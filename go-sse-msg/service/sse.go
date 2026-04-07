package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"go-sse-msg/config"
	"go-sse-msg/model"
	"go-sse-msg/pkg/ack"
	"go-sse-msg/pkg/offline"
	"go-sse-msg/pkg/stats"
	"go-sse-msg/repository"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// SSEService SSE 服务
type SSEService struct {
	config         *config.Config
	redis          *redis.Client
	ackManager     *ack.Manager
	presenceSvc    *PresenceService
	offlineMgr     *offline.Manager
	statsCollector *stats.Collector
	templateMgr    *model.TemplateManager

	// 连接管理
	clients     map[string]*Client // key: deviceID
	userClients map[string]map[string]bool // userID -> set of deviceIDs
	mu          sync.RWMutex
}

// Client SSE 客户端
type Client struct {
	UserID     string
	DeviceID   string
	DeviceType string
	Channel    chan []byte
	Quit       chan struct{}
	Connected  time.Time
	LastActive time.Time
}

// NewSSEService 创建 SSE 服务
func NewSSEService(cfg *config.Config, redisClient *redis.Client) *SSEService {
	s := &SSEService{
		config:      cfg,
		redis:       redisClient,
		clients:     make(map[string]*Client),
		userClients: make(map[string]map[string]bool),
	}

	// 初始化各个管理器
	s.ackManager = ack.NewManager(redisClient, cfg.SSE.AckTimeout, 3)
	s.presenceSvc = NewPresenceService(redisClient, cfg.SSE.HeartbeatInterval)
	s.offlineMgr = offline.NewManager(redisClient, cfg.Redis.OfflineQueue, 1000, 7)
	s.statsCollector = stats.NewCollector(redisClient, cfg.Stats.RetentionDays)
	s.templateMgr = model.NewTemplateManager(cfg.Template.MaxTemplates)

	// 注册默认模板
	for _, t := range model.GetDefaultTemplates() {
		s.templateMgr.Register(t)
	}

	// 设置 ACK 回调
	s.ackManager.SetCallbacks(
		s.onAckSuccess,
		s.onAckTimeout,
	)

	// 启动 Redis Pub/Sub 监听
	go s.subscribeRedis()

	return s
}

// Subscribe 客户端订阅 SSE
func (s *SSEService) Subscribe(c *gin.Context, userID, deviceID, deviceType string) {
	// 检查连接数限制
	s.mu.RLock()
	if len(s.clients) >= s.config.SSE.MaxConnections {
		s.mu.RUnlock()
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "max connections reached"})
		return
	}
	s.mu.RUnlock()

	// 注册设备
	if _, err := s.presenceSvc.RegisterDevice(userID, deviceID, deviceType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register device"})
		return
	}

	// 创建客户端
	client := &Client{
		UserID:     userID,
		DeviceID:   deviceID,
		DeviceType: deviceType,
		Channel:    make(chan []byte, s.config.SSE.MessageBufferSize),
		Quit:       make(chan struct{}),
		Connected:  time.Now(),
		LastActive: time.Now(),
	}

	s.mu.Lock()
	s.clients[deviceID] = client
	if s.userClients[userID] == nil {
		s.userClients[userID] = make(map[string]bool)
	}
	s.userClients[userID][deviceID] = true
	s.mu.Unlock()

	// 更新统计
	s.statsCollector.SetActiveConnections(len(s.clients))
	s.statsCollector.RecordUserActive(userID)

	// 推送离线消息
	go s.pushOfflineMessages(userID, client)

	// 设置 SSE 响应头
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("X-Accel-Buffering", "no")

	// 发送初始连接成功消息
	fmt.Fprintf(c.Writer, "event: connected\ndata: %s\n\n", `{"status":"connected"}`)
	c.Writer.Flush()

	// 监听连接关闭
	ctx := c.Request.Context()
	notify := c.Writer.CloseNotify()

	// 心跳协程
	heartbeatTicker := time.NewTicker(s.config.SSE.HeartbeatInterval)
	defer heartbeatTicker.Stop()

	// 消息处理循环
	for {
		select {
		case data := <-client.Channel:
			// 发送消息
			fmt.Fprintf(c.Writer, "event: message\ndata: %s\n\n", string(data))
			c.Writer.Flush()
			client.LastActive = time.Now()
			s.presenceSvc.UpdateHeartbeat(deviceID)

		case <-heartbeatTicker.C:
			// 发送心跳
			fmt.Fprintf(c.Writer, "event: ping\ndata: {}\n\n")
			c.Writer.Flush()
			client.LastActive = time.Now()
			s.presenceSvc.UpdateHeartbeat(deviceID)

		case <-notify:
			// 客户端断开
			goto cleanup

		case <-ctx.Done():
			// 请求上下文结束
			goto cleanup

		case <-client.Quit:
			// 主动退出
			goto cleanup
		}
	}

cleanup:
	s.cleanupClient(userID, deviceID)
}

// SendMessage 发送消息给指定用户
func (s *SSEService) SendMessage(msg *model.Message) error {
	// 生成消息ID
	if msg.ID == "" {
		msg.ID = generateID()
	}
	if msg.Timestamp.IsZero() {
		msg.Timestamp = time.Now()
	}

	// 统计
	s.statsCollector.RecordMessageSent(string(msg.Type), fmt.Sprintf("%d", msg.Priority))

	data, _ := json.Marshal(msg)

	// 如果指定了设备，只发给该设备
	if msg.DeviceID != "" {
		return s.sendToDevice(msg.DeviceID, data, msg.RequireAck)
	}

	// 获取用户所有在线设备
	s.mu.RLock()
	devices, ok := s.userClients[msg.UserID]
	s.mu.RUnlock()

	if !ok || len(devices) == 0 {
		// 用户不在线，存储离线消息
		return s.offlineMgr.Store(msg.UserID, data, int(msg.Priority), 0)
	}

	// 发送给所有在线设备（多端同步）
	for deviceID := range devices {
		s.sendToDevice(deviceID, data, msg.RequireAck)
	}

	return nil
}

// Broadcast 广播消息
func (s *SSEService) Broadcast(req *model.BroadcastRequest) error {
	msg := &model.Message{
		ID:         generateID(),
		Type:       req.Type,
		Priority:   req.Priority,
		Data:       req.Data,
		Timestamp:  time.Now(),
		RequireAck: req.RequireAck,
	}

	if req.TTL > 0 {
		expireAt := time.Now().Add(req.TTL)
		msg.ExpireAt = &expireAt
	}

	data, _ := json.Marshal(msg)

	// 发布到 Redis，让所有实例都能收到
	ctx := context.Background()
	return s.redis.Publish(ctx, "sse:broadcast", data).Err()
}

// SendWithTemplate 使用模板发送消息
func (s *SSEService) SendWithTemplate(userID, templateID, lang string, variables map[string]interface{}, requireAck bool) error {
	rendered, err := s.templateMgr.Render(templateID, lang, variables)
	if err != nil {
		return err
	}

	msg := &model.Message{
		ID:         generateID(),
		Type:       rendered.Type,
		Priority:   rendered.Priority,
		UserID:     userID,
		TemplateID: templateID,
		Data:       rendered.Data,
		Timestamp:  time.Now(),
		RequireAck: requireAck,
	}

	return s.SendMessage(msg)
}

// ProcessACK 处理 ACK 确认
func (s *SSEService) ProcessACK(messageID, userID string) bool {
	return s.ackManager.Confirm(messageID, userID)
}

// SyncMessages 多端同步消息
func (s *SSEService) SyncMessages(req *model.SyncRequest) (*model.SyncResponse, error) {
	// 从 Redis Stream 获取消息
	ctx := context.Background()
	streamKey := s.config.Redis.StreamName

	// 使用 XRead 或 XRange 获取消息
	// 简化版：获取最近的消息
	msgs, err := s.redis.XRevRangeN(ctx, streamKey, "+", "-", 100).Result()
	if err != nil {
		return nil, err
	}

	var messages []model.Message
	for _, msg := range msgs {
		// 解析消息
		if data, ok := msg.Values["data"]; ok {
			var m model.Message
			if err := json.Unmarshal([]byte(data.(string)), &m); err == nil {
				messages = append(messages, m)
			}
		}
	}

	return &model.SyncResponse{
		Messages:   messages,
		SyncID:     generateID(),
		SyncTime:   time.Now(),
		HasMore:    false,
		TotalCount: len(messages),
	}, nil
}

// 发送消息到指定设备
func (s *SSEService) sendToDevice(deviceID string, data []byte, requireAck bool) error {
	s.mu.RLock()
	client, ok := s.clients[deviceID]
	s.mu.RUnlock()

	if !ok {
		// 设备不在线，查找对应的用户并存储离线消息
		return fmt.Errorf("device not connected: %s", deviceID)
	}

	// 解析消息获取信息
	var msgData model.Message
	json.Unmarshal(data, &msgData)

	select {
	case client.Channel <- data:
		if requireAck && msgData.ID != "" {
			s.ackManager.Register(msgData.ID, client.UserID, deviceID, data)
		}
		if msgData.Type != "" {
			s.statsCollector.RecordMessageDelivered(string(msgData.Type), time.Since(msgData.Timestamp))
		}
		return nil
	default:
		// channel 已满，记录失败
		s.statsCollector.RecordMessageFailed("channel_full")
		return fmt.Errorf("client channel full")
	}
}

// 推送离线消息
func (s *SSEService) pushOfflineMessages(userID string, client *Client) {
	messages, err := s.offlineMgr.Fetch(userID, 100)
	if err != nil {
		return
	}

	for _, msg := range messages {
		select {
		case client.Channel <- msg.Message:
		case <-client.Quit:
			return
		case <-time.After(5 * time.Second):
			// 超时，消息可能未送达
		}
	}
}

// 清理客户端
func (s *SSEService) cleanupClient(userID, deviceID string) {
	s.mu.Lock()
	if client, ok := s.clients[deviceID]; ok {
		close(client.Quit)
		close(client.Channel)
		delete(s.clients, deviceID)
	}
	if devices, ok := s.userClients[userID]; ok {
		delete(devices, deviceID)
		if len(devices) == 0 {
			delete(s.userClients, userID)
		}
	}
	s.mu.Unlock()

	s.presenceSvc.UnregisterDevice(userID, deviceID)
	s.statsCollector.SetActiveConnections(len(s.clients))
}

// Redis Pub/Sub 订阅
func (s *SSEService) subscribeRedis() {
	ctx := context.Background()
	pubsub := s.redis.Subscribe(ctx, "sse:broadcast")
	defer pubsub.Close()

	ch := pubsub.Channel()
	for msg := range ch {
		var data map[string]interface{}
		if err := json.Unmarshal([]byte(msg.Payload), &data); err != nil {
			continue
		}

		// 解析广播消息
		msgBytes, _ := json.Marshal(data)
		var m model.Message
		if err := json.Unmarshal(msgBytes, &m); err != nil {
			continue
		}

		// 本地发送给所有在线用户
		s.mu.RLock()
		for userID := range s.userClients {
			// 检查是否在排除列表
			m.UserID = userID
			data, _ := json.Marshal(m)
			devices := s.userClients[userID]
			for deviceID := range devices {
				s.sendToDevice(deviceID, data, m.RequireAck)
			}
		}
		s.mu.RUnlock()
	}
}

// ACK 回调
func (s *SSEService) onAckSuccess(messageID, userID string) {
	// 更新统计
	s.statsCollector.RecordMessageDelivered("", 0)
}

func (s *SSEService) onAckTimeout(messageID, userID string, retryCount int) {
	// 重试发送
	data, ok := s.ackManager.GetRetryMessage(messageID, userID)
	if ok {
		// 获取用户的在线设备并重新发送
		devices, _ := s.presenceSvc.GetUserDevices(userID)
		for _, device := range devices {
			s.sendToDevice(device.ID, data, true)
		}
	}
}

// GetPresenceService 获取在线状态服务
func (s *SSEService) GetPresenceService() *PresenceService {
	return s.presenceSvc
}

// GetStatsCollector 获取统计收集器
func (s *SSEService) GetStatsCollector() *stats.Collector {
	return s.statsCollector
}

// GetOfflineManager 获取离线消息管理器
func (s *SSEService) GetOfflineManager() *offline.Manager {
	return s.offlineMgr
}

// GetAckManager 获取 ACK 管理器
func (s *SSEService) GetAckManager() *ack.Manager {
	return s.ackManager
}

// GetMessageRepository 获取消息仓库（V1 版本不支持）
func (s *SSEService) GetMessageRepository() *repository.MessageRepository {
	return nil
}

func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

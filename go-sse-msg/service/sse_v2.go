package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
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

// 对象池 - 复用缓冲区减少 GC
var (
	bufferPool = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}

	messagePool = sync.Pool{
		New: func() interface{} {
			return new(model.Message)
		},
	}
)

// SSEServiceV2 优化版 SSE 服务
type SSEServiceV2 struct {
	config         *config.Config
	redis          *redis.Client
	ackManager     *ack.Manager
	presenceSvc    *PresenceService
	offlineMgr     *offline.Manager
	statsCollector *stats.Collector
	templateMgr    *model.TemplateManager
	msgRepo        *repository.MessageRepository

	// 使用 sync.Map 替代 map + mutex，减少锁竞争
	clients     sync.Map // key: deviceID, value: *ClientV2
	userClients sync.Map // key: userID, value: *sync.Map[deviceID]bool

	// 连接数统计使用原子操作
	connectionCount int64

	// 批量发送通道
	broadcastChan chan *broadcastTask
	batchSize     int
	batchTimeout  time.Duration

	// 关闭信号
	stopChan chan struct{}
}

// ClientV2 优化版客户端
type ClientV2 struct {
	UserID      string
	DeviceID    string
	DeviceType  string
	Channel     chan []byte
	Quit        chan struct{}
	Connected   time.Time
	LastActive  int64 // 使用原子操作更新时间戳
	SendCounter int64 // 发送计数
}

// broadcastTask 广播任务
type broadcastTask struct {
	data       []byte
	requireAck bool
	priority   int
}

// NewSSEServiceV2 创建优化版 SSE 服务
func NewSSEServiceV2(cfg *config.Config, redisClient *redis.Client, msgRepo *repository.MessageRepository) *SSEServiceV2 {
	s := &SSEServiceV2{
		config:        cfg,
		redis:         redisClient,
		msgRepo:       msgRepo,
		broadcastChan: make(chan *broadcastTask, 10000),
		batchSize:     100,
		batchTimeout:  10 * time.Millisecond,
		stopChan:      make(chan struct{}),
	}

	// 初始化管理器
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
	s.ackManager.SetCallbacks(s.onAckSuccess, s.onAckTimeout)

	// 启动批量处理协程
	go s.batchProcessor()

	// 启动 Redis Pub/Sub
	go s.subscribeRedis()

	return s
}

// Subscribe 客户端订阅 - 优化版
func (s *SSEServiceV2) Subscribe(c *gin.Context, userID, deviceID, deviceType string) {
	// 快速检查连接数限制
	if atomic.LoadInt64(&s.connectionCount) >= int64(s.config.SSE.MaxConnections) {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "max connections reached"})
		return
	}

	// 异步注册设备
	go s.presenceSvc.RegisterDevice(userID, deviceID, deviceType)

	// 创建客户端
	client := &ClientV2{
		UserID:     userID,
		DeviceID:   deviceID,
		DeviceType: deviceType,
		Channel:    make(chan []byte, s.config.SSE.MessageBufferSize),
		Quit:       make(chan struct{}),
		Connected:  time.Now(),
		LastActive: time.Now().Unix(),
	}

	// 存储客户端
	s.clients.Store(deviceID, client)

	// 存储用户-设备关系
	devices, _ := s.userClients.LoadOrStore(userID, &sync.Map{})
	devices.(*sync.Map).Store(deviceID, true)

	// 原子增加连接数
	count := atomic.AddInt64(&s.connectionCount, 1)
	s.statsCollector.SetActiveConnections(int(count))
	s.statsCollector.RecordUserActive(userID)

	// 异步推送离线消息
	select {
	case <-s.stopChan:
		return
	default:
		go s.pushOfflineMessagesV2(userID, client)
	}

	// 设置 SSE 响应头
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("X-Accel-Buffering", "no")

	// 使用缓冲写入器
	flusher, ok := c.Writer.(http.Flusher)
	if !ok {
		s.cleanupClient(userID, deviceID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "streaming unsupported"})
		return
	}

	// 发送初始连接成功消息
	fmt.Fprintf(c.Writer, "event: connected\ndata: %s\n\n", `{"status":"connected"}`)
	flusher.Flush()

	// 监听连接关闭
	ctx := c.Request.Context()
	notify := c.Writer.CloseNotify()

	// 心跳协程
	heartbeatTicker := time.NewTicker(s.config.SSE.HeartbeatInterval)
	defer heartbeatTicker.Stop()

	// 消息批处理缓冲
	msgBuffer := make([][]byte, 0, 10)
	flushTicker := time.NewTicker(50 * time.Millisecond)
	defer flushTicker.Stop()

	// 消息处理循环
	for {
		select {
		case data := <-client.Channel:
			msgBuffer = append(msgBuffer, data)
			if len(msgBuffer) >= 10 {
				s.flushMessages(c.Writer, flusher, msgBuffer)
				msgBuffer = msgBuffer[:0]
			}
			atomic.StoreInt64(&client.LastActive, time.Now().Unix())

		case <-flushTicker.C:
			if len(msgBuffer) > 0 {
				s.flushMessages(c.Writer, flusher, msgBuffer)
				msgBuffer = msgBuffer[:0]
			}

		case <-heartbeatTicker.C:
			fmt.Fprintf(c.Writer, "event: ping\ndata: {}\n\n")
			flusher.Flush()
			client.LastActive = time.Now().Unix()
			go s.presenceSvc.UpdateHeartbeat(deviceID)

		case <-notify:
			goto cleanup

		case <-ctx.Done():
			goto cleanup

		case <-client.Quit:
			goto cleanup

		case <-s.stopChan:
			goto cleanup
		}
	}

cleanup:
	// 刷新剩余消息
	if len(msgBuffer) > 0 {
		s.flushMessages(c.Writer, flusher, msgBuffer)
	}
	s.cleanupClient(userID, deviceID)
}

// flushMessages 批量刷新消息
func (s *SSEServiceV2) flushMessages(w http.ResponseWriter, flusher http.Flusher, messages [][]byte) {
	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufferPool.Put(buf)

	for _, data := range messages {
		buf.WriteString("event: message\ndata: ")
		buf.Write(data)
		buf.WriteString("\n\n")
	}

	w.Write(buf.Bytes())
	flusher.Flush()
}

// SendMessageV2 优化版发送消息
func (s *SSEServiceV2) SendMessage(msg *model.Message) error {
	if msg.ID == "" {
		msg.ID = generateID()
	}
	if msg.Timestamp.IsZero() {
		msg.Timestamp = time.Now()
	}

	// 异步统计
	select {
	case <-s.stopChan:
	default:
		go s.statsCollector.RecordMessageSent(string(msg.Type), fmt.Sprintf("%d", msg.Priority))
	}

	// 持久化消息（异步双写）
	if s.msgRepo != nil {
		go func(m *model.Message) {
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			if err := s.msgRepo.Save(ctx, m); err != nil {
				// 记录错误但不影响发送
				// log.Printf("Failed to save message: %v", err)
			}
		}(msg)
	}

	// 使用对象池获取缓冲区
	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufferPool.Put(buf)

	encoder := json.NewEncoder(buf)
	if err := encoder.Encode(msg); err != nil {
		return err
	}
	data := buf.Bytes()

	// 如果指定了设备，直接发送
	if msg.DeviceID != "" {
		return s.sendToDevice(msg.DeviceID, data, msg.RequireAck, msg)
	}

	// 获取用户设备 - 使用 LoadOrStore 避免重复查找
	devicesI, ok := s.userClients.Load(msg.UserID)
	if !ok {
		// 用户不在线，存储离线消息（异步）
		go s.offlineMgr.Store(msg.UserID, data, int(msg.Priority), 0)
		return nil
	}

	// 发送给所有在线设备（多端同步）- 使用 Range 遍历
	devices := devicesI.(*sync.Map)
	var wg sync.WaitGroup

	devices.Range(func(key, value interface{}) bool {
		deviceID := key.(string)
		wg.Add(1)
		go func(did string) {
			defer wg.Done()
			s.sendToDevice(did, data, msg.RequireAck, msg)
		}(deviceID)
		return true
	})

	// 不等待发送完成，立即返回
	go wg.Wait()

	return nil
}

// batchProcessor 批量处理广播消息
func (s *SSEServiceV2) batchProcessor() {
	batch := make([]*broadcastTask, 0, s.batchSize)
	timer := time.NewTimer(s.batchTimeout)
	defer timer.Stop()

	for {
		select {
		case task := <-s.broadcastChan:
			batch = append(batch, task)
			if len(batch) >= s.batchSize {
				s.processBatch(batch)
				batch = batch[:0]
				timer.Reset(s.batchTimeout)
			}

		case <-timer.C:
			if len(batch) > 0 {
				s.processBatch(batch)
				batch = batch[:0]
			}
			timer.Reset(s.batchTimeout)

		case <-s.stopChan:
			if len(batch) > 0 {
				s.processBatch(batch)
			}
			return
		}
	}
}

// processBatch 处理批量广播
func (s *SSEServiceV2) processBatch(batch []*broadcastTask) {
	// 批量处理：合并相同优先级的消息
	s.clients.Range(func(key, value interface{}) bool {
		client := value.(*ClientV2)
		for _, task := range batch {
			select {
			case client.Channel <- task.data:
				atomic.AddInt64(&client.SendCounter, 1)
			default:
				// channel 已满，记录失败
				go s.statsCollector.RecordMessageFailed("channel_full")
			}
		}
		return true
	})
}

// sendToDevice 发送消息到指定设备 - 优化版
func (s *SSEServiceV2) sendToDevice(deviceID string, data []byte, requireAck bool, msg *model.Message) error {
	clientI, ok := s.clients.Load(deviceID)
	if !ok {
		return fmt.Errorf("device not connected: %s", deviceID)
	}
	client := clientI.(*ClientV2)

	select {
	case client.Channel <- data:
		if requireAck && msg.ID != "" {
			// 异步注册 ACK
			go s.ackManager.Register(msg.ID, client.UserID, deviceID, data)
		}
		// 异步统计
		go s.statsCollector.RecordMessageDelivered(string(msg.Type), time.Since(msg.Timestamp))
		return nil
	default:
		go s.statsCollector.RecordMessageFailed("channel_full")
		return fmt.Errorf("client channel full")
	}
}

// BroadcastV2 优化版广播
func (s *SSEServiceV2) Broadcast(req *model.BroadcastRequest) error {
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

	// 使用对象池
	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufferPool.Put(buf)

	encoder := json.NewEncoder(buf)
	if err := encoder.Encode(msg); err != nil {
		return err
	}

	// 放入批量处理队列
	select {
	case s.broadcastChan <- &broadcastTask{data: buf.Bytes(), requireAck: req.RequireAck, priority: int(req.Priority)}:
		return nil
	default:
		// 队列已满，直接处理
		go func() {
			s.processBatch([]*broadcastTask{{data: buf.Bytes(), requireAck: req.RequireAck}})
		}()
		return nil
	}
}

// 其他方法保持兼容
func (s *SSEServiceV2) SendWithTemplate(userID, templateID, lang string, variables map[string]interface{}, requireAck bool) error {
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

func (s *SSEServiceV2) ProcessACK(messageID, userID string) bool {
	return s.ackManager.Confirm(messageID, userID)
}

func (s *SSEServiceV2) SyncMessages(req *model.SyncRequest) (*model.SyncResponse, error) {
	ctx := context.Background()
	streamKey := s.config.Redis.StreamName

	msgs, err := s.redis.XRevRangeN(ctx, streamKey, "+", "-", 100).Result()
	if err != nil {
		return nil, err
	}

	messages := make([]model.Message, 0, len(msgs))
	for _, msg := range msgs {
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

// pushOfflineMessages 异步推送离线消息
func (s *SSEServiceV2) pushOfflineMessagesV2(userID string, client *ClientV2) {
	messages, err := s.offlineMgr.Fetch(userID, 100)
	if err != nil || len(messages) == 0 {
		return
	}

	// 批量发送
	for _, msg := range messages {
		select {
		case client.Channel <- msg.Message:
		case <-client.Quit:
			return
		case <-s.stopChan:
			return
		case <-time.After(5 * time.Second):
			return
		}
	}
}

// cleanupClient 清理客户端
func (s *SSEServiceV2) cleanupClient(userID, deviceID string) {
	s.clients.Delete(deviceID)

	if devicesI, ok := s.userClients.Load(userID); ok {
		devices := devicesI.(*sync.Map)
		devices.Delete(deviceID)

		// 检查是否还有设备
		hasDevices := false
		devices.Range(func(_, _ interface{}) bool {
			hasDevices = true
			return false
		})
		if !hasDevices {
			s.userClients.Delete(userID)
		}
	}

	count := atomic.AddInt64(&s.connectionCount, -1)
	s.statsCollector.SetActiveConnections(int(count))

	go s.presenceSvc.UnregisterDevice(userID, deviceID)
}

// subscribeRedis Redis 订阅
func (s *SSEServiceV2) subscribeRedis() {
	ctx := context.Background()
	pubsub := s.redis.Subscribe(ctx, "sse:broadcast")
	defer pubsub.Close()

	ch := pubsub.Channel()
	for {
		select {
		case msg := <-ch:
			if msg == nil {
				return
			}
			var data map[string]interface{}
			if err := json.Unmarshal([]byte(msg.Payload), &data); err != nil {
				continue
			}

			msgBytes, _ := json.Marshal(data)
			var m model.Message
			if err := json.Unmarshal(msgBytes, &m); err != nil {
				continue
			}

			// 批量广播给本地客户端
			s.clients.Range(func(key, value interface{}) bool {
				client := value.(*ClientV2)
				m.UserID = client.UserID
				data, _ := json.Marshal(m)
				select {
				case client.Channel <- data:
				default:
				}
				return true
			})

		case <-s.stopChan:
			return
		}
	}
}

// onAckSuccess ACK 成功回调
func (s *SSEServiceV2) onAckSuccess(messageID, userID string) {
	go s.statsCollector.RecordMessageDelivered("", 0)
}

// onAckTimeout ACK 超时回调
func (s *SSEServiceV2) onAckTimeout(messageID, userID string, retryCount int) {
	data, ok := s.ackManager.GetRetryMessage(messageID, userID)
	if !ok {
		return
	}

	// 异步重试
	go func() {
		devices, _ := s.presenceSvc.GetUserDevices(userID)
		var msg model.Message
		json.Unmarshal(data, &msg)

		for _, device := range devices {
			s.sendToDevice(device.ID, data, true, &msg)
		}
	}()
}

// Getters
func (s *SSEServiceV2) GetPresenceService() *PresenceService { return s.presenceSvc }
func (s *SSEServiceV2) GetStatsCollector() *stats.Collector  { return s.statsCollector }
func (s *SSEServiceV2) GetOfflineManager() *offline.Manager  { return s.offlineMgr }
func (s *SSEServiceV2) GetAckManager() *ack.Manager                  { return s.ackManager }
func (s *SSEServiceV2) GetMessageRepository() *repository.MessageRepository { return s.msgRepo }
func (s *SSEServiceV2) GetConnectionCount() int64                    { return atomic.LoadInt64(&s.connectionCount) }

// Stop 优雅关闭
func (s *SSEServiceV2) Stop() {
	close(s.stopChan)
	time.Sleep(100 * time.Millisecond)

	// 关闭所有客户端
	s.clients.Range(func(key, value interface{}) bool {
		client := value.(*ClientV2)
		close(client.Quit)
		return true
	})
}

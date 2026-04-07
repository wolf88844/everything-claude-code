package service

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"go-sse-msg/model"
	"github.com/redis/go-redis/v9"
)

// PresenceService 在线状态服务
type PresenceService struct {
	redis        *redis.Client
	localClients map[string]*model.Device // 本地内存中的设备
	mu           sync.RWMutex
	heartbeatTTL time.Duration
}

// NewPresenceService 创建在线状态服务
func NewPresenceService(redis *redis.Client, heartbeatTTL time.Duration) *PresenceService {
	s := &PresenceService{
		redis:        redis,
		localClients: make(map[string]*model.Device),
		heartbeatTTL: heartbeatTTL,
	}
	// 启动心跳检查
	go s.heartbeatLoop()
	return s
}

// RegisterDevice 注册设备上线
func (s *PresenceService) RegisterDevice(userID, deviceID, deviceType string) (*model.Device, error) {
	device := &model.Device{
		ID:        deviceID,
		UserID:    userID,
		Type:      deviceType,
		Online:    true,
		LastSeen:  time.Now(),
		Connected: time.Now(),
	}

	s.mu.Lock()
	s.localClients[deviceID] = device
	s.mu.Unlock()

	// 写入 Redis
	ctx := context.Background()
	deviceKey := fmt.Sprintf("presence:device:%s", deviceID)
	data, _ := json.Marshal(device)

	pipe := s.redis.Pipeline()
	pipe.Set(ctx, deviceKey, data, s.heartbeatTTL*2)
	pipe.HSet(ctx, fmt.Sprintf("presence:user:%s", userID), deviceID, time.Now().Unix())
	pipe.SAdd(ctx, fmt.Sprintf("presence:devices:%s", userID), deviceID)
	pipe.Expire(ctx, fmt.Sprintf("presence:user:%s", userID), s.heartbeatTTL*2)
	pipe.Expire(ctx, fmt.Sprintf("presence:devices:%s", userID), s.heartbeatTTL*2)
	_, err := pipe.Exec(ctx)

	return device, err
}

// UpdateHeartbeat 更新心跳
func (s *PresenceService) UpdateHeartbeat(deviceID string) error {
	s.mu.Lock()
	if device, ok := s.localClients[deviceID]; ok {
		device.LastSeen = time.Now()
	}
	s.mu.Unlock()

	ctx := context.Background()
	deviceKey := fmt.Sprintf("presence:device:%s", deviceID)
	return s.redis.Expire(ctx, deviceKey, s.heartbeatTTL*2).Err()
}

// UnregisterDevice 设备下线
func (s *PresenceService) UnregisterDevice(userID, deviceID string) error {
	s.mu.Lock()
	delete(s.localClients, deviceID)
	s.mu.Unlock()

	ctx := context.Background()
	deviceKey := fmt.Sprintf("presence:device:%s", deviceID)

	pipe := s.redis.Pipeline()
	pipe.Del(ctx, deviceKey)
	pipe.HDel(ctx, fmt.Sprintf("presence:user:%s", userID), deviceID)
	pipe.SRem(ctx, fmt.Sprintf("presence:devices:%s", userID), deviceID)
	_, err := pipe.Exec(ctx)

	return err
}

// GetUserDevices 获取用户所有设备
func (s *PresenceService) GetUserDevices(userID string) ([]model.Device, error) {
	ctx := context.Background()
	deviceIDs, err := s.redis.SMembers(ctx, fmt.Sprintf("presence:devices:%s", userID)).Result()
	if err != nil {
		return nil, err
	}

	var devices []model.Device
	for _, id := range deviceIDs {
		data, err := s.redis.Get(ctx, fmt.Sprintf("presence:device:%s", id)).Result()
		if err != nil {
			continue
		}
		var device model.Device
		if err := json.Unmarshal([]byte(data), &device); err != nil {
			continue
		}
		devices = append(devices, device)
	}

	return devices, nil
}

// GetOnlineStatus 获取用户在线状态
func (s *PresenceService) GetOnlineStatus(userID string) (*model.OnlineStatus, error) {
	devices, err := s.GetUserDevices(userID)
	if err != nil {
		return nil, err
	}

	status := &model.OnlineStatus{
		UserID:   userID,
		Devices:  devices,
		Platform: make(map[string]bool),
	}

	online := false
	var lastSeen time.Time
	var connectTime time.Time

	for _, d := range devices {
		if d.Online {
			online = true
			status.Platform[d.Type] = true
			if lastSeen.IsZero() || d.LastSeen.After(lastSeen) {
				lastSeen = d.LastSeen
			}
			if connectTime.IsZero() || d.Connected.Before(connectTime) {
				connectTime = d.Connected
			}
		}
	}

	status.Online = online
	status.LastSeenAt = lastSeen
	status.ConnectTime = connectTime

	return status, nil
}

// IsUserOnline 检查用户是否在线
func (s *PresenceService) IsUserOnline(userID string) bool {
	ctx := context.Background()
	count, _ := s.redis.SCard(ctx, fmt.Sprintf("presence:devices:%s", userID)).Result()
	return count > 0
}

// GetAllOnlineUsers 获取所有在线用户（简化版）
func (s *PresenceService) GetAllOnlineUsers() ([]string, error) {
	ctx := context.Background()
	// 使用 Scan 扫描所有 presence:user:* 键
	var users []string
	iter := s.redis.Scan(ctx, 0, "presence:devices:*", 100).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		// 提取 userID
		var userID string
		fmt.Sscanf(key, "presence:devices:%s", &userID)
		if userID != "" {
			users = append(users, userID)
		}
	}
	return users, iter.Err()
}

// heartbeatLoop 心跳检查循环
func (s *PresenceService) heartbeatLoop() {
	ticker := time.NewTicker(s.heartbeatTTL)
	defer ticker.Stop()

	for range ticker.C {
		s.checkStaleDevices()
	}
}

// checkStaleDevices 检查过期设备
func (s *PresenceService) checkStaleDevices() {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	for deviceID, device := range s.localClients {
		if now.Sub(device.LastSeen) > s.heartbeatTTL*2 {
			// 设备已过期
			device.Online = false
			delete(s.localClients, deviceID)
			// 异步清理 Redis
			go s.UnregisterDevice(device.UserID, deviceID)
		}
	}
}

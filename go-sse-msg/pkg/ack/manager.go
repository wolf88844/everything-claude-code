package ack

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

// Manager ACK 管理器
type Manager struct {
	redis      *redis.Client
	pending    map[string]*AckRequest // 等待确认的消息
	mu         sync.RWMutex
	timeout    time.Duration
	maxRetry   int
	retryDelay time.Duration
	// 回调函数
	onAckSuccess func(messageID, userID string)
	onAckTimeout func(messageID, userID string, retryCount int)
}

// AckRequest ACK 请求
type AckRequest struct {
	MessageID   string
	UserID      string
	DeviceID    string
	SentAt      time.Time
	TimeoutAt   time.Time
	RetryCount  int
	MaxRetry    int
	Data        []byte
	Confirmed   bool
}

// NewManager 创建 ACK 管理器
func NewManager(redis *redis.Client, timeout time.Duration, maxRetry int) *Manager {
	m := &Manager{
		redis:      redis,
		pending:    make(map[string]*AckRequest),
		timeout:    timeout,
		maxRetry:   maxRetry,
		retryDelay: 2 * time.Second,
	}
	// 启动超时检查协程
	go m.checkTimeoutLoop()
	return m
}

// SetCallbacks 设置回调
func (m *Manager) SetCallbacks(onSuccess func(messageID, userID string), onTimeout func(messageID, userID string, retryCount int)) {
	m.onAckSuccess = onSuccess
	m.onAckTimeout = onTimeout
}

// Register 注册需要确认的消息
func (m *Manager) Register(messageID, userID, deviceID string, data []byte) {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := m.key(messageID, userID)
	m.pending[key] = &AckRequest{
		MessageID:  messageID,
		UserID:     userID,
		DeviceID:   deviceID,
		SentAt:     time.Now(),
		TimeoutAt:  time.Now().Add(m.timeout),
		RetryCount: 0,
		MaxRetry:   m.maxRetry,
		Data:       data,
		Confirmed:  false,
	}

	// 同时写入 Redis 持久化
	ctx := context.Background()
	m.redis.HSet(ctx, "ack:pending:"+userID, messageID, time.Now().Unix())
	m.redis.Expire(ctx, "ack:pending:"+userID, m.timeout*2)
}

// Confirm 确认收到消息
func (m *Manager) Confirm(messageID, userID string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := m.key(messageID, userID)
	req, ok := m.pending[key]
	if !ok {
		// 检查 Redis 中是否还有记录（可能已超时但在重试）
		ctx := context.Background()
		exists := m.redis.HExists(ctx, "ack:pending:"+userID, messageID).Val()
		if !exists {
			return false
		}
	} else {
		req.Confirmed = true
	}

	// 清理记录
	delete(m.pending, key)
	ctx := context.Background()
	m.redis.HDel(ctx, "ack:pending:"+userID, messageID)
	m.redis.ZAdd(ctx, "ack:confirmed", redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: fmt.Sprintf("%s:%s", userID, messageID),
	})

	if m.onAckSuccess != nil {
		m.onAckSuccess(messageID, userID)
	}

	return true
}

// IsPending 检查消息是否在等待确认
func (m *Manager) IsPending(messageID, userID string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	key := m.key(messageID, userID)
	_, ok := m.pending[key]
	return ok
}

// GetPendingList 获取用户的待确认消息列表
func (m *Manager) GetPendingList(userID string) []AckRequest {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var list []AckRequest
	prefix := userID + ":"
	for key, req := range m.pending {
		if len(key) > len(prefix) && key[:len(prefix)] == prefix {
			list = append(list, *req)
		}
	}
	return list
}

// 超时检查循环
func (m *Manager) checkTimeoutLoop() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		m.checkTimeout()
	}
}

// 检查超时
func (m *Manager) checkTimeout() {
	m.mu.Lock()
	defer m.mu.Unlock()

	now := time.Now()
	for key, req := range m.pending {
		if req.Confirmed {
			continue
		}

		if now.After(req.TimeoutAt) {
			// 超时处理
			if req.RetryCount < req.MaxRetry {
				// 重试
				req.RetryCount++
				req.TimeoutAt = now.Add(m.timeout)
				req.SentAt = now

				if m.onAckTimeout != nil {
					m.onAckTimeout(req.MessageID, req.UserID, req.RetryCount)
				}
			} else {
				// 超过最大重试次数，标记为失败
				delete(m.pending, key)
				ctx := context.Background()
				m.redis.HDel(ctx, "ack:pending:"+req.UserID, req.MessageID)
				m.redis.ZAdd(ctx, "ack:failed", redis.Z{
					Score:  float64(now.Unix()),
					Member: fmt.Sprintf("%s:%s", req.UserID, req.MessageID),
				})

				if m.onAckTimeout != nil {
					m.onAckTimeout(req.MessageID, req.UserID, req.RetryCount)
				}
			}
		}
	}
}

// GetRetryMessage 获取需要重试的消息
func (m *Manager) GetRetryMessage(messageID, userID string) ([]byte, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	key := m.key(messageID, userID)
	req, ok := m.pending[key]
	if !ok {
		return nil, false
	}
	return req.Data, true
}

func (m *Manager) key(messageID, userID string) string {
	return fmt.Sprintf("%s:%s", userID, messageID)
}

// GetStats 获取 ACK 统计
func (m *Manager) GetStats(userID string) map[string]interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()

	pendingCount := 0
	prefix := userID + ":"
	for key := range m.pending {
		if len(key) > len(prefix) && key[:len(prefix)] == prefix {
			pendingCount++
		}
	}

	ctx := context.Background()
	confirmedCount := m.redis.ZCard(ctx, "ack:confirmed").Val()
	failedCount := m.redis.ZCard(ctx, "ack:failed").Val()

	return map[string]interface{}{
		"pending_count":   pendingCount,
		"confirmed_count": confirmedCount,
		"failed_count":    failedCount,
	}
}

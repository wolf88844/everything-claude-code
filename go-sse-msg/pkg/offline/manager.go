package offline

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// Manager 离线消息管理器
type Manager struct {
	redis      *redis.Client
	queuePrefix string
	maxQueueSize int
	retentionDays int
}

// OfflineMessage 离线消息
type OfflineMessage struct {
	ID         string          `json:"id"`
	UserID     string          `json:"user_id"`
	Message    json.RawMessage `json:"message"`
	SentAt     time.Time       `json:"sent_at"`
	ExpireAt   *time.Time      `json:"expire_at,omitempty"`
	Priority   int             `json:"priority"`
}

// NewManager 创建离线消息管理器
func NewManager(redis *redis.Client, queuePrefix string, maxSize, retentionDays int) *Manager {
	return &Manager{
		redis:         redis,
		queuePrefix:   queuePrefix,
		maxQueueSize:  maxSize,
		retentionDays: retentionDays,
	}
}

// Store 存储离线消息
func (m *Manager) Store(userID string, message []byte, priority int, ttl time.Duration) error {
	ctx := context.Background()
	queueKey := m.queuePrefix + userID

	// 检查队列大小
	queueLen, err := m.redis.LLen(ctx, queueKey).Result()
	if err != nil {
		return err
	}

	// 如果队列已满，根据优先级替换
	if queueLen >= int64(m.maxQueueSize) {
		// 获取队列中最低优先级的消息
		// 简化处理：直接丢弃最旧的消息
		m.redis.RPop(ctx, queueKey)
	}

	msg := OfflineMessage{
		ID:       generateID(),
		UserID:   userID,
		Message:  message,
		SentAt:   time.Now(),
		Priority: priority,
	}

	if ttl > 0 {
		expireAt := time.Now().Add(ttl)
		msg.ExpireAt = &expireAt
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	// 使用 LPUSH 添加到队列头部（最新消息在前）
	if err := m.redis.LPush(ctx, queueKey, data).Err(); err != nil {
		return err
	}

	// 设置过期时间
	m.redis.Expire(ctx, queueKey, time.Duration(m.retentionDays)*24*time.Hour)

	// 更新用户离线消息计数
	m.redis.HIncrBy(ctx, "offline:count", userID, 1)

	return nil
}

// Fetch 获取用户的离线消息
func (m *Manager) Fetch(userID string, limit int64) ([]OfflineMessage, error) {
	ctx := context.Background()
	queueKey := m.queuePrefix + userID

	// 获取队列中的消息
	datas, err := m.redis.LRange(ctx, queueKey, 0, limit-1).Result()
	if err != nil {
		return nil, err
	}

	var messages []OfflineMessage
	var expiredKeys []int64

	for i, data := range datas {
		var msg OfflineMessage
		if err := json.Unmarshal([]byte(data), &msg); err != nil {
			continue
		}

		// 检查是否过期
		if msg.ExpireAt != nil && time.Now().After(*msg.ExpireAt) {
			expiredKeys = append(expiredKeys, int64(i))
			continue
		}

		messages = append(messages, msg)
	}

	// 删除已获取的消息
	if len(datas) > 0 {
		m.redis.LTrim(ctx, queueKey, int64(len(datas)), -1)
		m.redis.HIncrBy(ctx, "offline:count", userID, -int64(len(datas)))
	}

	// 清理过期消息
	for _, idx := range expiredKeys {
		_ = idx // 简化处理
	}

	return messages, nil
}

// GetCount 获取用户离线消息数量
func (m *Manager) GetCount(userID string) int64 {
	ctx := context.Background()
	count, _ := m.redis.HGet(ctx, "offline:count", userID).Int64()
	return count
}

// Clear 清空用户离线消息
func (m *Manager) Clear(userID string) error {
	ctx := context.Background()
	queueKey := m.queuePrefix + userID

	if err := m.redis.Del(ctx, queueKey).Err(); err != nil {
		return err
	}

	m.redis.HDel(ctx, "offline:count", userID)
	return nil
}

// Peek 预览离线消息（不删除）
func (m *Manager) Peek(userID string, limit int64) ([]OfflineMessage, error) {
	ctx := context.Background()
	queueKey := m.queuePrefix + userID

	datas, err := m.redis.LRange(ctx, queueKey, 0, limit-1).Result()
	if err != nil {
		return nil, err
	}

	var messages []OfflineMessage
	for _, data := range datas {
		var msg OfflineMessage
		if err := json.Unmarshal([]byte(data), &msg); err != nil {
			continue
		}
		messages = append(messages, msg)
	}

	return messages, nil
}

// GetAllUserIDs 获取所有有待处理离线消息的用户
func (m *Manager) GetAllUserIDs() ([]string, error) {
	ctx := context.Background()
	return m.redis.HKeys(ctx, "offline:count").Result()
}

// CleanupExpired 清理过期消息
func (m *Manager) CleanupExpired() error {
	_, err := m.GetAllUserIDs()
	if err != nil {
		return err
	}

	// 简化处理：遍历检查每条消息
	// 生产环境可以使用 Redis 过期时间或定期扫描

	return nil
}

func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

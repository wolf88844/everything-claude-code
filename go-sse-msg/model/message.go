package model

import (
	"encoding/json"
	"time"
)

// MessageType 消息类型
type MessageType string

const (
	MessageTypeNotification MessageType = "notification"
	MessageTypeSystem       MessageType = "system"
	MessageTypeBusiness     MessageType = "business"
	MessageTypeChat         MessageType = "chat"
)

// MessagePriority 消息优先级
type MessagePriority int

const (
	PriorityLow    MessagePriority = 1
	PriorityNormal MessagePriority = 5
	PriorityHigh   MessagePriority = 10
	PriorityUrgent MessagePriority = 20
)

// Message 业务消息结构
type Message struct {
	ID         string          `json:"id"`
	Type       MessageType     `json:"type"`
	Priority   MessagePriority `json:"priority"`
	UserID     string          `json:"user_id"`
	DeviceID   string          `json:"device_id"`
	TemplateID string          `json:"template_id,omitempty"`
	Data       interface{}     `json:"data"`
	Timestamp  time.Time       `json:"timestamp"`
	ExpireAt   *time.Time      `json:"expire_at,omitempty"`
	RequireAck bool            `json:"require_ack,omitempty"`
}

// ToJSON 转换为 JSON 字节
func (m *Message) ToJSON() []byte {
	data, _ := json.Marshal(m)
	return data
}

// FromJSON 从 JSON 解析
func (m *Message) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// AckMessage ACK 确认消息
type AckMessage struct {
	MessageID string    `json:"message_id"`
	UserID    string    `json:"user_id"`
	DeviceID  string    `json:"device_id"`
	Timestamp time.Time `json:"timestamp"`
	Status    AckStatus `json:"status"`
}

type AckStatus string

const (
	AckStatusReceived  AckStatus = "received"
	AckStatusProcessed AckStatus = "processed"
	AckStatusFailed    AckStatus = "failed"
)

// Device 设备信息
type Device struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Type      string    `json:"type"` // web, ios, android, desktop
	Online    bool      `json:"online"`
	LastSeen  time.Time `json:"last_seen"`
	Connected time.Time `json:"connected"`
}

// OnlineStatus 用户在线状态
type OnlineStatus struct {
	UserID      string            `json:"user_id"`
	Online      bool              `json:"online"`
	Devices     []Device          `json:"devices"`
	LastSeenAt  time.Time         `json:"last_seen_at"`
	ConnectTime time.Time         `json:"connect_time"`
	Platform    map[string]bool   `json:"platform"` // web, mobile, desktop
}

// DeliveryStatus 消息投递状态
type DeliveryStatus struct {
	MessageID     string            `json:"message_id"`
	UserID        string            `json:"user_id"`
	DeviceID      string            `json:"device_id"`
	SentAt        time.Time         `json:"sent_at"`
	DeliveredAt   *time.Time        `json:"delivered_at,omitempty"`
	AckedAt       *time.Time        `json:"acked_at,omitempty"`
	Status        string            `json:"status"` // pending, sent, delivered, acked, failed
	RetryCount    int               `json:"retry_count"`
	ErrorMessage  string            `json:"error_message,omitempty"`
}

// BroadcastRequest 广播消息请求
type BroadcastRequest struct {
	UserIDs    []string        `json:"user_ids,omitempty"`    // 指定用户，为空则广播全部
	ExcludeIDs []string        `json:"exclude_ids,omitempty"` // 排除用户
	Type       MessageType     `json:"type"`
	Priority   MessagePriority `json:"priority"`
	TemplateID string          `json:"template_id,omitempty"`
	Data       interface{}     `json:"data"`
	RequireAck bool            `json:"require_ack"`
	TTL        time.Duration   `json:"ttl,omitempty"`
}

// SyncRequest 多端同步请求
type SyncRequest struct {
	UserID      string    `json:"user_id"`
	DeviceID    string    `json:"device_id"`
	LastSyncID  string    `json:"last_sync_id"`
	LastSyncTime time.Time `json:"last_sync_time"`
}

// SyncResponse 同步响应
type SyncResponse struct {
	Messages   []Message `json:"messages"`
	SyncID     string    `json:"sync_id"`
	SyncTime   time.Time `json:"sync_time"`
	HasMore    bool      `json:"has_more"`
	TotalCount int       `json:"total_count"`
}

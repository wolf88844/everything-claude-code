package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"go-sse-msg/model"
	"github.com/redis/go-redis/v9"
)

// MessageRepository 消息存储仓库
type MessageRepository struct {
	db     *sql.DB
	redis  *redis.Client
	config *RepositoryConfig
}

// RepositoryConfig 仓库配置
type RepositoryConfig struct {
	RedisStreamName     string
	RedisRetentionHours int
	DBTableName         string
	BatchSize           int
	FlushInterval       time.Duration
}

// DefaultRepositoryConfig 默认配置
func DefaultRepositoryConfig() *RepositoryConfig {
	return &RepositoryConfig{
		RedisStreamName:     "sse:messages",
		RedisRetentionHours: 24,
		DBTableName:         "messages",
		BatchSize:           100,
		FlushInterval:       5 * time.Second,
	}
}

// NewMessageRepository 创建消息仓库
func NewMessageRepository(db *sql.DB, redis *redis.Client, config *RepositoryConfig) *MessageRepository {
	if config == nil {
		config = DefaultRepositoryConfig()
	}

	r := &MessageRepository{
		db:     db,
		redis:  redis,
		config: config,
	}

	// 确保表存在
	r.ensureTable()

	return r
}

// ensureTable 确保数据库表存在
func (r *MessageRepository) ensureTable() {
	sql := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id VARCHAR(64) PRIMARY KEY,
			user_id VARCHAR(64) NOT NULL,
			device_id VARCHAR(64),
			message_type VARCHAR(32) NOT NULL,
			priority INT DEFAULT 5,
			template_id VARCHAR(64),
			data JSON NOT NULL,
			require_ack BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			expire_at TIMESTAMP NULL,
			delivered BOOLEAN DEFAULT FALSE,
			delivered_at TIMESTAMP NULL,
			acked BOOLEAN DEFAULT FALSE,
			acked_at TIMESTAMP NULL,
			INDEX idx_user_id (user_id, created_at DESC),
			INDEX idx_created_at (created_at DESC),
			INDEX idx_message_type (message_type)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
	`, r.config.DBTableName)

	r.db.Exec(sql)
}

// Save 保存消息（双写：Redis Stream + MySQL）
func (r *MessageRepository) Save(ctx context.Context, msg *model.Message) error {
	// 1. 写入 Redis Stream（用于实时分发）
	if err := r.saveToRedisStream(ctx, msg); err != nil {
		return fmt.Errorf("failed to save to redis stream: %w", err)
	}

	// 2. 写入 MySQL（用于持久化存储）
	if err := r.saveToMySQL(ctx, msg); err != nil {
		return fmt.Errorf("failed to save to mysql: %w", err)
	}

	return nil
}

// saveToRedisStream 保存到 Redis Stream
func (r *MessageRepository) saveToRedisStream(ctx context.Context, msg *model.Message) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	values := map[string]interface{}{
		"id":         msg.ID,
		"user_id":    msg.UserID,
		"type":       string(msg.Type),
		"priority":   msg.Priority,
		"data":       string(data),
		"created_at": msg.Timestamp.Unix(),
	}

	if msg.RequireAck {
		values["require_ack"] = 1
	}

	// 添加到 Stream，保留最近 10000 条
	_, err = r.redis.XAdd(ctx, &redis.XAddArgs{
		Stream: r.config.RedisStreamName,
		MaxLen: 10000,
		Approx: true,
		Values: values,
	}).Result()

	return err
}

// saveToMySQL 保存到 MySQL
func (r *MessageRepository) saveToMySQL(ctx context.Context, msg *model.Message) error {
	data, err := json.Marshal(msg.Data)
	if err != nil {
		return err
	}

	sql := fmt.Sprintf(`
		INSERT INTO %s (id, user_id, device_id, message_type, priority, template_id, data, require_ack, created_at, expire_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE id=id
	`, r.config.DBTableName)

	_, err = r.db.ExecContext(ctx, sql,
		msg.ID,
		msg.UserID,
		msg.DeviceID,
		string(msg.Type),
		msg.Priority,
		msg.TemplateID,
		data,
		msg.RequireAck,
		msg.Timestamp,
		msg.ExpireAt,
	)

	return err
}

// GetByID 根据 ID 获取消息
func (r *MessageRepository) GetByID(ctx context.Context, messageID string) (*model.Message, error) {
	sql := fmt.Sprintf(`
		SELECT id, user_id, device_id, message_type, priority, template_id, data, require_ack, created_at, expire_at
		FROM %s WHERE id = ?
	`, r.config.DBTableName)

	row := r.db.QueryRowContext(ctx, sql, messageID)
	return r.scanMessage(row)
}

// GetByUser 获取用户的消息列表（分页）
func (r *MessageRepository) GetByUser(ctx context.Context, userID string, limit, offset int) ([]*model.Message, error) {
	sql := fmt.Sprintf(`
		SELECT id, user_id, device_id, message_type, priority, template_id, data, require_ack, created_at, expire_at
		FROM %s
		WHERE user_id = ?
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?
	`, r.config.DBTableName)

	rows, err := r.db.QueryContext(ctx, sql, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*model.Message
	for rows.Next() {
		msg, err := r.scanMessage(rows)
		if err != nil {
			continue
		}
		messages = append(messages, msg)
	}

	return messages, rows.Err()
}

// GetByTimeRange 按时间范围获取消息
func (r *MessageRepository) GetByTimeRange(ctx context.Context, userID string, start, end time.Time, limit int) ([]*model.Message, error) {
	sql := fmt.Sprintf(`
		SELECT id, user_id, device_id, message_type, priority, template_id, data, require_ack, created_at, expire_at
		FROM %s
		WHERE user_id = ? AND created_at >= ? AND created_at <= ?
		ORDER BY created_at DESC
		LIMIT ?
	`, r.config.DBTableName)

	rows, err := r.db.QueryContext(ctx, sql, userID, start, end, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*model.Message
	for rows.Next() {
		msg, err := r.scanMessage(rows)
		if err != nil {
			continue
		}
		messages = append(messages, msg)
	}

	return messages, rows.Err()
}

// GetUnreadMessages 获取用户的未读消息（用于断线恢复）
func (r *MessageRepository) GetUnreadMessages(ctx context.Context, userID string, lastReadID string, limit int) ([]*model.Message, error) {
	// 从 Redis Stream 获取最近消息
	if lastReadID == "" {
		lastReadID = "0"
	}

	// 先尝试从 Redis Stream 读取
	msgs, err := r.redis.XRead(ctx, &redis.XReadArgs{
		Streams: []string{r.config.RedisStreamName, lastReadID},
		Count:   int64(limit),
		Block:   0,
	}).Result()

	if err == nil && len(msgs) > 0 {
		var messages []*model.Message
		for _, stream := range msgs {
			for _, msg := range stream.Messages {
				if m := r.parseStreamMessage(&msg); m != nil {
					messages = append(messages, m)
				}
			}
		}
		if len(messages) > 0 {
			return messages, nil
		}
	}

	// 从 MySQL 获取
	var lastReadTime time.Time
	if lastReadID != "" && lastReadID != "0" {
		// 尝试解析时间戳
		msg, _ := r.GetByID(ctx, lastReadID)
		if msg != nil {
			lastReadTime = msg.Timestamp
		}
	}

	if lastReadTime.IsZero() {
		lastReadTime = time.Now().Add(-24 * time.Hour) // 默认24小时
	}

	return r.GetByTimeRange(ctx, userID, lastReadTime, time.Now(), limit)
}

// MarkDelivered 标记消息已送达
func (r *MessageRepository) MarkDelivered(ctx context.Context, messageID string) error {
	sql := fmt.Sprintf(`
		UPDATE %s SET delivered = TRUE, delivered_at = NOW() WHERE id = ?
	`, r.config.DBTableName)
	_, err := r.db.ExecContext(ctx, sql, messageID)
	return err
}

// MarkAcked 标记消息已确认
func (r *MessageRepository) MarkAcked(ctx context.Context, messageID string) error {
	sql := fmt.Sprintf(`
		UPDATE %s SET acked = TRUE, acked_at = NOW() WHERE id = ?
	`, r.config.DBTableName)
	_, err := r.db.ExecContext(ctx, sql, messageID)
	return err
}

// GetUserMessageStats 获取用户消息统计
func (r *MessageRepository) GetUserMessageStats(ctx context.Context, userID string) (map[string]interface{}, error) {
	sql := fmt.Sprintf(`
		SELECT
			COUNT(*) as total,
			SUM(CASE WHEN delivered = TRUE THEN 1 ELSE 0 END) as delivered,
			SUM(CASE WHEN acked = TRUE THEN 1 ELSE 0 END) as acked,
			SUM(CASE WHEN created_at >= DATE_SUB(NOW(), INTERVAL 24 HOUR) THEN 1 ELSE 0 END) as last_24h
		FROM %s WHERE user_id = ?
	`, r.config.DBTableName)

	var total, delivered, acked, last24h int64
	err := r.db.QueryRowContext(ctx, sql, userID).Scan(&total, &delivered, &acked, &last24h)
	if err != nil {
		return nil, err
	}

	deliveryRate := float64(0)
	if total > 0 {
		deliveryRate = float64(delivered) / float64(total) * 100
	}

	return map[string]interface{}{
		"total":         total,
		"delivered":     delivered,
		"acked":         acked,
		"last_24h":      last24h,
		"delivery_rate": deliveryRate,
	}, nil
}

// ArchiveOldMessages 归档旧消息（删除或移动到冷存储）
func (r *MessageRepository) ArchiveOldMessages(ctx context.Context, before time.Time) (int64, error) {
	sql := fmt.Sprintf(`
		DELETE FROM %s WHERE created_at < ? AND acked = TRUE
	`, r.config.DBTableName)

	result, err := r.db.ExecContext(ctx, sql, before)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// RowScanner 接口，兼容 *sql.Row 和 *sql.Rows
type RowScanner interface {
	Scan(dest ...interface{}) error
}

// scanMessage 扫描消息行
func (r *MessageRepository) scanMessage(row RowScanner) (*model.Message, error) {
	var msg model.Message
	var msgType string
	var dataBytes []byte
	var templateID sql.NullString
	var expireAt sql.NullTime

	err := row.Scan(
		&msg.ID,
		&msg.UserID,
		&msg.DeviceID,
		&msgType,
		&msg.Priority,
		&templateID,
		&dataBytes,
		&msg.RequireAck,
		&msg.Timestamp,
		&expireAt,
	)

	if err != nil {
		return nil, err
	}

	msg.Type = model.MessageType(msgType)
	if templateID.Valid {
		msg.TemplateID = templateID.String
	}
	if expireAt.Valid {
		msg.ExpireAt = &expireAt.Time
	}

	// 解析 data JSON
	if len(dataBytes) > 0 {
		json.Unmarshal(dataBytes, &msg.Data)
	}

	return &msg, nil
}

// parseStreamMessage 解析 Redis Stream 消息
func (r *MessageRepository) parseStreamMessage(msg *redis.XMessage) *model.Message {
	data, ok := msg.Values["data"].(string)
	if !ok {
		return nil
	}

	var m model.Message
	if err := json.Unmarshal([]byte(data), &m); err != nil {
		return nil
	}

	return &m
}

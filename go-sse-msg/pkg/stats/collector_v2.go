package stats

import (
	"context"
	"encoding/json"
	"sync/atomic"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/redis/go-redis/v9"
)

// CollectorV2 优化版统计收集器
type CollectorV2 struct {
	redis         *redis.Client
	retentionDays int

	// Prometheus 指标
	messagesTotal     prometheus.Counter
	messagesSent      *prometheus.CounterVec
	messagesAcked     *prometheus.CounterVec
	messagesFailed    *prometheus.CounterVec
	activeConnections prometheus.Gauge
	latencyHistogram  prometheus.Histogram

	// 批量写入缓冲
	writeBuffer   chan *statsEvent
	batchSize     int
	flushInterval time.Duration

	// 内存统计 - 使用原子操作
	sentCount     int64
	deliveredCount int64
	failedCount   int64

	stopChan chan struct{}
}

// statsEvent 统计事件
type statsEvent struct {
	EventType string
	Params    map[string]string
	Timestamp time.Time
	Latency   time.Duration
}

// NewCollectorV2 创建优化版统计收集器
func NewCollectorV2(redis *redis.Client, retentionDays int) *CollectorV2 {
	c := &CollectorV2{
		redis:         redis,
		retentionDays: retentionDays,
		writeBuffer:   make(chan *statsEvent, 10000),
		batchSize:     100,
		flushInterval: 5 * time.Second,
		stopChan:      make(chan struct{}),

		messagesTotal: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "sse_messages_total",
			Help: "Total number of messages",
		}),
		messagesSent: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "sse_messages_sent_total",
			Help: "Total number of messages sent",
		}, []string{"type", "priority"}),
		messagesAcked: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "sse_messages_acked_total",
			Help: "Total number of messages acknowledged",
		}, []string{"type"}),
		messagesFailed: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "sse_messages_failed_total",
			Help: "Total number of failed messages",
		}, []string{"reason"}),
		activeConnections: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "sse_active_connections",
			Help: "Number of active SSE connections",
		}),
		latencyHistogram: prometheus.NewHistogram(prometheus.HistogramOpts{
			Name:    "sse_message_latency_seconds",
			Help:    "Message delivery latency",
			Buckets: prometheus.DefBuckets,
		}),
	}

	// 注册指标
	prometheus.MustRegister(
		c.messagesTotal,
		c.messagesSent,
		c.messagesAcked,
		c.messagesFailed,
		c.activeConnections,
		c.latencyHistogram,
	)

	// 启动批量写入协程
	go c.batchProcessor()

	return c
}

// batchProcessor 批量处理统计事件
func (c *CollectorV2) batchProcessor() {
	batch := make([]*statsEvent, 0, c.batchSize)
	timer := time.NewTimer(c.flushInterval)
	defer timer.Stop()

	for {
		select {
		case event := <-c.writeBuffer:
			batch = append(batch, event)
			if len(batch) >= c.batchSize {
				c.flushBatch(batch)
				batch = batch[:0]
				timer.Reset(c.flushInterval)
			}

		case <-timer.C:
			if len(batch) > 0 {
				c.flushBatch(batch)
				batch = batch[:0]
			}
			timer.Reset(c.flushInterval)

		case <-c.stopChan:
			if len(batch) > 0 {
				c.flushBatch(batch)
			}
			return
		}
	}
}

// flushBatch 批量写入 Redis
func (c *CollectorV2) flushBatch(batch []*statsEvent) {
	if len(batch) == 0 {
		return
	}

	ctx := context.Background()
	pipe := c.redis.Pipeline()

	date := time.Now().Format("2006-01-02")
	statsKey := "stats:daily:" + date

	for _, event := range batch {
		switch event.EventType {
		case "sent":
			pipe.HIncrBy(ctx, statsKey, "total_sent", 1)
		case "delivered":
			pipe.HIncrBy(ctx, statsKey, "total_delivered", 1)
		case "failed":
			pipe.HIncrBy(ctx, statsKey, "total_failed", 1)
		}
	}

	// 设置过期时间
	pipe.Expire(ctx, statsKey, time.Duration(c.retentionDays)*24*time.Hour)

	// 批量执行
	if _, err := pipe.Exec(ctx); err != nil {
		// 记录错误但不中断
	}
}

// RecordMessageSent 记录消息发送 - 异步
func (c *CollectorV2) RecordMessageSent(msgType, priority string) {
	c.messagesTotal.Inc()
	c.messagesSent.WithLabelValues(msgType, priority).Inc()
	atomic.AddInt64(&c.sentCount, 1)

	select {
	case c.writeBuffer <- &statsEvent{
		EventType: "sent",
		Params:    map[string]string{"type": msgType, "priority": priority},
		Timestamp: time.Now(),
	}:
	default:
		// 缓冲已满，丢弃事件
	}
}

// RecordMessageDelivered 记录消息送达 - 异步
func (c *CollectorV2) RecordMessageDelivered(msgType string, latency time.Duration) {
	c.messagesAcked.WithLabelValues(msgType).Inc()
	c.latencyHistogram.Observe(latency.Seconds())
	atomic.AddInt64(&c.deliveredCount, 1)

	select {
	case c.writeBuffer <- &statsEvent{
		EventType: "delivered",
		Params:    map[string]string{"type": msgType},
		Timestamp: time.Now(),
		Latency:   latency,
	}:
	default:
	}
}

// RecordMessageFailed 记录消息失败 - 异步
func (c *CollectorV2) RecordMessageFailed(reason string) {
	c.messagesFailed.WithLabelValues(reason).Inc()
	atomic.AddInt64(&c.failedCount, 1)

	select {
	case c.writeBuffer <- &statsEvent{
		EventType: "failed",
		Params:    map[string]string{"reason": reason},
		Timestamp: time.Now(),
	}:
	default:
	}
}

// SetActiveConnections 设置活跃连接数
func (c *CollectorV2) SetActiveConnections(count int) {
	c.activeConnections.Set(float64(count))
}

// RecordUserActive 记录活跃用户 - 批量处理
func (c *CollectorV2) RecordUserActive(userID string) {
	ctx := context.Background()
	date := time.Now().Format("2006-01-02")

	// 使用 Pipeline 批量执行
	pipe := c.redis.Pipeline()
	pipe.SAdd(ctx, "stats:users:"+date, userID)
	pipe.Expire(ctx, "stats:users:"+date, time.Duration(c.retentionDays)*24*time.Hour)
	pipe.Exec(ctx)
}

// GetDailyStats 获取每日统计 - 使用缓存
func (c *CollectorV2) GetDailyStats(date string) (*DailyStats, error) {
	ctx := context.Background()

	// 使用 Pipeline 获取所有数据
	pipe := c.redis.Pipeline()
	statsCmd := pipe.HGetAll(ctx, "stats:daily:"+date)
	usersCmd := pipe.SCard(ctx, "stats:users:"+date)

	_, err := pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}

	data, _ := statsCmd.Result()
	uniqueUsers, _ := usersCmd.Result()

	stats := &DailyStats{
		Date:        date,
		UniqueUsers: uniqueUsers,
	}

	// 解析数据
	if v, ok := data["total_sent"]; ok {
		json.Unmarshal([]byte(v), &stats.TotalSent)
	}
	if v, ok := data["total_delivered"]; ok {
		json.Unmarshal([]byte(v), &stats.TotalDelivered)
	}
	if v, ok := data["total_failed"]; ok {
		json.Unmarshal([]byte(v), &stats.TotalFailed)
	}
	if v, ok := data["avg_latency_ms"]; ok {
		json.Unmarshal([]byte(v), &stats.AvgLatencyMs)
	}

	return stats, nil
}

// GetRealtimeStats 获取实时统计 - 从内存读取
func (c *CollectorV2) GetRealtimeStats() map[string]interface{} {
	return map[string]interface{}{
		"active_connections": c.activeConnections,
		"sent_count":         atomic.LoadInt64(&c.sentCount),
		"delivered_count":    atomic.LoadInt64(&c.deliveredCount),
		"failed_count":       atomic.LoadInt64(&c.failedCount),
	}
}

// Stop 停止收集器
func (c *CollectorV2) Stop() {
	close(c.stopChan)
}

// DailyStatsV2 每日统计（别名避免重复定义）
type DailyStatsV2 = DailyStats

package stats

import (
	"context"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/redis/go-redis/v9"
)

// Collector 推送统计收集器
type Collector struct {
	redis        *redis.Client
	retentionDays int

	// Prometheus 指标
	messagesTotal   prometheus.Counter
	messagesSent    *prometheus.CounterVec
	messagesAcked   *prometheus.CounterVec
	messagesFailed  *prometheus.CounterVec
	activeConnections prometheus.Gauge
	latencyHistogram  prometheus.Histogram

	// 内存缓存
	mu              sync.RWMutex
	dailyStats      map[string]*DailyStats
}

// DailyStats 每日统计
type DailyStats struct {
	Date           string  `json:"date"`
	TotalSent      int64   `json:"total_sent"`
	TotalDelivered int64   `json:"total_delivered"`
	TotalAcked     int64   `json:"total_acked"`
	TotalFailed    int64   `json:"total_failed"`
	UniqueUsers    int64   `json:"unique_users"`
	AvgLatencyMs   float64 `json:"avg_latency_ms"`
}

// NewCollector 创建统计收集器
func NewCollector(redis *redis.Client, retentionDays int) *Collector {
	c := &Collector{
		redis:         redis,
		retentionDays: retentionDays,
		dailyStats:    make(map[string]*DailyStats),

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

	// 启动定时保存
	go c.saveLoop()

	return c
}

// RecordMessageSent 记录消息发送
func (c *Collector) RecordMessageSent(msgType, priority string) {
	c.messagesTotal.Inc()
	c.messagesSent.WithLabelValues(msgType, priority).Inc()

	date := time.Now().Format("2006-01-02")
	c.mu.Lock()
	defer c.mu.Unlock()

	stats, ok := c.dailyStats[date]
	if !ok {
		stats = &DailyStats{Date: date}
		c.dailyStats[date] = stats
	}
	stats.TotalSent++

	// 写入 Redis
	ctx := context.Background()
	c.redis.HIncrBy(ctx, "stats:daily:"+date, "total_sent", 1)
}

// RecordMessageDelivered 记录消息送达
func (c *Collector) RecordMessageDelivered(msgType string, latency time.Duration) {
	c.messagesAcked.WithLabelValues(msgType).Inc()
	c.latencyHistogram.Observe(latency.Seconds())

	date := time.Now().Format("2006-01-02")
	c.mu.Lock()
	defer c.mu.Unlock()

	stats, ok := c.dailyStats[date]
	if !ok {
		stats = &DailyStats{Date: date}
		c.dailyStats[date] = stats
	}
	stats.TotalDelivered++

	// 更新平均延迟
	if stats.TotalDelivered == 1 {
		stats.AvgLatencyMs = float64(latency.Milliseconds())
	} else {
		stats.AvgLatencyMs = (stats.AvgLatencyMs*float64(stats.TotalDelivered-1) + float64(latency.Milliseconds())) / float64(stats.TotalDelivered)
	}

	ctx := context.Background()
	c.redis.HIncrBy(ctx, "stats:daily:"+date, "total_delivered", 1)
}

// RecordMessageFailed 记录消息失败
func (c *Collector) RecordMessageFailed(reason string) {
	c.messagesFailed.WithLabelValues(reason).Inc()

	date := time.Now().Format("2006-01-02")
	c.mu.Lock()
	defer c.mu.Unlock()

	stats, ok := c.dailyStats[date]
	if !ok {
		stats = &DailyStats{Date: date}
		c.dailyStats[date] = stats
	}
	stats.TotalFailed++

	ctx := context.Background()
	c.redis.HIncrBy(ctx, "stats:daily:"+date, "total_failed", 1)
}

// SetActiveConnections 设置活跃连接数
func (c *Collector) SetActiveConnections(count int) {
	c.activeConnections.Set(float64(count))
}

// RecordUserActive 记录活跃用户
func (c *Collector) RecordUserActive(userID string) {
	ctx := context.Background()
	date := time.Now().Format("2006-01-02")
	c.redis.SAdd(ctx, "stats:users:"+date, userID)
	c.redis.Expire(ctx, "stats:users:"+date, time.Duration(c.retentionDays)*24*time.Hour)
}

// GetDailyStats 获取每日统计
func (c *Collector) GetDailyStats(date string) (*DailyStats, error) {
	ctx := context.Background()
	data, err := c.redis.HGetAll(ctx, "stats:daily:"+date).Result()
	if err != nil {
		return nil, err
	}

	stats := &DailyStats{Date: date}
	// 解析数据...
	_ = data

	// 获取独立用户数
	uniqueUsers, _ := c.redis.SCard(ctx, "stats:users:"+date).Result()
	stats.UniqueUsers = uniqueUsers

	return stats, nil
}

// GetRealtimeStats 获取实时统计
func (c *Collector) GetRealtimeStats() map[string]interface{} {
	return map[string]interface{}{
		"active_connections": c.activeConnections,
	}
}

// saveLoop 定时保存统计
func (c *Collector) saveLoop() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		c.saveToRedis()
	}
}

// saveToRedis 保存到 Redis
func (c *Collector) saveToRedis() {
	c.mu.RLock()
	defer c.mu.RUnlock()

	ctx := context.Background()
	for date, stats := range c.dailyStats {
		key := "stats:daily:" + date
		c.redis.HSet(ctx, key, "total_sent", stats.TotalSent)
		c.redis.HSet(ctx, key, "total_delivered", stats.TotalDelivered)
		c.redis.HSet(ctx, key, "total_acked", stats.TotalAcked)
		c.redis.HSet(ctx, key, "total_failed", stats.TotalFailed)
		c.redis.HSet(ctx, key, "avg_latency_ms", stats.AvgLatencyMs)
		c.redis.Expire(ctx, key, time.Duration(c.retentionDays)*24*time.Hour)
	}
}

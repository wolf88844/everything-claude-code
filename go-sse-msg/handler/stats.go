package handler

import (
	"net/http"
	"time"

	"go-sse-msg/pkg/stats"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// StatsHandler 统计处理器
type StatsHandler struct {
	statsCollector *stats.Collector
}

// NewStatsHandler 创建统计处理器
func NewStatsHandler(statsCollector *stats.Collector) *StatsHandler {
	return &StatsHandler{statsCollector: statsCollector}
}

// GetDailyStats 获取每日统计
// GET /api/stats/daily/:date
func (h *StatsHandler) GetDailyStats(c *gin.Context) {
	date := c.Param("date")
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	s, err := h.statsCollector.GetDailyStats(date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, s)
}

// GetRealtimeStats 获取实时统计
// GET /api/stats/realtime
func (h *StatsHandler) GetRealtimeStats(c *gin.Context) {
	stats := h.statsCollector.GetRealtimeStats()
	c.JSON(http.StatusOK, stats)
}

// PrometheusMetrics Prometheus 指标端点
// GET /metrics
func (h *StatsHandler) PrometheusMetrics() gin.HandlerFunc {
	handler := promhttp.Handler()
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

package handler

import (
	"net/http"

	"go-sse-msg/service"

	"github.com/gin-gonic/gin"
)

// PresenceHandler 在线状态处理器
type PresenceHandler struct {
	presenceService *service.PresenceService
}

// NewPresenceHandler 创建在线状态处理器
func NewPresenceHandler(presenceService *service.PresenceService) *PresenceHandler {
	return &PresenceHandler{presenceService: presenceService}
}

// GetMyStatus 获取自己的在线状态
// GET /api/presence/me
func (h *PresenceHandler) GetMyStatus(c *gin.Context) {
	userID := c.GetString("userID")

	status, err := h.presenceService.GetOnlineStatus(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, status)
}

// GetUserStatus 获取指定用户在线状态
// GET /api/presence/user/:user_id
func (h *PresenceHandler) GetUserStatus(c *gin.Context) {
	userID := c.Param("user_id")

	status, err := h.presenceService.GetOnlineStatus(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, status)
}

// GetUserDevices 获取用户的设备列表
// GET /api/presence/user/:user_id/devices
func (h *PresenceHandler) GetUserDevices(c *gin.Context) {
	userID := c.Param("user_id")

	devices, err := h.presenceService.GetUserDevices(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"devices": devices,
		"count":   len(devices),
	})
}

// Heartbeat 心跳接口
// POST /api/presence/heartbeat
func (h *PresenceHandler) Heartbeat(c *gin.Context) {
	deviceID := c.GetString("deviceID")

	if err := h.presenceService.UpdateHeartbeat(deviceID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// IsOnline 检查用户是否在线
// GET /api/presence/user/:user_id/online
func (h *PresenceHandler) IsOnline(c *gin.Context) {
	userID := c.Param("user_id")
	online := h.presenceService.IsUserOnline(userID)

	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"online":  online,
	})
}

package handler

import (
	"net/http"

	"go-sse-msg/model"
	"go-sse-msg/service"

	"github.com/gin-gonic/gin"
)

// SSEHandler SSE 处理器
type SSEHandler struct {
	sseService service.SSEServiceInterface
}

// NewSSEHandler 创建 SSE 处理器
func NewSSEHandler(sseService service.SSEServiceInterface) *SSEHandler {
	return &SSEHandler{sseService: sseService}
}

// Subscribe SSE 订阅端点
// GET /events
func (h *SSEHandler) Subscribe(c *gin.Context) {
	userID := c.GetString("userID")
	deviceID := c.GetString("deviceID")
	deviceType := c.GetString("deviceType")

	// 获取 last_sync_id 用于断线恢复
	lastSyncID := c.GetHeader("X-Last-Sync-ID")
	if lastSyncID != "" {
		// 执行消息同步
		// h.sseService.SyncMessages(...)
	}

	h.sseService.Subscribe(c, userID, deviceID, deviceType)
}

// SendMessage 发送消息
// POST /api/messages/send
func (h *SSEHandler) SendMessage(c *gin.Context) {
	var req struct {
		UserID     string      `json:"user_id" binding:"required"`
		DeviceID   string      `json:"device_id,omitempty"`
		Type       string      `json:"type" binding:"required"`
		Priority   int         `json:"priority,omitempty"`
		Data       interface{} `json:"data" binding:"required"`
		RequireAck bool        `json:"require_ack,omitempty"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg := &model.Message{
		UserID:     req.UserID,
		DeviceID:   req.DeviceID,
		Type:       model.MessageType(req.Type),
		Priority:   model.MessagePriority(req.Priority),
		Data:       req.Data,
		RequireAck: req.RequireAck,
	}

	if err := h.sseService.SendMessage(msg); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message_id": msg.ID,
		"status":     "sent",
	})
}

// Broadcast 广播消息
// POST /api/messages/broadcast
func (h *SSEHandler) Broadcast(c *gin.Context) {
	var req model.BroadcastRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.sseService.Broadcast(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "broadcasting"})
}

// SendWithTemplate 使用模板发送
// POST /api/messages/template
func (h *SSEHandler) SendWithTemplate(c *gin.Context) {
	var req struct {
		UserID     string                 `json:"user_id" binding:"required"`
		TemplateID string                 `json:"template_id" binding:"required"`
		Language   string                 `json:"language,omitempty"`
		Variables  map[string]interface{} `json:"variables"`
		RequireAck bool                   `json:"require_ack,omitempty"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.sseService.SendWithTemplate(req.UserID, req.TemplateID, req.Language, req.Variables, req.RequireAck); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "sent"})
}

// Ack 确认消息
// POST /api/messages/ack
func (h *SSEHandler) Ack(c *gin.Context) {
	var req struct {
		MessageID string `json:"message_id" binding:"required"`
		UserID    string `json:"user_id,omitempty"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 如果没有提供 userID，从上下文获取
	if req.UserID == "" {
		req.UserID = c.GetString("userID")
	}

	if ok := h.sseService.ProcessACK(req.MessageID, req.UserID); !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "message not found or already acked"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "acknowledged"})
}

// GetPendingAcks 获取待确认消息
// GET /api/messages/pending-acks
func (h *SSEHandler) GetPendingAcks(c *gin.Context) {
	userID := c.GetString("userID")
	list := h.sseService.GetAckManager().GetPendingList(userID)

	c.JSON(http.StatusOK, gin.H{
		"pending_acks": list,
		"count":        len(list),
	})
}

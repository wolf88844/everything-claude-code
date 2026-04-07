package handler

import (
	"net/http"
	"strconv"

	"go-sse-msg/repository"

	"github.com/gin-gonic/gin"
)

// MessageHandler 消息历史处理器
type MessageHandler struct {
	msgRepo *repository.MessageRepository
}

// NewMessageHandler 创建消息处理器
func NewMessageHandler(msgRepo *repository.MessageRepository) *MessageHandler {
	return &MessageHandler{msgRepo: msgRepo}
}

// GetMessageHistory 获取消息历史
// GET /api/messages/history
func (h *MessageHandler) GetMessageHistory(c *gin.Context) {
	userID := c.GetString("userID")

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	if limit > 100 {
		limit = 100
	}

	ctx := c.Request.Context()
	messages, err := h.msgRepo.GetByUser(ctx, userID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messages": messages,
		"limit":    limit,
		"offset":   offset,
		"count":    len(messages),
	})
}

// GetMessageByID 获取单条消息详情
// GET /api/messages/:id
func (h *MessageHandler) GetMessageByID(c *gin.Context) {
	messageID := c.Param("id")

	ctx := c.Request.Context()
	msg, err := h.msgRepo.GetByID(ctx, messageID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "message not found"})
		return
	}

	c.JSON(http.StatusOK, msg)
}

// GetUnreadMessages 获取未读消息（断线恢复）
// GET /api/messages/unread
func (h *MessageHandler) GetUnreadMessages(c *gin.Context) {
	userID := c.GetString("userID")
	lastReadID := c.Query("last_id")

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))
	if limit > 500 {
		limit = 500
	}

	ctx := c.Request.Context()
	messages, err := h.msgRepo.GetUnreadMessages(ctx, userID, lastReadID, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messages": messages,
		"count":    len(messages),
	})
}

// GetMessageStats 获取消息统计
// GET /api/messages/stats
func (h *MessageHandler) GetMessageStats(c *gin.Context) {
	userID := c.GetString("userID")

	ctx := c.Request.Context()
	stats, err := h.msgRepo.GetUserMessageStats(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stats)
}

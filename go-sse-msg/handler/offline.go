package handler

import (
	"net/http"
	"strconv"

	"go-sse-msg/pkg/offline"

	"github.com/gin-gonic/gin"
)

// OfflineHandler 离线消息处理器
type OfflineHandler struct {
	offlineManager *offline.Manager
}

// NewOfflineHandler 创建离线消息处理器
func NewOfflineHandler(offlineManager *offline.Manager) *OfflineHandler {
	return &OfflineHandler{offlineManager: offlineManager}
}

// GetOfflineMessages 获取离线消息
// GET /api/offline/messages
func (h *OfflineHandler) GetOfflineMessages(c *gin.Context) {
	userID := c.GetString("userID")
	limitStr := c.DefaultQuery("limit", "100")
	limit, _ := strconv.ParseInt(limitStr, 10, 64)

	messages, err := h.offlineManager.Fetch(userID, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messages": messages,
		"count":    len(messages),
	})
}

// PeekOfflineMessages 预览离线消息（不删除）
// GET /api/offline/messages/peek
func (h *OfflineHandler) PeekOfflineMessages(c *gin.Context) {
	userID := c.GetString("userID")
	limitStr := c.DefaultQuery("limit", "100")
	limit, _ := strconv.ParseInt(limitStr, 10, 64)

	messages, err := h.offlineManager.Peek(userID, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messages": messages,
		"count":    len(messages),
	})
}

// GetOfflineCount 获取离线消息数量
// GET /api/offline/count
func (h *OfflineHandler) GetOfflineCount(c *gin.Context) {
	userID := c.GetString("userID")
	count := h.offlineManager.GetCount(userID)

	c.JSON(http.StatusOK, gin.H{
		"count": count,
	})
}

// ClearOfflineMessages 清空离线消息
// DELETE /api/offline/messages
func (h *OfflineHandler) ClearOfflineMessages(c *gin.Context) {
	userID := c.GetString("userID")

	if err := h.offlineManager.Clear(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "cleared"})
}

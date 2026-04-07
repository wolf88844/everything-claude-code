package service

import (
	"go-sse-msg/model"
	"go-sse-msg/pkg/ack"
	"go-sse-msg/pkg/offline"
	"go-sse-msg/pkg/stats"
	"go-sse-msg/repository"

	"github.com/gin-gonic/gin"
)

// SSEServiceInterface SSE 服务接口
type SSEServiceInterface interface {
	Subscribe(c *gin.Context, userID, deviceID, deviceType string)
	SendMessage(msg *model.Message) error
	Broadcast(req *model.BroadcastRequest) error
	SendWithTemplate(userID, templateID, lang string, variables map[string]interface{}, requireAck bool) error
	ProcessACK(messageID, userID string) bool
	SyncMessages(req *model.SyncRequest) (*model.SyncResponse, error)
	GetPresenceService() *PresenceService
	GetStatsCollector() *stats.Collector
	GetOfflineManager() *offline.Manager
	GetAckManager() *ack.Manager
	GetMessageRepository() *repository.MessageRepository
}

// 确保两个版本都实现接口
var _ SSEServiceInterface = (*SSEService)(nil)
var _ SSEServiceInterface = (*SSEServiceV2)(nil)

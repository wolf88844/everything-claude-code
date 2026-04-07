package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-sse-msg/config"
	"go-sse-msg/handler"
	"go-sse-msg/middleware"
	"go-sse-msg/repository"
	"go-sse-msg/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 初始化 Redis 客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	// 测试 Redis 连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis")

	// 初始化 MySQL 连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySQL.User,
		cfg.MySQL.Password,
		cfg.MySQL.Host,
		cfg.MySQL.Port,
		cfg.MySQL.Database,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open MySQL connection: %v", err)
	}
	defer db.Close()

	// 配置连接池
	db.SetMaxOpenConns(cfg.MySQL.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MySQL.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.MySQL.ConnMaxLifetime)

	// 测试连接
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	log.Println("Connected to MySQL")

	// 初始化消息仓库（如果 MySQL 连接成功）
	var msgRepo *repository.MessageRepository
	if err := db.PingContext(ctx); err == nil {
		msgRepo = repository.NewMessageRepository(db, redisClient, nil)
		log.Println("Message persistence enabled")
	} else {
		log.Printf("MySQL not available, running without persistence: %v", err)
	}

	// 初始化服务（使用优化版本）
	sseService := service.NewSSEServiceV2(cfg, redisClient, msgRepo)

	// 初始化处理器
	sseHandler := handler.NewSSEHandler(sseService)
	presenceHandler := handler.NewPresenceHandler(sseService.GetPresenceService())
	statsHandler := handler.NewStatsHandler(sseService.GetStatsCollector())
	offlineHandler := handler.NewOfflineHandler(sseService.GetOfflineManager())
	messageHandler := handler.NewMessageHandler(msgRepo)

	// 设置 Gin 路由
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())
	r.Use(middleware.Security())
	r.Use(middleware.Logger())

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"time":   time.Now().Format(time.RFC3339),
		})
	})

	// SSE 订阅端点
	r.GET("/events", middleware.Auth(), sseHandler.Subscribe)

	// API 路由组
	api := r.Group("/api")
	api.Use(middleware.Auth())
	{
		// 消息相关
		msg := api.Group("/messages")
		{
			msg.POST("/send", sseHandler.SendMessage)
			msg.POST("/broadcast", sseHandler.Broadcast)
			msg.POST("/template", sseHandler.SendWithTemplate)
			msg.POST("/ack", sseHandler.Ack)
			msg.GET("/pending-acks", sseHandler.GetPendingAcks)
			msg.GET("/history", messageHandler.GetMessageHistory)
			msg.GET("/unread", messageHandler.GetUnreadMessages)
			msg.GET("/stats", messageHandler.GetMessageStats)
			msg.GET("/:id", messageHandler.GetMessageByID)
		}

		// 在线状态
		presence := api.Group("/presence")
		{
			presence.GET("/me", presenceHandler.GetMyStatus)
			presence.GET("/user/:user_id", presenceHandler.GetUserStatus)
			presence.GET("/user/:user_id/devices", presenceHandler.GetUserDevices)
			presence.GET("/user/:user_id/online", presenceHandler.IsOnline)
			presence.POST("/heartbeat", presenceHandler.Heartbeat)
		}

		// 离线消息
		offline := api.Group("/offline")
		{
			offline.GET("/messages", offlineHandler.GetOfflineMessages)
			offline.GET("/messages/peek", offlineHandler.PeekOfflineMessages)
			offline.GET("/count", offlineHandler.GetOfflineCount)
			offline.DELETE("/messages", offlineHandler.ClearOfflineMessages)
		}

		// 统计
		stats := api.Group("/stats")
		{
			stats.GET("/daily/:date", statsHandler.GetDailyStats)
			stats.GET("/realtime", statsHandler.GetRealtimeStats)
		}
	}

	// Prometheus 指标
	r.GET("/metrics", statsHandler.PrometheusMetrics())

	// 启动 HTTP 服务器
	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      r,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	// 优雅关闭
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	log.Printf("Server started on port %s", cfg.Server.Port)

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// 优雅关闭
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	// 关闭 Redis 连接
	if err := redisClient.Close(); err != nil {
		log.Printf("Failed to close Redis connection: %v", err)
	}

	log.Println("Server exited")
}

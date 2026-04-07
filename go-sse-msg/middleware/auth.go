package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Auth 认证中间件（简化版，实际项目应使用 JWT 等）
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 Header 获取用户信息
		// 实际项目应该验证 JWT Token
		userID := c.GetHeader("X-User-ID")
		if userID == "" {
			// 开发阶段允许匿名访问，使用 IP 作为临时 ID
			userID = c.ClientIP()
		}

		deviceID := c.GetHeader("X-Device-ID")
		if deviceID == "" {
			deviceID = "unknown"
		}

		deviceType := c.GetHeader("X-Device-Type")
		if deviceType == "" {
			// 根据 User-Agent 推断
			ua := c.Request.UserAgent()
			if strings.Contains(ua, "Mobile") {
				deviceType = "mobile"
			} else {
				deviceType = "web"
			}
		}

		c.Set("userID", userID)
		c.Set("deviceID", deviceID)
		c.Set("deviceType", deviceType)

		c.Next()
	}
}

// RequireAuth 强制认证中间件
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		// 这里应该验证 JWT Token
		// 简化处理：假设 token 就是 userID
		c.Set("userID", strings.TrimPrefix(token, "Bearer "))
		c.Next()
	}
}

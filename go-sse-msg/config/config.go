package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Server     ServerConfig
	Redis      RedisConfig
	MySQL      MySQLConfig
	SSE        SSEConfig
	Template   TemplateConfig
	Stats      StatsConfig
}

type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type RedisConfig struct {
	Addr         string
	Password     string
	DB           int
	StreamName   string
	OfflineQueue string
}

type SSEConfig struct {
	MaxConnections    int
	HeartbeatInterval time.Duration
	AckTimeout        time.Duration
	MessageBufferSize int
}

type MySQLConfig struct {
	Host            string
	Port            string
	Database        string
	User            string
	Password        string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type TemplateConfig struct {
	MaxTemplates int
}

type StatsConfig struct {
	RetentionDays int
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port:         getEnv("SERVER_PORT", "8080"),
			ReadTimeout:  getDuration("SERVER_READ_TIMEOUT", 10*time.Second),
			WriteTimeout: getDuration("SERVER_WRITE_TIMEOUT", 10*time.Second),
		},
		Redis: RedisConfig{
			Addr:         getEnv("REDIS_ADDR", "localhost:6379"),
			Password:     getEnv("REDIS_PASSWORD", ""),
			DB:           getInt("REDIS_DB", 0),
			StreamName:   getEnv("REDIS_STREAM_NAME", "sse:messages"),
			OfflineQueue: getEnv("REDIS_OFFLINE_QUEUE", "sse:offline:"),
		},
		MySQL: MySQLConfig{
			Host:            getEnv("MYSQL_HOST", "localhost"),
			Port:            getEnv("MYSQL_PORT", "3306"),
			Database:        getEnv("MYSQL_DATABASE", "sse_msg"),
			User:            getEnv("MYSQL_USER", "root"),
			Password:        getEnv("MYSQL_PASSWORD", ""),
			MaxOpenConns:    getInt("MYSQL_MAX_OPEN_CONNS", 25),
			MaxIdleConns:    getInt("MYSQL_MAX_IDLE_CONNS", 10),
			ConnMaxLifetime: getDuration("MYSQL_CONN_MAX_LIFETIME", 5*time.Minute),
		},
		SSE: SSEConfig{
			MaxConnections:    getInt("SSE_MAX_CONNECTIONS", 10000),
			HeartbeatInterval: getDuration("SSE_HEARTBEAT", 30*time.Second),
			AckTimeout:        getDuration("SSE_ACK_TIMEOUT", 5*time.Second),
			MessageBufferSize: getInt("SSE_BUFFER_SIZE", 100),
		},
		Template: TemplateConfig{
			MaxTemplates: getInt("TEMPLATE_MAX", 100),
		},
		Stats: StatsConfig{
			RetentionDays: getInt("STATS_RETENTION_DAYS", 7),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
	}
	return defaultValue
}

func getDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if d, err := time.ParseDuration(value); err == nil {
			return d
		}
	}
	return defaultValue
}

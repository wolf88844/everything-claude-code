# Go SSE 消息推送系统

企业级 Server-Sent Events 实时消息推送服务，支持 ACK 确认、用户在线状态、多端同步、消息模板、推送统计、离线消息等完整功能。

## 功能特性

| 功能 | 说明 |
|------|------|
| **实时推送** | 基于 SSE 的服务端主动推送 |
| **ACK 确认** | 消息送达确认，超时自动重试 |
| **用户在线状态** | 实时追踪用户在线/离线状态 |
| **多端同步** | 同一账号多设备消息同步 |
| **消息模板** | 支持模板渲染、国际化、变量替换 |
| **推送统计** | Prometheus 指标 + 每日统计报表 |
| **离线消息** | 用户离线时存储，上线后自动推送 |
| **消息持久化** | MySQL 长期存储 + Redis Stream 实时流，支持消息查询和归档 |
| **Redis 集群** | 支持多实例部署，消息跨节点广播 |

## 项目结构

```
go-sse-msg/
├── cmd/
│   └── main.go              # 入口文件
├── config/
│   └── config.go            # 配置管理
├── handler/                 # HTTP 处理器
│   ├── sse.go               # SSE 相关接口
│   ├── presence.go          # 在线状态接口
│   ├── stats.go             # 统计接口
│   ├── offline.go           # 离线消息接口
│   └── message.go           # 消息历史接口
├── repository/              # 数据持久层
│   └── message.go           # MySQL + Redis Stream 双写
├── middleware/              # 中间件
│   ├── cors.go              # 跨域、安全中间件
│   └── auth.go              # 认证中间件
├── model/                   # 数据模型
│   ├── message.go           # 消息、ACK、设备模型
│   └── template.go          # 消息模板模型
├── pkg/                     # 核心包
│   ├── ack/                 # ACK 管理
│   ├── offline/             # 离线消息管理
│   └── stats/               # 推送统计
├── service/                 # 业务服务
│   ├── sse.go               # SSE 核心服务
│   └── presence.go          # 在线状态服务
├── go.mod
└── README.md
```

## 快速开始

### 1. 安装依赖

 ```bash
 # 安装 Redis (Mac)
 brew install redis
 brew services start redis

 # 安装 Redis (Ubuntu)
 sudo apt-get install redis-server
 sudo systemctl start redis

 # 安装 MySQL (Mac)
 brew install mysql
 brew services start mysql

 # 安装 MySQL (Ubuntu)
 sudo apt-get install mysql-server
 sudo systemctl start mysql

 # 创建数据库
 mysql -u root -e "CREATE DATABASE IF NOT EXISTS sse_msg;"

 # 安装 Go 依赖
 go mod download
 ```

### 2. 启动服务

```bash
go run cmd/main.go
```

服务启动在 `http://localhost:8080`

### 3. 测试 SSE 连接

```bash
# 使用 curl 订阅 SSE
curl -N "http://localhost:8080/events" \
  -H "X-User-ID: user123" \
  -H "X-Device-ID: device001" \
  -H "X-Device-Type: web"
```

### 4. 发送消息

```bash
# 发送普通消息
curl -X POST "http://localhost:8080/api/messages/send" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user123",
    "type": "notification",
    "data": {"title": "Hello", "content": "World"},
    "require_ack": true
  }'

# 使用模板发送
curl -X POST "http://localhost:8080/api/messages/template" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user123",
    "template_id": "order_created",
    "variables": {"OrderID": "ORD-001", "Amount": "99.99"}
  }'
```

### 5. 确认消息 (ACK)

```bash
curl -X POST "http://localhost:8080/api/messages/ack" \
  -H "Content-Type: application/json" \
  -d '{"message_id": "xxx", "user_id": "user123"}'
```

## API 文档

### SSE 端点

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/events` | SSE 订阅，保持连接接收实时消息 |

**请求头:**
- `X-User-ID`: 用户 ID
- `X-Device-ID`: 设备 ID
- `X-Device-Type`: 设备类型 (web, ios, android, desktop)
- `X-Last-Sync-ID`: 断线恢复用的最后同步 ID

### 消息接口

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/api/messages/send` | 发送消息给指定用户 |
| POST | `/api/messages/broadcast` | 广播消息给所有用户 |
| POST | `/api/messages/template` | 使用模板发送消息 |
| POST | `/api/messages/ack` | 确认收到消息 |
| GET | `/api/messages/pending-acks` | 获取待确认消息列表 |
| GET | `/api/messages/history` | 获取消息历史（分页） |
| GET | `/api/messages/unread` | 获取未读消息（断线恢复） |
| GET | `/api/messages/stats` | 获取消息统计 |
| GET | `/api/messages/:id` | 获取单条消息详情 |

### 在线状态接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/api/presence/me` | 获取自己的在线状态 |
| GET | `/api/presence/user/:user_id` | 获取指定用户状态 |
| GET | `/api/presence/user/:user_id/devices` | 获取用户设备列表 |
| GET | `/api/presence/user/:user_id/online` | 检查用户是否在线 |
| POST | `/api/presence/heartbeat` | 发送心跳 |

### 离线消息接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/api/offline/messages` | 获取并清空离线消息 |
| GET | `/api/offline/messages/peek` | 预览离线消息（不清空） |
| GET | `/api/offline/count` | 获取离线消息数量 |
| DELETE | `/api/offline/messages` | 清空离线消息 |

### 统计接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/api/stats/daily/:date` | 获取每日统计 |
| GET | `/api/stats/realtime` | 获取实时统计 |
| GET | `/metrics` | Prometheus 指标 |

## 配置

环境变量配置:

| 变量 | 默认值 | 说明 |
|------|--------|------|
| `SERVER_PORT` | 8080 | HTTP 服务端口 |
| `REDIS_ADDR` | localhost:6379 | Redis 地址 |
| `REDIS_PASSWORD` | - | Redis 密码 |
| `REDIS_DB` | 0 | Redis 数据库 |
| `MYSQL_HOST` | localhost | MySQL 主机 |
| `MYSQL_PORT` | 3306 | MySQL 端口 |
| `MYSQL_DATABASE` | sse_msg | MySQL 数据库名 |
| `MYSQL_USER` | root | MySQL 用户名 |
| `MYSQL_PASSWORD` | - | MySQL 密码 |
| `SSE_MAX_CONNECTIONS` | 10000 | 最大 SSE 连接数 |
| `SSE_ACK_TIMEOUT` | 5s | ACK 超时时间 |
| `STATS_RETENTION_DAYS` | 7 | 统计数据保留天数 |

## 内存安全设计

项目针对 SSE 常见内存泄漏问题做了以下处理:

1. **连接自动清理**: 客户端断开时自动清理 map 和 channel
2. **心跳检测**: 定期检测过期连接并释放资源
3. **Channel 缓冲**: 带缓冲 channel 防止发送阻塞
4. **Context 管理**: 使用 context 管理 goroutine 生命周期
5. **ACK 超时**: 超时未确认的消息自动重试，超过次数后清理

## 多实例部署

服务支持多实例部署，通过 Redis Pub/Sub 实现消息广播:

```yaml
# docker-compose.yml
version: '3'
services:
  app1:
    build: .
    ports:
      - "8081:8080"
    environment:
      - REDIS_ADDR=redis:6379
  
  app2:
    build: .
    ports:
      - "8082:8080"
    environment:
      - REDIS_ADDR=redis:6379
  
  redis:
    image: redis:alpine
```

## 预设消息模板

系统内置以下模板:

| 模板 ID | 说明 | 变量 |
|---------|------|------|
| `order_created` | 订单创建通知 | `OrderID`, `Amount` |
| `payment_success` | 支付成功通知 | `OrderID`, `Amount` |
| `system_alert` | 系统警告 | `Level`, `Message`, `Time` |

## 监控

Prometheus 指标:

```
sse_messages_total              # 总消息数
sse_messages_sent_total         # 发送消息数（按类型、优先级）
sse_messages_acked_total        # 确认消息数
sse_messages_failed_total       # 失败消息数
sse_active_connections          # 活跃连接数
sse_message_latency_seconds     # 消息延迟分布
```

## License

MIT

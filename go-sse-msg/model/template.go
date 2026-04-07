package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"
	"text/template"
	"time"
)

// MessageTemplate 消息模板
type MessageTemplate struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Type        MessageType       `json:"type"`
	Priority    MessagePriority   `json:"priority"`
	Content     TemplateContent   `json:"content"`
	Variables   []TemplateVar     `json:"variables"`
	Version     int               `json:"version"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	Enabled     bool              `json:"enabled"`
}

// TemplateContent 模板内容（多语言）
type TemplateContent struct {
	Default string            `json:"default"`
	I18n    map[string]string `json:"i18n,omitempty"` // zh-CN, en-US, etc.
}

// TemplateVar 模板变量定义
type TemplateVar struct {
	Name        string `json:"name"`
	Type        string `json:"type"`        // string, number, bool, date
	Required    bool   `json:"required"`
	Default     string `json:"default"`
	Description string `json:"description"`
}

// RenderedMessage 渲染后的消息
type RenderedMessage struct {
	ID         string          `json:"id"`
	TemplateID string          `json:"template_id"`
	Type       MessageType     `json:"type"`
	Priority   MessagePriority `json:"priority"`
	Title      string          `json:"title,omitempty"`
	Content    string          `json:"content"`
	Data       interface{}     `json:"data"`
}

// TemplateManager 模板管理器
type TemplateManager struct {
	templates map[string]*MessageTemplate
	cache     map[string]*template.Template
	mu        sync.RWMutex
	maxSize   int
}

// NewTemplateManager 创建模板管理器
func NewTemplateManager(maxSize int) *TemplateManager {
	return &TemplateManager{
		templates: make(map[string]*MessageTemplate),
		cache:     make(map[string]*template.Template),
		maxSize:   maxSize,
	}
}

// Register 注册模板
func (tm *TemplateManager) Register(t *MessageTemplate) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if len(tm.templates) >= tm.maxSize {
		return fmt.Errorf("template manager reached max capacity: %d", tm.maxSize)
	}

	// 预编译模板
	tpl, err := template.New(t.ID).Parse(t.Content.Default)
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", t.ID, err)
	}

	tm.templates[t.ID] = t
	tm.cache[t.ID] = tpl

	// 编译国际化模板
	for lang, content := range t.Content.I18n {
		key := fmt.Sprintf("%s:%s", t.ID, lang)
		tpl, err := template.New(key).Parse(content)
		if err != nil {
			return fmt.Errorf("failed to parse template %s: %w", key, err)
		}
		tm.cache[key] = tpl
	}

	return nil
}

// Get 获取模板
func (tm *TemplateManager) Get(id string) (*MessageTemplate, bool) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	t, ok := tm.templates[id]
	return t, ok
}

// Render 渲染模板
func (tm *TemplateManager) Render(templateID string, lang string, variables map[string]interface{}) (*RenderedMessage, error) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	t, ok := tm.templates[templateID]
	if !ok {
		return nil, fmt.Errorf("template not found: %s", templateID)
	}

	if !t.Enabled {
		return nil, fmt.Errorf("template is disabled: %s", templateID)
	}

	// 获取对应语言的模板
	key := templateID
	if lang != "" && lang != "default" {
		langKey := fmt.Sprintf("%s:%s", templateID, lang)
		if _, ok := tm.cache[langKey]; ok {
			key = langKey
		}
	}

	tpl, ok := tm.cache[key]
	if !ok {
		return nil, fmt.Errorf("template cache not found: %s", key)
	}

	// 验证必填变量
	for _, v := range t.Variables {
		if v.Required {
			if _, ok := variables[v.Name]; !ok {
				return nil, fmt.Errorf("required variable missing: %s", v.Name)
			}
		}
	}

	// 渲染模板
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, variables); err != nil {
		return nil, fmt.Errorf("failed to render template: %w", err)
	}

	// 解析渲染结果为结构化数据
	var data map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &data); err != nil {
		// 如果不是 JSON，则作为纯文本内容
		data = map[string]interface{}{"text": buf.String()}
	}

	return &RenderedMessage{
		ID:         generateID(),
		TemplateID: templateID,
		Type:       t.Type,
		Priority:   t.Priority,
		Content:    buf.String(),
		Data:       data,
	}, nil
}

// Unregister 注销模板
func (tm *TemplateManager) Unregister(id string) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	delete(tm.templates, id)
	delete(tm.cache, id)
}

// List 列出所有模板
func (tm *TemplateManager) List() []*MessageTemplate {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	list := make([]*MessageTemplate, 0, len(tm.templates))
	for _, t := range tm.templates {
		list = append(list, t)
	}
	return list
}

func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// 预设模板示例
func GetDefaultTemplates() []*MessageTemplate {
	return []*MessageTemplate{
		{
			ID:          "order_created",
			Name:        "订单创建通知",
			Description: "用户下单成功后发送",
			Type:        MessageTypeBusiness,
			Priority:    PriorityNormal,
			Content: TemplateContent{
				Default: `{"title": "订单创建成功", "order_id": "{{.OrderID}}", "amount": "{{.Amount}}", "message": "您的订单 {{.OrderID}} 已创建，金额 ¥{{.Amount}}"}`,
				I18n: map[string]string{
					"en-US": `{"title": "Order Created", "order_id": "{{.OrderID}}", "amount": "{{.Amount}}", "message": "Your order {{.OrderID}} has been created, amount ${{.Amount}}"}`,
				},
			},
			Variables: []TemplateVar{
				{Name: "OrderID", Type: "string", Required: true},
				{Name: "Amount", Type: "number", Required: true},
			},
			Enabled: true,
		},
		{
			ID:          "payment_success",
			Name:        "支付成功通知",
			Description: "用户支付成功后发送",
			Type:        MessageTypeBusiness,
			Priority:    PriorityHigh,
			Content: TemplateContent{
				Default: `{"title": "支付成功", "order_id": "{{.OrderID}}", "amount": "{{.Amount}}", "message": "订单 {{.OrderID}} 支付成功，金额 ¥{{.Amount}}"}`,
			},
			Variables: []TemplateVar{
				{Name: "OrderID", Type: "string", Required: true},
				{Name: "Amount", Type: "number", Required: true},
			},
			Enabled: true,
		},
		{
			ID:          "system_alert",
			Name:        "系统警告",
			Description: "系统异常时发送给管理员",
			Type:        MessageTypeSystem,
			Priority:    PriorityUrgent,
			Content: TemplateContent{
				Default: `{"title": "系统警告", "level": "{{.Level}}", "message": "{{.Message}}", "time": "{{.Time}}"}`,
			},
			Variables: []TemplateVar{
				{Name: "Level", Type: "string", Required: true},
				{Name: "Message", Type: "string", Required: true},
				{Name: "Time", Type: "string", Required: false, Default: ""},
			},
			Enabled: true,
		},
	}
}

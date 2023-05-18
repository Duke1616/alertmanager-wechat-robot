package im

import "context"

// IM消息通知器
type ImNotifier interface {
	SendMessage(context.Context, *SendMessageRequest) error
}

func NewSendMessageRequest(targetId, title, content string) *SendMessageRequest {
	return &SendMessageRequest{
		TargetId: targetId,
		Title:    title,
		Content:  content,
	}
}

type SendMessageRequest struct {
	// 消息标题
	TargetId string `json:"target_id" validate:"required"`
	// 消息标题
	Title string `json:"title" validate:"required"`
	// 消息内容
	Content string `json:"content"`
	// 消息格式, 可以为空
	ContentType string `json:"content_type"`
}

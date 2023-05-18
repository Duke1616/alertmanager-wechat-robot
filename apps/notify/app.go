package notify

const (
	AppName = "notify"
)

func NewSendNotifyRequest() *SendNotifyRequest {
	return &SendNotifyRequest{
		NotifyType: NOTIFY_TYPE_WECHAT,
		Users:      []string{},
	}
}

func NewSendWechatRequest(title, content, targetId string) *SendNotifyRequest {
	req := NewSendNotifyRequest()
	req.NotifyType = NOTIFY_TYPE_WECHAT
	req.TargetId = targetId
	req.Title = title
	req.Content = content
	return req
}

func (req *SendNotifyRequest) AddUser(users ...string) {
	req.Users = append(req.Users, users...)
}

func NewWechatRequest(url, secret string) *WechatRequest {
	return &WechatRequest{
		Url:    url,
		Secret: secret,
	}
}

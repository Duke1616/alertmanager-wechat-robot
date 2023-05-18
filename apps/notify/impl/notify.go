package impl

import (
	"context"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/notify"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/notify/provider/im"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/notify/provider/im/wechat"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/target"
)

func (s *service) SendNotify(ctx context.Context, req *notify.SendNotifyRequest) (*notify.SendResponse, error) {
	resp := &notify.SendResponse{
		Success: true,
	}

	switch req.NotifyType {
	case notify.NOTIFY_TYPE_WECHAT:
		sendReq := im.NewSendMessageRequest(req.TargetId, req.Title, req.Content)
		sendReq.ContentType = req.ContentType
		err := s.SendWechat(ctx, sendReq)

		if err != nil {
			return nil, err
		}
	case notify.NOTIFY_TYPE_EMAIL:

	}

	return resp, nil
}

func (s *service) SendWechat(ctx context.Context, req *im.SendMessageRequest) error {
	t, err := s.target.DescribeTarget(ctx, target.NewDescribeTargetRequestById(req.TargetId))

	if err != nil {
		return err
	}

	wechatReq := notify.NewWechatRequest(t.Spec.Url, t.Spec.Secret)

	notifier := wechat.NewWechatNotify(wechatReq)

	err = notifier.SendMessage(ctx, req)

	if err != nil {
		return err
	}

	return nil
}

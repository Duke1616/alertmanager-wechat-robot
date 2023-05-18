package wechat

import (
	"context"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/notify"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/notify/provider/im"
	robot "github.com/group-robot/work-weixin-robot"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewWechatNotify(conf *notify.WechatRequest) im.ImNotifier {
	return &Wechat{
		conf: conf,
		log:  zap.L().Named("notifier.wechat"),
	}
}

type Wechat struct {
	conf *notify.WechatRequest
	log  logger.Logger
}

func (w *Wechat) SendMessage(ctx context.Context, req *im.SendMessageRequest) error {
	content := req.Title + "\n" + req.Content
	msg := robot.NewMarkdownMessage(content)

	webhookUrl := w.conf.Url + "/cgi-bin/webhook/send?key=" + w.conf.Secret

	client := robot.NewRobotClientByWebHook(webhookUrl)

	_, err := client.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}

package impl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/notifier"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/rule"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/target"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
)

func (s *service) SendWechatRobot(ctx context.Context, req *notifier.NotificationInfo) (*emptypb.Empty, error) {
	t, err := s.target.DescribeTarget(ctx, target.NewDescribeTargetRequestByName(req.Name))
	if err != nil {
		return nil, err
	}

	notifierWechat := notifier.NewNotifierWechat(req.Notification)

	// 查询所有rule
	rs, err := s.rule.QueryRule(ctx, rule.NewQueryRuleRequest(t.Id))
	if err != nil {
		return nil, err
	}

	// 规则判断，数据预处理
	err = notifierWechat.HasRule(rs)
	if err != nil {
		return nil, err
	}

	// 报警信息拼接
	md, err := s.TransFormToWechat(ctx, notifierWechat)
	if err != nil {
		return nil, err
	}

	// 发送消息通知
	wechatUrl := fmt.Sprintf("%s%s%s", t.Url, "/cgi-bin/webhook/send?key=", t.Secret)
	err = s.send(wechatUrl, md)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *service) TransFormToWechat(ctx context.Context, req *notifier.NotificationWechat) (*notifier.Message, error) {
	status := req.Notification.Status

	var buffer bytes.Buffer
	// 查看状态
	if status == "firing" {
		buffer.WriteString(fmt.Sprintf("<font color=\"warning\">**告警状态：%s  报警数量: %d**</font> \n", status, len(req.Notification.Alerts)))
	} else {
		buffer.WriteString(fmt.Sprintf("<font color=\"info\">**告警状态：%s  报警数量: %d**</font> \n", status, len(req.Notification.Alerts)))
	}

	// 消息通知
	if req.Mention.All == true {
		// TODO 目前企业微信机器人不支持@所有人，需要单独扩展
		buffer.WriteString(fmt.Sprintf("\n通知相关人员: <font color=\"info\"><@all></font>"))
	} else {
		for _, men := range req.Mention.Mobiles {
			buffer.WriteString(fmt.Sprintf("\n通知相关人员: <font color=\"info\"><@%s></font>", men))
		}
	}

	// 报警相关信息
	for _, alert := range req.Notification.Alerts {
		labels := alert.Labels
		buffer.WriteString(fmt.Sprintf("\n>告警级别: %s", labels["severity"]))
		buffer.WriteString(fmt.Sprintf("\n>告警类型: %s", labels["alertname"]))
		buffer.WriteString(fmt.Sprintf("\n>告警主机: %s", labels["instance"]))
		annotations := alert.Annotations
		buffer.WriteString(fmt.Sprintf("\n>告警主题: <font color=\"info\">**%s**</font>", annotations["summary"]))
		buffer.WriteString(fmt.Sprintf("\n>告警详情: %s", annotations["description"]))
		buffer.WriteString(fmt.Sprintf("\n>告警时间: %s\n", alert.StartsAt.AsTime().Format("2006-01-02 15:04:05")))
		if len(req.Notification.Alerts) != 1 {
			buffer.WriteString(fmt.Sprint("\n=============华丽的分割线线=============="))
		}

	}

	markdown := &notifier.Message{
		Msgtype: "markdown",
		Markdown: &notifier.Markdown{
			Content: buffer.String(),
		},
	}

	return markdown, nil
}

func (s *service) send(wechatUrl string, md *notifier.Message) error {
	data, err := json.Marshal(md)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", wechatUrl, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	request.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	s.log.Info(resp.Status)
	s.log.Info(resp.Header)

	return nil
}

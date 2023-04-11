package impl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/notifier"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/target"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
)

func (s *service) SendWechatRobot(ctx context.Context, req *notifier.NotificationInfo) (*emptypb.Empty, error) {
	t, err := s.target.DescribeTarget(ctx, target.NewDescribeTargetRequestByName(req.Name))
	if err != nil {
		return nil, err
	}

	for _, v := range t.Rule {
		value, ok := req.Notification.CommonLabels[v.Label]
		if ok {
			fmt.Println(value)
		}
	}

	//md, err := s.TransFormToWechat(ctx, req)
	//if err != nil {
	//	return nil, nil
	//}
	//
	//wechatUrl := fmt.Sprintf("%s%s", t.Url, t.Secret)

	//err = send(wechatUrl, md)
	//if err != nil {
	//	return nil, err
	//}
	//if req.Notification.CommonAnnotations[t.Rule]
	//fmt.Println(req.Notification.CommonLabels)
	return nil, nil
}

func (s *service) TransFormToWechat(ctx context.Context, req *notifier.NotificationInfo) (*notifier.Message, error) {
	status := req.Notification.Status

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("### 当前状态:%s, 数量: %d \n", status, len(req.Notification.Alerts)))

	for _, alert := range req.Notification.Alerts {
		labels := alert.Labels
		buffer.WriteString(fmt.Sprintf("\n>告警级别: %s", labels["severity"]))
		buffer.WriteString(fmt.Sprintf("\n>告警类型: %s", labels["alertname"]))
		buffer.WriteString(fmt.Sprintf("\n>告警主机: %s", labels["instance"]))
		annotations := alert.Annotations
		buffer.WriteString(fmt.Sprintf("\n>告警主题: %s", annotations["summary"]))
		buffer.WriteString(fmt.Sprintf("\n>告警详情: %s", annotations["description"]))
		buffer.WriteString(fmt.Sprintf("\n>告警时间: %s", alert.StartsAt.AsTime().Format("2006-01-02 15:04:05")))
		buffer.WriteString(fmt.Sprint("\n=============华丽的分割线条=============="))
	}

	markdown := &notifier.Message{
		Msgtype: "markdown",
		Markdown: &notifier.Markdown{
			Content: buffer.String(),
		},
	}

	return markdown, nil
}

func send(wechatUrl string, md *notifier.Message) error {
	fmt.Println(md)

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
	fmt.Println("response Status", resp.Status)
	fmt.Println("response Status", resp.Header)

	return nil
}

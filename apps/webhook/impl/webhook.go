package impl

import (
	"bytes"
	"context"
	"fmt"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/filter"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/history"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/notify"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/policy"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/webhook"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *service) TransformAlert(ctx context.Context, req *webhook.TransData) (*emptypb.Empty, error) {
	// 查看规则、根据优先级进行排序
	query := &policy.QueryPolicyRequest{}

	// 查询filter
	if req.IsFilter == "true" {
		queryFilter := filter.NewDefaultQueryFilterRequest()
		queryFilter.Meta = req.GroupLabels
		set, err := s.filter.QueryFilter(ctx, queryFilter)
		if err != nil {
			return nil, err
		}

		if set != nil || set.Total > 0 {
			// 根据条件进行判断
			query = policy.NewQueryPolicyRequestByFilter("priority", set.Items[0].Spec.Name)
		}
	} else {
		query = policy.NewQueryPolicyRequestBySort("priority")
	}

	// 查询策略
	policySet, err := s.policy.QueryPolicy(ctx, query)
	if err != nil {
		return nil, err
	}

	// 策略匹配、验证
	transSet := req.HasPolicy(policySet)

	// TODO 数据合并处理, 合并重复报警项 k,v

	// 发送告警
	for _, data := range transSet.Items {
		err = s.TransNotify(ctx, data)
		if err != nil {
			// 报警发送错误记录
			return nil, err
		}

		_, err = s.history.SaveHistory(ctx, history.NewSaveHistoryRequest(data))

		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (s *service) TransNotify(ctx context.Context, trans *webhook.TransformAlertData) error {
	var buffer bytes.Buffer
	title := ""

	for _, alert := range trans.Alerts {
		labels := alert.Labels
		if labels["level"] == "P4" || labels["level"] == "P5" {
			title = fmt.Sprintf("\n><font color=\"info\">**告警级别：%s      告警数量：%d**</font>",
				labels["level"], len(trans.Alerts))
		} else if labels["level"] == "P3" || labels["level"] == "P2" {
			title = fmt.Sprintf("\n><font color=\"warning\">**告警级别：%s      告警数量：%d**</font>",
				labels["level"], len(trans.Alerts))
		} else {
			title = fmt.Sprintf("\n><font color=\"red\">**告警级别：%s      告警数量：%d**</font>",
				labels["level"], len(trans.Alerts))
		}

		buffer.WriteString(fmt.Sprintf("\n>告警名称: %s", labels["alertname"]))
		buffer.WriteString(fmt.Sprintf("\n>告警主机: %s", labels["instance"]))
		annotations := alert.Annotations
		buffer.WriteString(fmt.Sprintf("\n>告警主题: <font color=\"info\">**%s**</font>", annotations["summary"]))
		buffer.WriteString(fmt.Sprintf("\n>告警详情: %s", annotations["description"]))
		buffer.WriteString(fmt.Sprintf("\n>告警时间: %s\n", alert.StartsAt.AsTime().Local().
			Format("2006-01-02 15:04:05")))

		if len(trans.Alerts) != 1 {
			buffer.WriteString(fmt.Sprint("\n=============华丽的分割线线==============\n"))
		}

	}

	_, err := s.notify.SendNotify(ctx, notify.NewSendWechatRequest(title, buffer.String(), trans.TargetId))

	if err != nil {
		return err
	}

	return nil
}

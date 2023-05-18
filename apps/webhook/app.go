package webhook

import (
	"fmt"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/policy"
)

const (
	AppName = "webhook"
)

func newTransformAlertDataSet() *TransformAlertDataSet {
	return &TransformAlertDataSet{
		Items: []*TransformAlertData{},
	}
}

func newTransformAlertDataDefault() *TransformAlertData {
	return &TransformAlertData{
		Users:  []string{},
		Alerts: []*Alert{},
	}
}

func (s *TransformAlertDataSet) Add(item *TransformAlertData) {
	s.Total++
	s.Items = append(s.Items, item)
}

func (x *TransData) HasPolicy(policySet *policy.PolicySet) *TransformAlertDataSet {
	alertDataSet := newTransformAlertDataSet()

	for _, a := range x.Alerts {
		var flag bool
		matchNum := 0

		// 遍历所有策略
		for _, p := range policySet.Items {
			// 遍历策略定义的tag label
			for _, tag := range p.Spec.Tags {
				value, ok := a.Labels[tag.Label]
				if ok {
					if tag.Value == value {
						matchNum++
					}
				}
			}

			// 执行对应的策略
			if matchNum == len(p.Spec.Tags) {
				hasFlag := x.HasActive(p, alertDataSet, a)
				flag = hasFlag
			}

			// 判断是否跳出循环
			if flag {
				break
			}
		}
	}

	return alertDataSet
}

func (x *TransData) HasActive(p *policy.Policy, transSet *TransformAlertDataSet, alert *Alert) bool {
	switch p.Spec.Active {
	case policy.ACTIVE_NEWS:
		for _, t := range transSet.Items {
			// 判断是否已经存在、存在则插入，否则创建
			if p.Spec.TargetId == t.TargetId {
				x.AddMentionByAlert(p, alert)
				t.Alerts = append(t.Alerts, alert)
				return true
			}
		}

		trans := newTransformAlertDataDefault()
		trans.TargetId = p.Spec.TargetId
		x.AddMentionByTrans(p, trans)
		trans.Alerts = append(trans.Alerts, alert)
		transSet.Add(trans)

		return true
	case policy.ACTIVE_STATUS_TRANS:
		alert.Status = "resolved"
		return false
	case policy.ACTIVE_DROP:
		return true
	case policy.ACTIVE_SKIP:
		return true
	default:
		fmt.Println("暂无条件匹配")
		return false
	}
}

func (x *TransData) AddMentionByAlert(p *policy.Policy, alert *Alert) {
	if p.Spec.Mention.All == true {
		//	TODO 查询群组所有人插入数组
	} else {
		for _, user := range p.Spec.Mention.Users {
			alert.Users = append(alert.Users, user)
		}
	}
}

func (x *TransData) AddMentionByTrans(p *policy.Policy, trans *TransformAlertData) {
	if p.Spec.Mention.All == true {
		//	TODO 查询群组所有人插入数组
	} else {
		for _, user := range p.Spec.Mention.Users {
			trans.Users = append(trans.Users, user)
		}
	}
}

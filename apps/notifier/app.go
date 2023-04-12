package notifier

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/rule"
	"net/http"
)

const (
	AppName = "notifier"
)

func NewQueryTargetRequestFromHTTP(r *http.Request, name string) string {
	qs := r.URL.Query()
	return qs.Get(name)
}

func NewNotifierWechat(req *Notification) *NotificationWechat {
	return &NotificationWechat{
		Mention:      &rule.Mention{},
		Notification: req,
	}
}

func (n *NotificationWechat) HasRule(ruleSet *rule.RuleSet) error {
	// TODO 根据target ID 获取所有的rule
	for _, r := range ruleSet.Items {
		if r.Enabled != true {
			continue
		}

		switch r.Spec.LabelType {
		case rule.LABEL_TYPE_GROUP:
			value, ok := n.Notification.GroupLabels[r.Spec.Label]
			if ok {
				n.HasActive(r, value)
			}
		case rule.LABEL_TYPE_COMMON:
			value, ok := n.Notification.CommonLabels[r.Spec.Label]
			if ok {
				n.HasActive(r, value)
			}
		}
	}
	return nil
}

func (n *NotificationWechat) HasActive(r *rule.Rule, value string) error {
	for _, v := range r.Spec.Value {
		if v == value {
			switch r.Spec.Active {
			case rule.ACTIVE_NEWS:
				n.Mention = r.Spec.Mention
			case rule.ACTIVE_DROP:
				return nil
			case rule.ACTIVE_STATUS_TRANS:
				n.Notification.Status = value
			default:
				return nil
			}
		}
	}
	return nil
}

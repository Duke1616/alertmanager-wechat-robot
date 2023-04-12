package notifier

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/target"
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
		Mention:      &target.Mention{},
		Notification: req,
	}
}

func (n *NotificationWechat) HasRule(t *target.Target) error {
	for _, rule := range t.Rule {
		if rule.Enabled != true {
			continue
		}

		switch rule.LabelType {
		case target.LABEL_TYPE_GROUP:
			value, ok := n.Notification.GroupLabels[rule.Label]
			if ok {
				n.HasActive(rule, value)
			}
		case target.LABEL_TYPE_COMMON:
			value, ok := n.Notification.CommonLabels[rule.Label]
			if ok {
				n.HasActive(rule, value)
			}
		}
	}
	return nil
}

func (n *NotificationWechat) HasActive(rule *target.Rule, value string) error {
	for _, v := range rule.Value {
		if v == value {
			switch rule.Active {
			case target.ACTIVE_NEWS:
				n.Mention = rule.Mention
			case target.ACTIVE_DROP:
				return nil
			case target.ACTIVE_STATUS_TRANS:
				n.Notification.Status = value
			default:
				return nil
			}
		}
	}
	return nil
}

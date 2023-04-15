package notifier

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/policy"
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
		Mention:      &policy.Mention{},
		Notification: req,
	}
}

func (n *NotificationWechat) HasRule(ruleSet *policy.PolicySet) error {
	// TODO 根据target ID 获取所有的rule
	for _, r := range ruleSet.Items {
		if r.Enabled != true {
			continue
		}

		switch r.Spec.LabelType {
		case policy.LABEL_TYPE_GROUP:
			value, ok := n.Notification.GroupLabels[r.Spec.Label]
			if ok {
				n.HasActive(r, value)
			}
		case policy.LABEL_TYPE_COMMON:
			value, ok := n.Notification.CommonLabels[r.Spec.Label]
			if ok {
				n.HasActive(r, value)
			}
		}
	}
	return nil
}

func (n *NotificationWechat) HasActive(r *policy.Policy, value string) error {
	for _, v := range r.Spec.Value {
		if v == value {
			switch r.Spec.Active {
			case policy.ACTIVE_NEWS:
				n.Mention = r.Spec.Mention
			case policy.ACTIVE_DROP:
				return nil
			case policy.ACTIVE_STATUS_TRANS:
				n.Notification.Status = value
			default:
				return nil
			}
		}
	}
	return nil
}

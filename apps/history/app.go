package history

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/webhook"
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
	"github.com/rs/xid"
	"net/http"
	"time"
)

const (
	AppName = "history"
)

var (
	validate = validator.New()
)

func NewSaveHistoryRequest(data *webhook.TransformAlertData) *SaveHistoryRequest {
	return &SaveHistoryRequest{
		Data: data,
	}
}

func (req *SaveHistoryRequest) HistorySet(targetName string) []interface{} {
	his := make([]interface{}, 0, len(req.Data.Alerts))
	for _, v := range req.Data.Alerts {
		information := &History{}
		labels := v.Labels
		annotations := v.Annotations

		if v.Status == "firing" {
			information.Type = "告警通知"
		} else {
			information.Type = "告警恢复"
		}

		information.Id = xid.New().String()
		information.Alertname = labels["alertname"]
		information.Level = labels["level"]
		information.EndAt = v.EndsAt.AsTime().Local().UnixMilli()
		information.StartAt = v.StartsAt.AsTime().Local().UnixMilli()
		information.NotificationAt = time.Now().Local().UnixMilli()
		information.Description = annotations["description"]
		information.InstanceName = labels["customer_city"] + labels["customer_name"]
		information.Target = targetName

		information.Users = req.Data.Users
		information.Service = labels["service_name"]
		information.Source = labels["source"]
		his = append(his, information)
	}

	return his
}

func NewQueryHistoryRequestFromHTTP(r *http.Request) *QueryHistoryRequest {
	page := request.NewPageRequestFromHTTP(r)

	qs := r.URL.Query()
	req := NewDefaultQueryAlertRequest()
	req.Type = qs.Get("type")
	req.Level = qs.Get("level")
	req.Service = qs.Get("service")
	req.Target = qs.Get("target")
	req.User = qs.Get("user")
	req.Instance = qs.Get("instance")
	req.Alertname = qs.Get("alertname")
	req.Page = page
	return req
}

func (req *QueryHistoryRequest) Validate() error {
	return validate.Struct(req)
}

func NewDefaultQueryAlertRequest() *QueryHistoryRequest {
	return &QueryHistoryRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func NewHistorySet() *HistorySet {
	return &HistorySet{
		Items: []*History{},
	}
}

func (s *HistorySet) Add(item *History) {
	s.Total++
	s.Items = append(s.Items, item)
}

func NewDefaultHistory() *History {
	return &History{}
}

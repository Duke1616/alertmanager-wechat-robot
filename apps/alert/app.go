package alert

import (
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
	"net/http"
	"strings"
	"time"
)

const (
	AppName = "alert"
)

var (
	validate = validator.New()
)

func (req *SaveAlertRequest) AlertInformationS() []interface{} {
	ais := make([]interface{}, 0, len(req.Alerts))
	for _, v := range req.Alerts {
		information := &AlertInformation{}
		labels := v.Labels
		annotations := v.Annotations

		if v.Status == "firing" {
			information.Type = "告警通知"
		} else {
			information.Type = "告警恢复"
		}
		information.AlertName = labels["alertname"]
		information.Level = labels["level"]
		information.EndAt = v.EndsAt.AsTime().UnixMilli()
		information.StartAt = v.StartsAt.AsTime().UnixMilli()
		information.NotificationAt = time.Now().UnixMilli()
		information.Description = annotations["description"]
		information.InstanceName = labels["customer_city"] + labels["customer_name"]
		information.TargetName = req.TargetName
		information.Others = strings.Join(req.Mention.Mobiles, ",")
		information.ServiceName = labels["service_name"]
		information.Source = labels["source"]
		ais = append(ais, information)
	}
	return ais
}

func NewQueryAlertRequestFromHTTP(r *http.Request) *QueryAlertRequest {
	page := request.NewPageRequestFromHTTP(r)

	qs := r.URL.Query()
	req := NewDefaultQueryAlertRequest()
	req.Type = qs.Get("type")
	req.Level = qs.Get("level")
	req.ServiceName = qs.Get("service_name")
	req.TargetName = qs.Get("target_name")
	req.Others = qs.Get("others")
	req.InstanceName = qs.Get("instance_name")
	req.AlertName = qs.Get("alert_name")
	req.Page = page
	return req
}

func NewDefaultQueryAlertRequest() *QueryAlertRequest {
	return &QueryAlertRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func (req *QueryAlertRequest) Validate() error {
	return validate.Struct(req)
}

func NewAlertInformationSet() *AlertInformationSet {
	return &AlertInformationSet{
		Items: []*AlertInformation{},
	}
}

func (s *AlertInformationSet) Add(item *AlertInformation) {
	s.Total++
	s.Items = append(s.Items, item)
}

func NewDefaultAlertInformation() *AlertInformation {
	return &AlertInformation{}
}

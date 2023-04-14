package alert

import (
	"strings"
	"time"
)

const (
	AppName = "alert"
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

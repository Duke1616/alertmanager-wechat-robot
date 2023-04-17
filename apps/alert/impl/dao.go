package impl

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/alert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newQueryAlertRequest(req *alert.QueryAlertRequest) (*queryAlertRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &queryAlertRequest{
		QueryAlertRequest: req}, nil
}

type queryAlertRequest struct {
	*alert.QueryAlertRequest
}

func (r *queryAlertRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryAlertRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Type != "" {
		filter["type"] = r.Type
	}
	if r.Level != "" {
		filter["level"] = r.Level
	}
	if r.ServiceName != "" {
		filter["service_name"] = r.ServiceName
	}
	if r.TargetName != "" {
		filter["target_name"] = r.TargetName
	}
	if r.Others != "" {
		filter["others"] = r.Others
	}
	if r.InstanceName != "" {
		filter["instance_name"] = r.InstanceName
	}

	if r.AlertName != "" {
		filter["alert_name"] = r.AlertName
	}
	return filter
}

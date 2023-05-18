package impl

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/history"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newQueryHistoryRequest(req *history.QueryHistoryRequest) (*queryHistoryRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &queryHistoryRequest{
		QueryHistoryRequest: req}, nil
}

type queryHistoryRequest struct {
	*history.QueryHistoryRequest
}

func (r *queryHistoryRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryHistoryRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Type != "" {
		filter["type"] = r.Type
	}
	if r.Level != "" {
		filter["level"] = r.Level
	}
	if r.Service != "" {
		filter["service"] = r.Service
	}
	if r.Target != "" {
		filter["target"] = r.Target
	}
	if r.User != "" {
		filter["user"] = r.User
	}
	if r.Instance != "" {
		filter["instance"] = r.Instance
	}

	if r.Alertname != "" {
		filter["alertname"] = r.Alertname
	}
	return filter
}

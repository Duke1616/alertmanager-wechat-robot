package impl

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/target"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newQueryTargetRequest(req *target.QueryTargetRequest) (*queryTargetRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &queryTargetRequest{
		QueryTargetRequest: req}, nil
}

type queryTargetRequest struct {
	*target.QueryTargetRequest
}

func (r *queryTargetRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryTargetRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Url != "" {
		filter["url"] = r.Url
	}

	if r.Secret != "" {
		filter["secret"] = r.Secret
	}

	return filter
}

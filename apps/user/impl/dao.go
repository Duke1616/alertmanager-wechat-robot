package impl

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newQueryUserRequest(req *user.QueryUserRequest) (*queryUserRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &queryUserRequest{
		QueryUserRequest: req}, nil
}

type queryUserRequest struct {
	*user.QueryUserRequest
}

func (r *queryUserRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryUserRequest) FindFilter() bson.M {
	filter := bson.M{}

	if len(r.TargetIds) > 1 {
		filter["spec.target_ids"] = bson.M{"$in": r.TargetIds}
	}

	return filter
}

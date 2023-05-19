package impl

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/filter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newQueryFilterRequest(req *filter.QueryFilterRequest) (*queryFilterRequest, error) {
	//if err := req.Validate(); err != nil {
	//	return nil, err
	//}
	//
	return &queryFilterRequest{
		QueryFilterRequest: req}, nil
}

type queryFilterRequest struct {
	*filter.QueryFilterRequest
}

func (r *queryFilterRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryFilterRequest) FindFilter() bson.M {
	filters := bson.M{}

	return filters
}

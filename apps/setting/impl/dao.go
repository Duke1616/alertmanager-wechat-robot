package impl

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/setting"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newQuerySettingRequest(req *setting.QuerySettingRequest) (*querySettingRequest, error) {
	//if err := req.Validate(); err != nil {
	//	return nil, err
	//}
	return &querySettingRequest{
		QuerySettingRequest: req}, nil
}

type querySettingRequest struct {
	*setting.QuerySettingRequest
}

func (r *querySettingRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *querySettingRequest) FindFilter() bson.M {
	filter := bson.M{}

	return filter
}

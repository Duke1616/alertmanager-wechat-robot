package impl

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/rule"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newQueryRuleRequest(req *rule.QueryRuleRequest) (*queryRuleRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &queryRuleRequest{
		QueryRuleRequest: req}, nil
}

type queryRuleRequest struct {
	*rule.QueryRuleRequest
}

func (r *queryRuleRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryRuleRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.TargetId != "" {
		filter["spec.target_id"] = r.TargetId
	}

	// TODO 根据标签过滤
	//switch r.LabelType {
	//case rule.LABEL_TYPE_GROUP:
	//	filter["spec.label_type"] = r.LabelType
	//case rule.LABEL_TYPE_COMMON:
	//	filter["spec.label_type"] = r.LabelType
	//}

	return filter
}

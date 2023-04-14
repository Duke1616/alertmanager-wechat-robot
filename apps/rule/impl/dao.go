package impl

import (
	"context"
	"fmt"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/rule"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/target"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

	if len(r.RuleIds) > 1 {
		filter["_id"] = bson.M{"$in": r.RuleIds}
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

func (s *service) update(ctx context.Context, ins *target.Target) error {
	if _, err := s.col.UpdateByID(ctx, ins.Id, bson.M{"$set": ins}); err != nil {
		return exception.NewInternalServerError("update target(%s) document error, %s",
			ins.Id, err)
	}

	return nil
}

func (s *service) delete(ctx context.Context, set *rule.RuleSet) error {
	if set == nil || len(set.Items) == 0 {
		return fmt.Errorf("rule is nil")
	}

	var result *mongo.DeleteResult
	var err error
	if len(set.Items) == 1 {
		result, err = s.col.DeleteMany(ctx, bson.M{"_id": set.RuleIds()[0]})
	} else {
		result, err = s.col.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": set.RuleIds()}})
	}

	if err != nil {
		return exception.NewInternalServerError("delete rule(%s) error, %s", set, err)
	}

	if result.DeletedCount == 0 {
		return exception.NewNotFound("rule %s not found", set)
	}

	return nil
}

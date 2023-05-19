package impl

import (
	"context"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/rule"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/setting"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *service) saveRules(ctx context.Context, rules *rule.Rules, groupIds []string, conf *setting.Setting) error {
	// 删除rule
	_, err := s.DeleteRule(ctx, rule.NewDeleteRuleRequest(groupIds))
	if err != nil {
		return err
	}

	// 新增rule
	for _, v := range rules.Data.Groups {
		ruleset := rules.RuleSet(v, conf)

		if _, err = s.rule.InsertMany(ctx, ruleset); err != nil {
			return exception.NewInternalServerError("inserted a rule document error, %s", err)
		}
	}

	return nil
}

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

	if len(r.GroupIds) > 0 {
		filter["group_id"] = bson.M{"$in": r.GroupIds}
	}

	if r.Service != "" {
		filter["service"] = r.Service
	}

	if r.Level != "" {
		filter["level"] = r.Level
	}

	return filter
}

func (s *service) deleteRule(ctx context.Context, set *rule.RuleSet) error {
	if set == nil || len(set.Items) == 0 {
		return nil
	}

	var result *mongo.DeleteResult
	var err error
	if len(set.Items) == 1 {
		result, err = s.rule.DeleteMany(ctx, bson.M{"_id": set.RuleIds()[0]})
	} else {
		result, err = s.rule.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": set.RuleIds()}})
	}

	if err != nil {
		return exception.NewInternalServerError("delete rule(%s) error, %s", set, err)
	}

	if result.DeletedCount == 0 {
		return exception.NewNotFound("rule %s not found", set)
	}

	return nil
}

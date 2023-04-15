package impl

import (
	"context"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/rule"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *service) saveRules(ctx context.Context, rules *rule.Rules) error {
	for _, v := range rules.Data.Groups {
		ruleset := rules.RuleSet(v)

		// 更新已有rule的记录
		newRule := make([]interface{}, 0, len(ruleset))
		for i := range ruleset {
			if err := s.rule.FindOneAndReplace(ctx, bson.M{"_id": ruleset[i].Id}, ruleset[i]).Err(); err != nil {
				if err == mongo.ErrNoDocuments {
					newRule = append(newRule, ruleset[i])
				} else {
					return err
				}
			}
		}

		// 插入rule新增记录
		if len(newRule) > 0 {
			if _, err := s.rule.InsertMany(ctx, newRule); err != nil {
				return exception.NewInternalServerError("inserted a rule document error, %s", err)
			}
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

	if r.Name != "" {
		filter["name"] = r.Name
	}

	if r.GroupName != "" {
		filter["group_name"] = r.GroupName
	}

	if r.ServiceName != "" {
		filter["service_name"] = r.ServiceName
	}

	if r.Level != "" {
		filter["level"] = r.Level
	}

	return filter
}

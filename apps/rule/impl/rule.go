package impl

import (
	"context"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/rule"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/target"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) CreateRule(ctx context.Context, req *rule.CreateRuleRequest) (*rule.Rule, error) {
	_, err := s.target.DescribeTarget(ctx, target.NewDescribeTargetRequestById(req.TargetId))
	if err != nil {
		return nil, err
	}

	r, err := rule.NewRule(req)

	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}
	if _, err := s.col.InsertOne(ctx, r); err != nil {
		return nil, exception.NewInternalServerError("inserted a rule document error, %s", err)
	}

	return r, nil
}

func (s *service) DescribeRule(ctx context.Context, req *rule.DescribeRuleRequest) (*rule.Rule, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	filter := bson.M{}
	switch req.DescribeBy {
	case rule.DESCRIBE_BY_id:
		filter["_id"] = req.Id
	}

	r := rule.NewDefaultRole()
	if err := s.col.FindOne(ctx, filter).Decode(r); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("rule %s not found", req)
		}

		return nil, exception.NewInternalServerError("find rule %s error, %s", req.Id, err)
	}

	return r, nil
}

func (s *service) QueryRule(ctx context.Context, req *rule.QueryRuleRequest) (*rule.RuleSet, error) {
	query, err := newQueryRuleRequest(req)
	if err != nil {
		return nil, err
	}

	s.log.Debugf("query rule filter: %s", query.FindFilter())
	resp, err := s.col.Find(ctx, query.FindFilter(), query.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find rule error, error is %s", err)
	}

	set := rule.NewRuleSet()
	// 循环插入数据
	for resp.Next(ctx) {
		ins := rule.NewDefaultRole()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode rule error, error is %s", err)
		}
		set.Add(ins)
	}

	// 计算数量
	count, err := s.col.CountDocuments(ctx, query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get rule count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

package impl

import (
	"context"
	"encoding/json"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/rule"
	"github.com/infraboard/mcube/exception"
	"io/ioutil"
	"net/http"
)

func (s *service) SyncRule(ctx context.Context, req *rule.SyncRuleRequest) (*rule.Empty, error) {
	rules := rule.NewRules()

	resp, err := http.Get(s.url)
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	err = json.Unmarshal(result, rules)
	if err != nil {
		return nil, err
	}

	if rules.Status != "success" {
		exception.NewInternalServerError("请求vmalert返回错误")
	}

	err = s.saveGroups(ctx, rules)
	if err != nil {
		return nil, err
	}

	err = s.saveRules(ctx, rules)
	if err != nil {
		return nil, err
	}

	return &rule.Empty{}, nil
}

func (s *service) QueryRule(ctx context.Context, req *rule.QueryRuleRequest) (*rule.RuleSet, error) {
	query, err := newQueryRuleRequest(req)
	if err != nil {
		return nil, err
	}

	s.log.Debugf("query rule filter: %s", query.FindFilter())
	resp, err := s.rule.Find(ctx, query.FindFilter(), query.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find rule error, error is %s", err)
	}

	set := rule.NewRuleSet()
	// 循环插入数据
	for resp.Next(ctx) {
		ins := rule.NewDefaultRule()
		if err = resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode rule error, error is %s", err)
		}
		set.Add(ins)
	}

	// 计算数量
	count, err := s.rule.CountDocuments(ctx, query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get rule count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

func (s *service) QueryGroup(ctx context.Context, req *rule.QueryGroupRequest) (*rule.GroupSet, error) {
	query, err := newQueryGroupRequest(req)
	if err != nil {
		return nil, err
	}

	s.log.Debugf("query group filter: %s", query.FindFilter())
	resp, err := s.group.Find(ctx, query.FindFilter(), query.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find group error, error is %s", err)
	}

	set := rule.NewGroupSet()
	// 循环插入数据
	for resp.Next(ctx) {
		ins := rule.NewDefaultGroup()
		if err = resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode rule error, error is %s", err)
		}
		set.Add(ins)
	}

	// 计算数量
	count, err := s.group.CountDocuments(ctx, query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get group count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

package impl

import (
	"context"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/rule"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/setting"
	"github.com/infraboard/mcube/exception"
)

func (s *service) SyncRule(ctx context.Context, req *rule.SyncRuleRequest) (*rule.SendResponse, error) {
	conf, err := s.setting.DescribeSetting(ctx, setting.NewDescribeSettingRequestById(req.SettingId))
	if err != nil {
		return nil, err
	}

	rules, err := rule.GetRules(conf.Spec.Url)

	if err != nil {
		s.log.Debugf("查询数据失败, %s", err)
		return nil, err
	}

	if rules.Status != "success" {
		return nil, exception.NewInternalServerError("请求vmalert返回错误")
	}

	// 保存组信息
	groupIds, err := s.saveGroups(ctx, rules, conf)
	if err != nil {
		return nil, err
	}

	// 保存告警信息
	err = s.saveRules(ctx, rules, groupIds)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *service) DeleteRule(ctx context.Context, req *rule.DeleteRuleRequest) (*rule.RuleSet, error) {
	queryReq := rule.NewDefaultQueryRuleRequest()
	queryReq.GroupIds = req.GroupIds

	set, err := s.QueryRule(ctx, queryReq)
	if err != nil {
		return nil, err
	}

	if err = s.deleteRule(ctx, set); err != nil {
		return nil, err
	}
	return set, nil
}

func (s *service) DeleteGroup(ctx context.Context, req *rule.DeleteGroupRequest) (*rule.GroupSet, error) {
	queryReq := rule.NewDefaultQueryGroupRequest()
	queryReq.Domain = req.Domain
	queryReq.Namespace = req.Namespace

	set, err := s.QueryGroup(ctx, queryReq)
	if err != nil {
		return nil, err
	}

	if err = s.deleteGroup(ctx, set); err != nil {
		return nil, err
	}
	return set, nil
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
			return nil, exception.NewInternalServerError("decode group error, error is %s", err)
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

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

	resp, err := http.Get("")
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
	return nil, nil
}

func (s *service) QueryGroup(ctx context.Context, req *rule.QueryGroupRequest) (*rule.GroupSet, error) {
	return nil, nil
}

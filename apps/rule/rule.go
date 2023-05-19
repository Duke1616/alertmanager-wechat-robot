package rule

import (
	"encoding/json"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/setting"
	"github.com/infraboard/mcube/http/request"
	"io/ioutil"
	"net/http"
)

func (req *QueryRuleRequest) Validate() error {
	return validate.Struct(req)
}

func NewRules() *Rules {
	return &Rules{}
}

func GetRules(url string) (*Rules, error) {
	rules := NewRules()
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(result, rules)

	if err != nil {
		return nil, err
	}

	return rules, nil
}

func (req *Rules) RuleSet(group *Group, conf *setting.Setting) []interface{} {
	rus := make([]interface{}, 0, len(group.Rules))
	for i := range group.Rules {
		ru := &Rule{
			Id:          group.Rules[i].Id,
			Name:        group.Rules[i].Name,
			Query:       group.Rules[i].Query,
			GroupId:     group.Rules[i].GroupId,
			Labels:      group.Rules[i].Labels,
			Domain:      conf.Spec.Domain,
			Namespace:   conf.Spec.Namespace,
			Annotations: group.Rules[i].Annotations,
			GroupName:   group.Name,
			Level:       group.Rules[i].Labels["level"],
			Service:     group.Rules[i].Labels["service"],
		}
		rus = append(rus, ru)
	}

	return rus
}

func NewQueryRuleRequestFromHTTP(r *http.Request) *QueryRuleRequest {
	page := request.NewPageRequestFromHTTP(r)

	qs := r.URL.Query()
	req := NewDefaultQueryRuleRequest()
	req.Service = qs.Get("service")
	req.Level = qs.Get("level")

	req.Page = page
	return req
}

func NewDefaultQueryRuleRequest() *QueryRuleRequest {
	return &QueryRuleRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func NewRuleSet() *RuleSet {
	return &RuleSet{
		Items: []*Rule{},
	}
}

func (s *RuleSet) Add(item *Rule) {
	s.Total++
	s.Items = append(s.Items, item)
}

func NewDefaultRule() *Rule {
	return &Rule{}
}

func NewDeleteRuleRequest(groupIds []string) *DeleteRuleRequest {
	return &DeleteRuleRequest{
		GroupIds: groupIds,
	}
}

func (s *RuleSet) RuleIds() (rus []string) {
	for i := range s.Items {
		rus = append(rus, s.Items[i].Id)
	}

	return
}

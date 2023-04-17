package rule

import (
	"github.com/infraboard/mcube/http/request"
	"net/http"
)

func (req *QueryRuleRequest) Validate() error {
	return validate.Struct(req)
}

func NewQueryRuleRequestFromHTTP(r *http.Request) *QueryRuleRequest {
	page := request.NewPageRequestFromHTTP(r)

	qs := r.URL.Query()
	req := NewDefaultQueryRuleRequest()
	req.ServiceName = qs.Get("service_name")
	req.Name = qs.Get("name")
	req.GroupName = qs.Get("group_name")
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

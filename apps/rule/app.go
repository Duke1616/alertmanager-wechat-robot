package rule

import (
	"github.com/go-playground/validator/v10"
	"github.com/imdario/mergo"
	"github.com/infraboard/mcube/http/request"
	"github.com/rs/xid"
	"net/http"
	"strings"
	"time"
)

const (
	AppName = "rule"
)

var (
	validate = validator.New()
)

func (req *CreateRuleRequest) Validate() error {
	return validate.Struct(req)
}

func NewRule(req *CreateRuleRequest) (*Rule, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &Rule{
		Id:       xid.New().String(),
		Enabled:  true,
		CreateAt: time.Now().UnixMilli(),
		Spec:     req,
	}, nil
}

func (req *DescribeRuleRequest) Validate() error {
	return validate.Struct(req)
}

func NewDefaultRole() *Rule {
	return &Rule{}
}

func NewDescribeRuleRequestById(id string) *DescribeRuleRequest {
	return &DescribeRuleRequest{
		DescribeBy: DESCRIBE_BY_id,
		Id:         id,
	}
}

func NewQueryRuleRequestFromHTTP(r *http.Request) *QueryRuleRequest {
	page := request.NewPageRequestFromHTTP(r)

	qs := r.URL.Query()
	req := NewDefaultQueryRuleRequest()
	req.TargetId = qs.Get("target_id")

	//lt := qs.Get("label_type")
	//req.LabelType, _ = ParseLABEL_TYPEFromString(lt)

	rids := qs.Get("rule_ids")
	if rids != "" {
		//index := strings.Index(rids, ",")
		//if index == -1 {
		//	rids += ","
		//}
		req.RuleIds = strings.Split(rids, ",")
	}

	req.Page = page
	return req
}

func NewDefaultQueryRuleRequest() *QueryRuleRequest {
	return &QueryRuleRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func (req *QueryRuleRequest) Validate() error {
	return validate.Struct(req)
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

func (s *RuleSet) HasTarget(ruleId string) bool {
	for i := range s.Items {
		if s.Items[i].Id == ruleId {
			return true
		}
	}

	return false
}

func (s *RuleSet) RuleIds() (rids []string) {
	for i := range s.Items {
		rids = append(rids, s.Items[i].Id)
	}

	return
}

func (i *Rule) Update(req *UpdateDeleteRequest) {
	i.UpdateAt = time.Now().UnixMicro()
}

func (i *Rule) Patch(req *UpdateDeleteRequest) error {
	i.UpdateAt = time.Now().UnixMicro()
	err := mergo.MergeWithOverwrite(i.Spec, req)
	if err != nil {
		return err
	}

	return nil
}

func NewQueryRuleRequest(targetId string) *QueryRuleRequest {
	return &QueryRuleRequest{
		Page:     request.NewPageRequest(100, 1),
		TargetId: targetId,
	}
}

func NewDeleteTargetRequest() *DeleteRuleRequest {
	return &DeleteRuleRequest{
		RuleIds: []string{},
	}
}

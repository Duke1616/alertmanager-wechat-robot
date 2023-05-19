package policy

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/imdario/mergo"
	"github.com/infraboard/mcube/http/request"
	"github.com/rs/xid"
	"net/http"
	"strings"
	"time"
)

const (
	AppName = "policy"
)

var (
	validate = validator.New()
)

func (req *CreatePolicyRequest) Validate() error {
	return validate.Struct(req)
}

func NewPolicy(req *CreatePolicyRequest) (*Policy, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	policyData := &Policy{
		Id:       xid.New().String(),
		Enabled:  true,
		CreateAt: time.Now().UnixMilli(),
		Spec:     req,
	}

	if req.Priority == 0 {
		switch req.PolicyType {
		case POLICY_TYPE_RADIO:
			policyData.Spec.Priority = 40
		case POLICY_TYPE_JOIN:
			policyData.Spec.Priority = 60
		case POLICY_TYPE_APPOINT:
			policyData.Spec.Priority = 80
		}
	}

	return policyData, nil
}

func (req *DeletePolicyRequest) Validate() error {
	return validate.Struct(req)
}

func (req *DescribePolicyRequest) Validate() error {
	return validate.Struct(req)
}

func NewDefaultPolicy() *Policy {
	return &Policy{}
}

func NewDescribePolicyRequestById(id string) *DescribePolicyRequest {
	return &DescribePolicyRequest{
		DescribeBy: DESCRIBE_BY_id,
		Id:         id,
	}
}

func NewQueryPolicyRequestFromHTTP(r *http.Request) *QueryPolicyRequest {
	page := request.NewPageRequestFromHTTP(r)

	qs := r.URL.Query()
	req := NewDefaultQueryPolicyRequest()
	req.TargetName = qs.Get("target_name")
	req.CreateType = qs.Get("create_type")
	req.PolicyType = qs.Get("policy_type")
	req.Sort = qs.Get("sort")

	rids := qs.Get("policy_ids")
	if rids != "" {
		req.PolicyIds = strings.Split(rids, ",")
	}

	req.Page = page
	return req
}

func NewDefaultQueryPolicyRequestEnum() *QueryPolicyRequest {
	return &QueryPolicyRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func NewDefaultQueryPolicyRequest() *QueryPolicyRequest {
	return &QueryPolicyRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func (req *QueryPolicyRequest) Validate() error {
	return validate.Struct(req)
}

func NewPolicySet() *PolicySet {
	return &PolicySet{
		Items: []*Policy{},
	}
}

func (s *PolicySet) Add(item *Policy) {
	s.Total++
	s.Items = append(s.Items, item)
}

func (s *PolicySet) HasPolicy(policyId string) bool {
	for i := range s.Items {
		if s.Items[i].Id == policyId {
			return true
		}
	}

	return false
}

func (s *PolicySet) PolicyIds() (rids []string) {
	for i := range s.Items {
		rids = append(rids, s.Items[i].Id)
	}

	return
}

func (i *Policy) Update(req *UpdatePolicyRequest) {
	i.UpdateAt = time.Now().UnixMicro()
}

func (i *Policy) Patch(req *UpdatePolicyRequest) error {
	i.UpdateAt = time.Now().UnixMicro()
	err := mergo.MergeWithOverwrite(i.Spec, req)
	if err != nil {
		return err
	}

	return nil
}

func NewQueryPolicyRequestByName(targetName string) *QueryPolicyRequest {
	return &QueryPolicyRequest{
		Page:       request.NewPageRequest(100, 1),
		TargetName: targetName,
	}
}

func NewQueryPolicyRequestBySort(sort string) *QueryPolicyRequest {
	return &QueryPolicyRequest{
		Page: request.NewPageRequest(100, 1),
		Sort: sort,
	}
}

func NewQueryPolicyRequestByFilter(sort, filter string) *QueryPolicyRequest {
	return &QueryPolicyRequest{
		Page:       request.NewPageRequest(100, 1),
		Sort:       sort,
		FilterName: filter,
	}
}

func NewQueryPolicyRequestByIds(ids []string) *QueryPolicyRequest {
	return &QueryPolicyRequest{
		Page:      request.NewPageRequest(100, 1),
		PolicyIds: ids,
	}
}

func NewDeletePolicyRequest() *DeletePolicyRequest {
	return &DeletePolicyRequest{
		PolicyIds: []string{},
	}
}

func (req *CreatePolicyRequest) AddTag(pos *Policy, set *PolicySet, targetName string) []interface{} {
	pSet := make([]interface{}, 0, len(set.Items))

	for _, v := range set.Items {
		cre := NewCreateRequest(req)
		p, _ := NewPolicy(cre)

		// 名称重组
		p.Spec.Name = fmt.Sprint(v.Spec.Name + pos.Spec.Name)

		// tag数组合并
		p.Spec.Tags = append(p.Spec.Tags, v.Spec.Tags[:]...)
		p.Spec.Tags = append(p.Spec.Tags, pos.Spec.Tags[:]...)
		p.TargetName = targetName

		// 父ID
		p.ParentId = append(p.ParentId, v.Id)
		p.ParentId = append(p.ParentId, pos.Id)

		// 过滤条件继承
		p.Spec.FilterName = v.Spec.FilterName

		pSet = append(pSet, p)
	}

	return pSet
}

func NewCreateRequest(req *CreatePolicyRequest) *CreatePolicyRequest {
	return &CreatePolicyRequest{
		Name:       req.Name,
		TargetId:   req.TargetId,
		Domain:     req.Domain,
		Namespace:  req.Namespace,
		PolicyType: req.PolicyType,
		Active:     req.Active,
		Mention:    req.Mention,
	}
}

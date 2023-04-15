package policy

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

	return &Policy{
		Id:       xid.New().String(),
		Enabled:  true,
		CreateAt: time.Now().UnixMilli(),
		Spec:     req,
	}, nil
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
	req.TargetId = qs.Get("target_id")

	//lt := qs.Get("label_type")
	//req.LabelType, _ = ParseLABEL_TYPEFromString(lt)

	rids := qs.Get("policy_ids")
	if rids != "" {
		//index := strings.Index(rids, ",")
		//if index == -1 {
		//	rids += ","
		//}
		req.PolicyIds = strings.Split(rids, ",")
	}

	req.Page = page
	return req
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

func NewQueryPolicyRequest(targetId string) *QueryPolicyRequest {
	return &QueryPolicyRequest{
		Page:     request.NewPageRequest(100, 1),
		TargetId: targetId,
	}
}

func NewDeletePolicyRequest() *DeletePolicyRequest {
	return &DeletePolicyRequest{
		PolicyIds: []string{},
	}
}

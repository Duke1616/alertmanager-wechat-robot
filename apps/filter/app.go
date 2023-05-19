package filter

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/policy"
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
	"github.com/rs/xid"
	"time"
)

const (
	AppName = "filter"
)

var (
	validate = validator.New()
)

func (req *CreateFilterRequest) Validate() error {
	return validate.Struct(req)
}

func NewFilter(req *CreateFilterRequest) (*Filter, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &Filter{
		Id:       xid.New().String(),
		Enabled:  true,
		CreateAt: time.Now().UnixMilli(),
		Spec:     req,
		Meta:     map[string]string{},
	}, nil
}

func (resp *Filter) AddMeta(tags []*policy.Tag) {
	for _, v := range tags {
		resp.Meta[v.Label] = v.Value
	}
}

func NewCreateFilter(name string, tags []*policy.Tag) *CreateFilterRequest {
	return &CreateFilterRequest{
		Name: name,
		Tags: tags,
	}
}

func NewFilterSet() *FilterSet {
	return &FilterSet{
		Items: []*Filter{},
	}
}

func (s *FilterSet) Add(item *Filter) {
	s.Total++
	s.Items = append(s.Items, item)
}

func NewDefaultFilter() *Filter {
	return &Filter{}
}

func NewDefaultQueryFilterRequest() *QueryFilterRequest {
	return &QueryFilterRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

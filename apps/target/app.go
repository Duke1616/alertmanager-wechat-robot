package target

import (
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
	"github.com/rs/xid"
	"net/http"
	"time"
)

const (
	AppName = "target"
)

var (
	validate = validator.New()
)

func (req *Target) Validate() error {
	return validate.Struct(req)
}

func NewTarget(req *Target) (*Target, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	t := &Target{
		Id:       xid.New().String(),
		CreateAt: time.Now().UnixMilli(),
		Name:     req.Name,
		Url:      req.Url,
		Secret:   req.Secret,
	}

	return t, nil
}

func (req *DescribeTargetRequest) Validate() error {
	return validate.Struct(req)
}

func NewDefaultTarget() *Target {
	return &Target{}
}

// NewDescribeTargetRequestById 查询详情请求
func NewDescribeTargetRequestById(id string) *DescribeTargetRequest {
	return &DescribeTargetRequest{
		DescribeBy: DESCRIBE_BY_ID,
		Id:         id,
	}
}

// NewDescribeTargetRequestByName 查询详情请求
func NewDescribeTargetRequestByName(name string) *DescribeTargetRequest {
	return &DescribeTargetRequest{
		DescribeBy: DESCRIBE_BY_NAME,
		Name:       name,
	}
}

func NewQueryTargetRequestFromHTTP(r *http.Request) *QueryTargetRequest {
	page := request.NewPageRequestFromHTTP(r)

	qs := r.URL.Query()
	req := NewDefaultQueryTargetRequest()
	req.Url = qs.Get("url")
	req.Secret = qs.Get("secret")

	req.Page = page
	return req
}

func NewDefaultQueryTargetRequest() *QueryTargetRequest {
	return &QueryTargetRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func (req *QueryTargetRequest) Validate() error {
	return validate.Struct(req)
}

func NewTargetSet() *TargetSet {
	return &TargetSet{
		Items: []*Target{},
	}
}

func (s *TargetSet) Add(item *Target) {
	s.Total++
	s.Items = append(s.Items, item)
}

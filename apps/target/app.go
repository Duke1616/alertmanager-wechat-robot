package target

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/xid"
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
		Rule:     req.Rule,
		Url:      req.Url,
		Secret:   req.Secret,
	}

	return t, nil
}

func (req *DescribeTargetRequest) Validate() error {
	return validate.Struct(req)
}

func NewDefaultTarget() *Target {
	return &Target{
		Rule: []*Rule{},
	}
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

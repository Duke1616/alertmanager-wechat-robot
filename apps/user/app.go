package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
	"github.com/rs/xid"
	"net/http"
	"time"
)

const (
	AppName = "user"
)

var (
	validate = validator.New()
)

func (req *CreateUserRequest) Validate() error {
	return validate.Struct(req)
}

func NewUser(req *CreateUserRequest) (*User, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	t := &User{
		Id:       xid.New().String(),
		CreateAt: time.Now().UnixMilli(),
		Spec:     req,
	}

	return t, nil
}

// NewDescribeUserRequestById 查询详情请求
func NewDescribeUserRequestById(id string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeBy: DESCRIBE_BY_ID,
		Id:         id,
	}
}

// NewDescribeUserRequestByName 查询详情请求
func NewDescribeUserRequestByName(name string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeBy: DESCRIBE_BY_NAME,
		Name:       name,
	}
}

func NewQueryUserRequestFromHTTP(r *http.Request) *QueryUserRequest {
	page := request.NewPageRequestFromHTTP(r)

	qs := r.URL.Query()
	req := NewDefaultQueryTargetRequest()

	req.TargetId = qs.Get("target_id")

	req.Page = page
	return req
}

func NewDefaultQueryTargetRequest() *QueryUserRequest {
	return &QueryUserRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func (req *QueryUserRequest) Validate() error {
	return validate.Struct(req)
}

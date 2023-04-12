package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
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
		UserType: USER_TYPE_STAFF,
	}

	return t, nil
}

func NewDefaultUser() *User {
	return &User{}
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

	//req.TargetId = qs.Get("target_id")

	ids := qs.Get("target_id")
	if ids != "" {
		index := strings.Index(ids, ",")
		if index == -1 {
			ids += ","
		}
		req.TargetIds = strings.Split(ids, ",")
	}

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

func (req *DescribeUserRequest) Validate() error {
	return validate.Struct(req)
}

func NewUserSet() *UserSet {
	return &UserSet{
		Items: []*User{},
	}
}

func (s *UserSet) Add(item *User) {
	s.Total++
	s.Items = append(s.Items, item)
}

func NewHashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func CheckPassword(password, haPass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(haPass), []byte(password))
	if err != nil {
		return exception.NewUnauthorized("user or password not connect")
	}

	return nil
}

func NewValidate(user, password string) *ValidateRequest {
	return &ValidateRequest{
		Name:     user,
		Password: password,
	}
}

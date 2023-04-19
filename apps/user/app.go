package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/imdario/mergo"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	pb_request "github.com/infraboard/mcube/pb/request"
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

// NewDescribeUserRequestByName 查询详情请求
func NewDescribeUserRequestByWechatId(wechatId string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeBy: DESCRIBE_BY_WECHAT_ID,
		Name:       wechatId,
	}
}

func NewQueryUserRequestFromHTTP(r *http.Request) *QueryUserRequest {
	page := request.NewPageRequestFromHTTP(r)

	qs := r.URL.Query()
	req := NewDefaultQueryUserRequest()

	//req.TargetId = qs.Get("target_id")

	//tnames := qs.Get("target_names")
	//if tnames != "" {
	//	index := strings.Index(tnames, ",")
	//	if index == -1 {
	//		tnames += ","
	//	}
	//	req.TargetNames = strings.Split(tnames, ",")
	//}

	tids := qs.Get("target_ids")
	if tids != "" {
		//index := strings.Index(tids, ",")
		//if index == -1 {
		//	tids += ","
		//}
		req.TargetNames = strings.Split(tids, ",")
	}

	uids := qs.Get("user_ids")
	if uids != "" {
		req.UserIds = strings.Split(uids, ",")
	}

	req.Page = page
	return req
}

func NewDefaultQueryUserRequest() *QueryUserRequest {
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

func (i *User) Update(req *UpdateUserRequest) {
	i.UpdateAt = time.Now().UnixMilli()
	i.Spec.Profile = req.Profile
	i.Spec.WechatId = req.WechatId
	i.Spec.TargetIds = req.TargetIds
}

func (i *User) Patch(req *UpdateUserRequest) error {
	i.UpdateAt = time.Now().UnixMicro()
	err := mergo.MergeWithOverwrite(i.Spec.Profile, req.Profile)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserSet) HasUser(userId string) bool {
	for i := range s.Items {
		if s.Items[i].Id == userId {
			return true
		}
	}

	return false
}

func (s *UserSet) UserIds() (uids []string) {
	for i := range s.Items {
		uids = append(uids, s.Items[i].Id)
	}

	return
}

func NewPutUserRequest(userId string) *UpdateUserRequest {
	return &UpdateUserRequest{
		UserId:     userId,
		UpdateMode: pb_request.UpdateMode_PUT,
		Profile:    &Profile{},
	}
}

func NewPatchUserRequest(userId string) *UpdateUserRequest {
	return &UpdateUserRequest{
		UserId:     userId,
		UpdateMode: pb_request.UpdateMode_PATCH,
		Profile:    &Profile{},
	}
}

func NewDeleteUserRequest() *DeleteUserRequest {
	return &DeleteUserRequest{
		UserIds: []string{},
	}
}

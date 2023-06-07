package setting

import (
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/request"
	"github.com/rs/xid"
	"net/http"
	"time"
)

const (
	AppName = "setting"
)

var (
	validate = validator.New()
)

func (req *CreateSettingRequest) Validate() error {
	return validate.Struct(req)
}

func NewSetting(req *CreateSettingRequest) (*Setting, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	t := &Setting{
		Id:       xid.New().String(),
		CreateAt: time.Now().UnixMilli(),
		Spec:     req,
	}

	return t, nil
}

func NewDefaultSetting() *Setting {
	return &Setting{}
}

// NewDescribeSettingRequestById 查询详情请求
func NewDescribeSettingRequestById(id string) *DescribeSettingRequest {
	return &DescribeSettingRequest{
		DescribeBy: DESCRIBE_BY_ID,
		Id:         id,
	}
}

// NewDescribeSettingRequestByName 查询详情请求
func NewDescribeSettingRequestByName(name string) *DescribeSettingRequest {
	return &DescribeSettingRequest{
		DescribeBy: DESCRIBE_BY_NAME,
		Name:       name,
	}
}

func NewQuerySettingRequestFromHTTP(r *http.Request) *QuerySettingRequest {
	page := request.NewPageRequestFromHTTP(r)

	//qs := r.URL.Query()
	req := NewDefaultQuerySettingRequest()

	req.Page = page
	return req
}

func NewDefaultQuerySettingRequest() *QuerySettingRequest {
	return &QuerySettingRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func NewSettingSet() *SettingSet {
	return &SettingSet{
		Items: []*Setting{},
	}
}

func (s *SettingSet) Add(item *Setting) {
	s.Total++
	s.Items = append(s.Items, item)
}

func NewSettingHistory() *Setting {
	return &Setting{}
}

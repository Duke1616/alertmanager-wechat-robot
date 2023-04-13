package target

import (
	"github.com/go-playground/validator/v10"
	"github.com/imdario/mergo"
	"github.com/infraboard/mcube/http/request"
	pb_request "github.com/infraboard/mcube/pb/request"
	"github.com/rs/xid"
	"net/http"
	"strings"
	"time"
)

const (
	AppName = "target"
)

var (
	validate = validator.New()
)

func (req *CreateTargetRequest) Validate() error {
	return validate.Struct(req)
}

func NewTarget(req *CreateTargetRequest) (*Target, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	t := &Target{
		Id:       xid.New().String(),
		CreateAt: time.Now().UnixMilli(),
		Spec:     req,
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

	tids := qs.Get("target_ids")
	if tids != "" {
		//index := strings.Index(tids, ",")
		//if index == -1 {
		//	tids += ","
		//}
		req.TargetIds = strings.Split(tids, ",")
	}

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

func (s *TargetSet) HasTarget(targetId string) bool {
	for i := range s.Items {
		if s.Items[i].Id == targetId {
			return true
		}
	}

	return false
}

func (s *TargetSet) TargetIds() (tids []string) {
	for i := range s.Items {
		tids = append(tids, s.Items[i].Id)
	}

	return
}

func (i *Target) Update(req *UpdateTargetRequest) {
	i.UpdateAt = time.Now().UnixMicro()
	i.Spec.Url = req.Url
	i.Spec.Secret = req.Secret
	i.Spec.Name = req.Name
	i.Spec.Description = req.Description
}

func (i *Target) Patch(req *UpdateTargetRequest) error {
	i.UpdateAt = time.Now().UnixMicro()
	err := mergo.MergeWithOverwrite(i.Spec, req)
	if err != nil {
		return err
	}

	return nil
}

func NewPutTargetRequest(targetId string) *UpdateTargetRequest {
	return &UpdateTargetRequest{
		TargetId:   targetId,
		UpdateMode: pb_request.UpdateMode_PUT,
	}
}

func NewPatchTargetRequest(targetId string) *UpdateTargetRequest {
	return &UpdateTargetRequest{
		TargetId:   targetId,
		UpdateMode: pb_request.UpdateMode_PATCH,
	}
}

func NewDeleteTargetRequest() *DeleteTargetRequest {
	return &DeleteTargetRequest{
		TargetIds: []string{},
	}
}

package setting

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/xid"
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

package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/setting"
	app "github.com/Duke1616/alertmanager-wechat-robot/register"
	"github.com/infraboard/mcube/http/label"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	h = &handler{}
)

type handler struct {
	service setting.RPCServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(setting.AppName)
	h.service = app.GetGrpcApp(setting.AppName).(setting.RPCServer)
	return nil
}

func (h *handler) Name() string {
	return setting.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"配置信息"}

	ws.Route(ws.POST("/").To(h.CreateSetting).
		Doc("创建配置").
		Metadata(label.Resource, setting.AppName).
		Metadata(label.Auth, false).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(setting.CreateSettingRequest{}).
		Writes(setting.Setting{}))
}

func init() {
	app.RegistryRestfulApp(h)
}

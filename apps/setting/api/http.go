package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/history"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/setting"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/target"
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

	ws.Route(ws.GET("/{id}").To(h.DescribeSetting).
		Doc("查询配置详情信息").
		Metadata(label.Resource, target.AppName).
		Metadata(label.Auth, false).
		Param(ws.PathParameter("id", "setting id").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(setting.Setting{}).
		Returns(200, "OK", setting.Setting{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/").To(h.QuerySetting).
		Doc("查询告警历史信息").
		Metadata(label.Resource, history.AppName).
		Metadata(label.Auth, true).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("page_size", "每一页的数据条数").DataType("uint64")).
		Param(ws.QueryParameter("page_number", "请求的第页数").DataType("uint64")).
		Param(ws.QueryParameter("offset", "偏移").DataType("int64")).
		Writes(setting.SettingSet{}).
		Returns(200, "OK", setting.SettingSet{}).
		Returns(404, "Not Found", nil))
}

func init() {
	app.RegistryRestfulApp(h)
}

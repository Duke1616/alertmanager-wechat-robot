package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/dashboard"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/policy"
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
	service dashboard.RPCServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(dashboard.AppName)
	h.service = app.GetGrpcApp(dashboard.AppName).(dashboard.RPCServer)
	return nil
}

func (h *handler) Name() string {
	return dashboard.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"dashboard模块"}

	ws.Route(ws.GET("/").To(h.QueryDashboard).
		Doc("根据条件查询策略").
		Metadata(label.Resource, policy.AppName).
		Metadata(label.Auth, true).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("page_size", "每一页的数据条数").DataType("uint64")).
		Param(ws.QueryParameter("page_number", "请求的第页数").DataType("uint64")).
		Param(ws.QueryParameter("offset", "偏移").DataType("int64")).
		Param(ws.QueryParameter("target_name", "查询在target_name下的所有policy规则").DataType("string")).
		Writes(policy.PolicySet{}).
		Returns(200, "OK", policy.PolicySet{}).
		Returns(404, "Not Found", nil))
}

func init() {
	app.RegistryRestfulApp(h)
}

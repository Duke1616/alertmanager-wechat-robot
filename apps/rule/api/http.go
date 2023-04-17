package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/rule"
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
	service rule.RPCServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(rule.AppName)
	h.service = app.GetGrpcApp(rule.AppName).(rule.RPCServer)
	return nil
}

func (h *handler) Name() string {
	return rule.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"告警规则模块"}

	ws.Route(ws.POST("/sync").To(h.SyncRule).
		Doc("同步规则").
		Metadata(label.Resource, rule.AppName).
		Metadata(label.Auth, false).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(rule.SyncRuleRequest{}).
		Writes(rule.Rule{}))

	ws.Route(ws.GET("/group").To(h.QueryGroup).
		Doc("查询告警配置组信息").
		Metadata(label.Resource, rule.AppName).
		Metadata(label.Auth, true).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("page_size", "每一页的数据条数").DataType("uint64")).
		Param(ws.QueryParameter("page_number", "请求的第页数").DataType("uint64")).
		Param(ws.QueryParameter("offset", "偏移").DataType("int64")).
		Param(ws.QueryParameter("url", "webhook地址").DataType("string")).
		Writes(rule.GroupSet{}).
		Returns(200, "OK", rule.GroupSet{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/").To(h.QueryRule).
		Doc("查询告警配置策略信息").
		Metadata(label.Resource, rule.AppName).
		Metadata(label.Auth, true).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("page_size", "每一页的数据条数").DataType("uint64")).
		Param(ws.QueryParameter("page_number", "请求的第页数").DataType("uint64")).
		Param(ws.QueryParameter("offset", "偏移").DataType("int64")).
		Param(ws.QueryParameter("name", "告警名称").DataType("string")).
		Param(ws.QueryParameter("group_name", "告警组").DataType("string")).
		Param(ws.QueryParameter("service_name", "告警服务").DataType("string")).
		Param(ws.QueryParameter("level", "告警级别").DataType("string")).
		Writes(rule.RuleSet{}).
		Returns(200, "OK", rule.RuleSet{}).
		Returns(404, "Not Found", nil))

}

func init() {
	app.RegistryRestfulApp(h)
}

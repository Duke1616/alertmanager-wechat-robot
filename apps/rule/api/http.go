package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/rule"
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
	tags := []string{"规则模块"}

	ws.Route(ws.POST("/").To(h.CreateRule).
		Doc("创建规则").
		Metadata(label.Resource, rule.AppName).
		Metadata(label.Auth, true).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(rule.CreateRuleRequest{}).
		Writes(rule.Rule{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeRule).
		Doc("查询规则详细信息").
		Metadata(label.Resource, rule.AppName).
		Metadata(label.Auth, true).
		Param(ws.PathParameter("id", "identifier of the rule").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(rule.Rule{}).
		Returns(200, "OK", rule.Rule{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/").To(h.QueryRule).
		Doc("条件查询规则").
		Metadata(label.Resource, rule.AppName).
		Metadata(label.Auth, true).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("page_size", "每一页的数据条数").DataType("uint64")).
		Param(ws.QueryParameter("page_number", "请求的第页数").DataType("uint64")).
		Param(ws.QueryParameter("offset", "偏移").DataType("int64")).
		Param(ws.QueryParameter("target_names", "查询在target_names下的所有rule规则").DataType("string")).
		Writes(rule.RuleSet{}).
		Returns(200, "OK", rule.RuleSet{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.DELETE("/{id}").To(h.DeleteRule).
		Doc("删除规则信息").
		Metadata(label.Resource, target.AppName).
		Metadata(label.Auth, true).
		Param(ws.PathParameter("id", "rule id").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags))
}

func init() {
	app.RegistryRestfulApp(h)
}

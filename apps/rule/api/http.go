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
	tags := []string{"策略信息"}

	ws.Route(ws.POST("/").To(h.CreateRule).
		Doc("创建策略").
		Metadata(label.Resource, rule.AppName).
		Metadata(label.Auth, true).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(rule.CreateRuleRequest{}).
		Writes(rule.Rule{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeRule).
		Doc("查询策略详细信息").
		Metadata(label.Resource, rule.AppName).
		Metadata(label.Auth, true).
		Param(ws.PathParameter("id", "identifier of the rule").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(target.Target{}).
		Returns(200, "OK", rule.Rule{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/").To(h.QueryRule).
		Doc("查询策略").
		Metadata(label.Resource, rule.AppName).
		Metadata(label.Auth, true).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(rule.QueryRuleRequest{}).
		Writes(rule.RuleSet{}).
		Returns(200, "OK", rule.RuleSet{}).
		Returns(404, "Not Found", nil))
}

func init() {
	app.RegistryRestfulApp(h)
}

package api

import (
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
	service policy.RPCServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(policy.AppName)
	h.service = app.GetGrpcApp(policy.AppName).(policy.RPCServer)
	return nil
}

func (h *handler) Name() string {
	return policy.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"策略模块"}

	ws.Route(ws.POST("/").To(h.CreatePolicy).
		Doc("创建策略").
		Metadata(label.Resource, policy.AppName).
		Metadata(label.Auth, true).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(policy.CreatePolicyRequest{}).
		Writes(policy.Policy{}))

	ws.Route(ws.GET("/{id}").To(h.DescribePolicy).
		Doc("查询策略详细信息").
		Metadata(label.Resource, policy.AppName).
		Metadata(label.Auth, true).
		Param(ws.PathParameter("id", "identifier of the policy").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(policy.Policy{}).
		Returns(200, "OK", policy.Policy{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/").To(h.QueryPolicy).
		Doc("根据条件查询策略").
		Metadata(label.Resource, policy.AppName).
		Metadata(label.Auth, true).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("page_size", "每一页的数据条数").DataType("uint64")).
		Param(ws.QueryParameter("page_number", "请求的第页数").DataType("uint64")).
		Param(ws.QueryParameter("offset", "偏移").DataType("int64")).
		Param(ws.QueryParameter("target_name", "查询在target_name下的所有policy规则").DataType("string")).
		Param(ws.QueryParameter("create_type", "根据创建策略所有policy规则").DataType("string")).
		Param(ws.QueryParameter("policy_type", "根据策略类型所有policy规则").DataType("string")).
		Writes(policy.PolicySet{}).
		Returns(200, "OK", policy.PolicySet{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.DELETE("/{id}").To(h.DeletePolicy).
		Doc("删除策略信息").
		Metadata(label.Resource, policy.AppName).
		Metadata(label.Auth, true).
		Param(ws.PathParameter("id", "policy id").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags))
}

func init() {
	app.RegistryRestfulApp(h)
}

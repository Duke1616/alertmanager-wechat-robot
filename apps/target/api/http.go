package api

import (
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
	service target.RPCServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(target.AppName)
	h.service = app.GetGrpcApp(target.AppName).(target.RPCServer)
	return nil
}

func (h *handler) Name() string {
	return target.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"群组信息"}

	ws.Route(ws.POST("/").To(h.CreateTarget).
		Doc("创建消息通知群组信息").
		Metadata(label.Resource, target.AppName).
		Metadata(label.Auth, true).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(target.Target{}).
		Writes(target.Target{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeTarget).
		Doc("查询群组详细信息").
		Metadata(label.Resource, target.AppName).
		Metadata(label.Auth, true).
		Param(ws.PathParameter("id", "identifier of the target").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(target.Target{}).
		Returns(200, "OK", target.Target{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/").To(h.QueryTarget).
		Doc("查询消息群组信息").
		Metadata(label.Resource, target.AppName).
		Metadata(label.Auth, true).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(target.QueryTargetRequest{}).
		Writes(target.TargetSet{}).
		Returns(200, "OK", target.TargetSet{}).
		Returns(404, "Not Found", nil))
}

func init() {
	app.RegistryRestfulApp(h)
}

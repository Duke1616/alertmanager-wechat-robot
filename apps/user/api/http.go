package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/target"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/user"
	app "github.com/Duke1616/alertmanager-wechat-robot/register"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	h = &handler{}
)

type handler struct {
	service user.RPCServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(user.AppName)
	h.service = app.GetGrpcApp(user.AppName).(user.RPCServer)
	return nil
}

func (h *handler) Name() string {
	return target.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"用户信息"}

	ws.Route(ws.POST("/").To(h.CreateUser).
		Doc("创建用户").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(user.User{}).
		Writes(user.User{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeUser).
		Doc("查询用户详情信息").
		Param(ws.PathParameter("id", "identifier of the user").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(user.DescribeUserRequest{}).
		Writes(user.User{}).
		Returns(200, "OK", user.User{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/").To(h.QueryUser).
		Doc("查询条件匹配用户信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(user.QueryUserRequest{}).
		Writes(user.UserSet{}).
		Returns(200, "OK", user.UserSet{}).
		Returns(404, "Not Found", nil))
}

func init() {
	app.RegistryRestfulApp(h)
}

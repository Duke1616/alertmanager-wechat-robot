package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/user"
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
	service user.RPCServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(user.AppName)
	h.service = app.GetGrpcApp(user.AppName).(user.RPCServer)
	return nil
}

func (h *handler) Name() string {
	return user.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"用户模块"}

	ws.Route(ws.POST("/").To(h.CreateUser).
		Doc("创建用户").
		Metadata(label.Resource, user.AppName).
		Metadata(label.Auth, false).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(user.CreateUserRequest{}).
		Writes(user.User{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeUser).
		Doc("查询用户详情信息").
		Metadata(label.Resource, user.AppName).
		Metadata(label.Auth, true).
		Param(ws.PathParameter("id", "用户id").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(user.User{}).
		Returns(200, "OK", user.User{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/").To(h.QueryUser).
		Doc("查询条件匹配用户信息").
		Metadata(label.Resource, user.AppName).
		Metadata(label.Auth, true).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("page_size", "每一页的数据条数").DataType("uint64")).
		Param(ws.QueryParameter("page_number", "请求的第页数").DataType("uint64")).
		Param(ws.QueryParameter("offset", "偏移").DataType("int64")).
		Param(ws.QueryParameter("user_ids", "查询user_ids的用户，以逗号分割传递").DataType("string")).
		Param(ws.QueryParameter("target_names", "查询存在target中的用户，以逗号分割传递").DataType("string")).
		Writes(user.UserSet{}).
		Returns(200, "OK", user.UserSet{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{id}").To(h.PutUser).
		Doc("修改用户信息").
		Metadata(label.Resource, user.AppName).
		Metadata(label.Auth, true).
		Param(ws.PathParameter("id", "用户id").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(user.UpdateUserRequest{}))

	//ws.Route(ws.PATCH("/{id}").To(h.PatchUser).
	//	Doc("修改用户信息").
	//	Metadata(label.Resource, user.AppName).
	//	Metadata(label.Auth, true).
	//	Param(ws.PathParameter("id", "用户id").DataType("string")).
	//	Metadata(restfulspec.KeyOpenAPITags, tags).
	//	Reads(user.UpdateUserRequest{}))

	ws.Route(ws.DELETE("/{id}").To(h.DeleteUser).
		Doc("删除用户").
		Metadata(label.Resource, user.AppName).
		Metadata(label.Auth, true).
		Param(ws.PathParameter("id", "用户id").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags))
}

func init() {
	app.RegistryRestfulApp(h)
}

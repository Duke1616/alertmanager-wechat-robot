package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/token"
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
	service token.RPCServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(token.AppName)
	h.service = app.GetGrpcApp(token.AppName).(token.RPCServer)
	return nil
}

func (h *handler) Name() string {
	return token.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"token模块"}

	ws.Route(ws.POST("/issue").To(h.IssueToken).
		Doc("创建token").
		Metadata(label.Resource, token.AppName).
		Metadata(label.Auth, false).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(token.IssueTokenRequest{}).
		Writes(token.Token{}))

	ws.Route(ws.GET("/validate").To(h.ValidateToken).
		Doc("验证token").
		Metadata(label.Resource, token.AppName).
		Metadata(label.Auth, false).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("access_token", "token信息").DataType("string")).
		Writes(token.Token{}).
		Returns(200, "OK", token.Token{}).
		Returns(404, "Not Found", nil))
}

func init() {
	app.RegistryRestfulApp(h)
}

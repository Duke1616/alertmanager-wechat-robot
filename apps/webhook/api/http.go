package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/webhook"
	app "github.com/Duke1616/alertmanager-wechat-robot/register"
	"github.com/infraboard/mcube/http/label"
	"google.golang.org/protobuf/types/known/emptypb"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	h = &handler{}
)

type handler struct {
	service webhook.RPCServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(webhook.AppName)
	h.service = app.GetGrpcApp(webhook.AppName).(webhook.RPCServer)
	return nil
}

func (h *handler) Name() string {
	return webhook.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"转换信息"}

	ws.Route(ws.POST("/send").To(h.TransformToMarkdown).
		Doc("发送wechat robot进行报警").
		Metadata(label.Resource, webhook.AppName).
		Metadata(label.Auth, false).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(webhook.TransData{}).
		Writes(emptypb.Empty{}))
}

func init() {
	app.RegistryRestfulApp(h)
}

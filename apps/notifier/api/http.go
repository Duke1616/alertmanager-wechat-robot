package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/notifier"
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
	service notifier.RPCServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(notifier.AppName)
	h.service = app.GetGrpcApp(notifier.AppName).(notifier.RPCServer)
	return nil
}

func (h *handler) Name() string {
	return notifier.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{""}

	ws.Route(ws.POST("/send").To(h.SendWechatRobot).
		Doc("发送wechat robot进行报警").
		Metadata(label.Resource, notifier.AppName).
		Metadata(label.Auth, false).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(notifier.Notification{}).
		Writes(emptypb.Empty{}))
}

func init() {
	app.RegistryRestfulApp(h)
}

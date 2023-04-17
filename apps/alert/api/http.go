package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/alert"
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
	service alert.RPCServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(alert.AppName)
	h.service = app.GetGrpcApp(alert.AppName).(alert.RPCServer)
	return nil
}

func (h *handler) Name() string {
	return alert.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"告警历史模块"}

	ws.Route(ws.GET("/").To(h.QueryAlertInformation).
		Doc("查询告警历史信息").
		Metadata(label.Resource, alert.AppName).
		Metadata(label.Auth, true).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("page_size", "每一页的数据条数").DataType("uint64")).
		Param(ws.QueryParameter("page_number", "请求的第页数").DataType("uint64")).
		Param(ws.QueryParameter("offset", "偏移").DataType("int64")).
		Param(ws.QueryParameter("type", "告警类型").DataType("string")).
		Param(ws.QueryParameter("service_name", "告警服务").DataType("string")).
		Param(ws.QueryParameter("level", "告警级别").DataType("string")).
		Param(ws.QueryParameter("alert_name", "告警名称").DataType("string")).
		Param(ws.QueryParameter("others", "告警通知人").DataType("string")).
		Param(ws.QueryParameter("target_name", "告警接收群组").DataType("string")).
		Param(ws.QueryParameter("instance_name", "告警主机").DataType("string")).
		Writes(alert.AlertInformationSet{}).
		Returns(200, "OK", alert.AlertInformationSet{}).
		Returns(404, "Not Found", nil))

}

func init() {
	app.RegistryRestfulApp(h)
}

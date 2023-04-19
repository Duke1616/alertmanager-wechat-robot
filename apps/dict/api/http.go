package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/dict"
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
	log logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(dict.AppName)
	return nil
}

func (h *handler) Name() string {
	return dict.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"枚举字典"}
	ws.Route(ws.GET("/rule_label_type").To(h.RuleLabelType).
		Doc("策略标签枚举").
		Metadata(label.Resource, dict.AppName).
		Metadata(label.Auth, false).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(dict.PolicyLabelType).
		Returns(200, "OK", dict.PolicyLabelType))

	ws.Route(ws.GET("/rule_active").To(h.RuleActive).
		Doc("策略动作枚举").
		Metadata(label.Resource, dict.AppName).
		Metadata(label.Auth, false).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(dict.PolicyActive).
		Returns(200, "OK", dict.PolicyActive))

	ws.Route(ws.GET("/user_type").To(h.UserType).
		Doc("用户类型枚举").
		Metadata(label.Resource, dict.AppName).
		Metadata(label.Auth, false).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(dict.UserType).
		Returns(200, "OK", dict.UserType))
}

func init() {
	app.RegistryRestfulApp(h)
}

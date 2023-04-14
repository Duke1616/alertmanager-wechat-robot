package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/dict"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/rule"
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
		Doc("规则标签").
		Metadata(label.Resource, dict.AppName).
		Metadata(label.Auth, false).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(dict.RuleLabelType).
		Returns(200, "OK", dict.RuleLabelType))

	ws.Route(ws.GET("/rule_active").To(h.RuleActive).
		Doc("规则动作").
		Metadata(label.Resource, dict.AppName).
		Metadata(label.Auth, false).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(rule.RuleSet{}).
		Writes(dict.RuleActive).
		Returns(200, "OK", dict.RuleActive))
}

func init() {
	app.RegistryRestfulApp(h)
}

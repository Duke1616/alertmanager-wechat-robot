package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/dict"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
)

func (h *handler) RuleLabelType(r *restful.Request, w *restful.Response) {
	response.Success(w, dict.RuleLabelType)
}

func (h *handler) RuleActive(r *restful.Request, w *restful.Response) {
	response.Success(w, dict.RuleActive)
}

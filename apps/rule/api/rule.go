package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/rule"
	"github.com/Duke1616/alertmanager-wechat-robot/utils"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
)

func (h *handler) SyncRule(r *restful.Request, w *restful.Response) {
	req := &rule.SyncRuleRequest{}
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.SyncRule(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	utils.Success(w, set, "同步数据成功")
}

func (h *handler) QueryGroup(r *restful.Request, w *restful.Response) {
	req := rule.NewQueryGroupRequestFromHTTP(r.Request)
	ins, err := h.service.QueryGroup(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	utils.Success(w, ins, "查询告警组成功")
}

func (h *handler) QueryRule(r *restful.Request, w *restful.Response) {
	req := rule.NewQueryRuleRequestFromHTTP(r.Request)
	ins, err := h.service.QueryRule(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	utils.Success(w, ins, "查询告警信息成功")
}

package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/rule"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
)

func (h *handler) CreateRule(r *restful.Request, w *restful.Response) {
	req := &rule.CreateRuleRequest{}
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.CreateRule(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) DescribeRule(r *restful.Request, w *restful.Response) {
	req := rule.NewDescribeRuleRequestById(r.PathParameter("id"))

	ins, err := h.service.DescribeRule(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) QueryRule(r *restful.Request, w *restful.Response) {
	req := rule.NewQueryRuleRequestFromHTTP(r.Request)
	ins, err := h.service.QueryRule(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) DeleteRule(r *restful.Request, w *restful.Response) {
	req := rule.NewDeleteTargetRequest()
	req.RuleIds = append(req.RuleIds, r.PathParameter("id"))

	set, err := h.service.DeleteRule(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

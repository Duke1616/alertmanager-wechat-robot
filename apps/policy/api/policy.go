package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/policy"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
)

func (h *handler) CreatePolicy(r *restful.Request, w *restful.Response) {
	req := &policy.CreatePolicyRequest{}
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.CreatePolicy(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) DescribePolicy(r *restful.Request, w *restful.Response) {
	req := policy.NewDescribePolicyRequestById(r.PathParameter("id"))

	ins, err := h.service.DescribePolicy(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) QueryPolicy(r *restful.Request, w *restful.Response) {
	req := policy.NewQueryPolicyRequestFromHTTP(r.Request)
	ins, err := h.service.QueryPolicy(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) DeletePolicy(r *restful.Request, w *restful.Response) {
	req := policy.NewDeletePolicyRequest()
	req.PolicyIds = append(req.PolicyIds, r.PathParameter("id"))

	set, err := h.service.DeletePolicy(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

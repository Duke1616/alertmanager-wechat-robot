package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/target"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
)

func (h *handler) CreateTarget(r *restful.Request, w *restful.Response) {
	req := &target.Target{}
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.CreateTarget(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) DescribeTarget(r *restful.Request, w *restful.Response) {
	req := target.NewDescribeTargetRequestById(r.PathParameter("id"))

	ins, err := h.service.DescribeTarget(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) QueryTarget(r *restful.Request, w *restful.Response) {
	req := target.NewQueryTargetRequestFromHTTP(r.Request)
	ins, err := h.service.QueryTarget(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

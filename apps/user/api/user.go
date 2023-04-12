package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/user"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
)

func (h *handler) CreateUser(r *restful.Request, w *restful.Response) {
	req := &user.CreateUserRequest{}
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.CreateUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) DescribeUser(r *restful.Request, w *restful.Response) {
	req := user.NewDescribeUserRequestById(r.PathParameter("id"))

	ins, err := h.service.DescribeUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) QueryUser(r *restful.Request, w *restful.Response) {
	req := user.NewQueryUserRequestFromHTTP(r.Request)
	ins, err := h.service.QueryUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

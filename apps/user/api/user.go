package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/user"
	"github.com/Duke1616/alertmanager-wechat-robot/utils"
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

	utils.Success(w, set, "创建用户成功")
}

func (h *handler) DescribeUser(r *restful.Request, w *restful.Response) {
	req := user.NewDescribeUserRequestById(r.PathParameter("id"))

	ins, err := h.service.DescribeUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	utils.Success(w, ins, "查看用户详情成功")
}

func (h *handler) QueryUser(r *restful.Request, w *restful.Response) {
	req := user.NewQueryUserRequestFromHTTP(r.Request)
	ins, err := h.service.QueryUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	utils.Success(w, ins, "查询用户成功")
}

func (h *handler) PutUser(r *restful.Request, w *restful.Response) {
	req := user.NewPutUserRequest(r.PathParameter("id"))
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.UpdateUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	utils.Success(w, set, "修改用户成功")
}

func (h *handler) PatchUser(r *restful.Request, w *restful.Response) {
	req := user.NewPatchUserRequest(r.PathParameter("id"))
	if err := r.ReadEntity(req.Profile); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.UpdateUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	utils.Success(w, set, "修改用户成功")
}

func (h *handler) DeleteUser(r *restful.Request, w *restful.Response) {
	req := user.NewDeleteUserRequest()
	req.UserIds = append(req.UserIds, r.PathParameter("id"))

	set, err := h.service.DeleteUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	utils.Success(w, set, "删除用户成功")
}

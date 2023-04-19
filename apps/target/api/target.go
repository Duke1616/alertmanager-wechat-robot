package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/target"
	"github.com/Duke1616/alertmanager-wechat-robot/utils"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
)

func (h *handler) CreateTarget(r *restful.Request, w *restful.Response) {
	req := &target.CreateTargetRequest{}
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.CreateTarget(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	utils.Success(w, set, "创建群组信息成功")
}

func (h *handler) DescribeTarget(r *restful.Request, w *restful.Response) {
	req := target.NewDescribeTargetRequestById(r.PathParameter("id"))

	ins, err := h.service.DescribeTarget(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	utils.Success(w, ins, "查询群组详情成功")
}

func (h *handler) QueryTarget(r *restful.Request, w *restful.Response) {
	req := target.NewQueryTargetRequestFromHTTP(r.Request)
	ins, err := h.service.QueryTarget(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	utils.Success(w, ins, "查询群组信息成功")
}

func (h *handler) PutTarget(r *restful.Request, w *restful.Response) {
	req := target.NewPutTargetRequest(r.PathParameter("id"))
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.UpdateTarget(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	utils.Success(w, set, "修改群组信息成功")
}

func (h *handler) DeleteTarget(r *restful.Request, w *restful.Response) {
	req := target.NewDeleteTargetRequest()
	req.TargetIds = append(req.TargetIds, r.PathParameter("id"))

	set, err := h.service.DeleteTarget(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	utils.Success(w, set, "删除群组信息成功")
}

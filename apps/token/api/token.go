package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/token"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
)

func (h *handler) IssueToken(r *restful.Request, w *restful.Response) {
	req := &token.IssueTokenRequest{}
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.IssueToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) ValidateToken(r *restful.Request, w *restful.Response) {

}

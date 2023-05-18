package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/history"
	"github.com/Duke1616/alertmanager-wechat-robot/utils"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
)

func (h *handler) QueryHistory(r *restful.Request, w *restful.Response) {
	req := history.NewQueryHistoryRequestFromHTTP(r.Request)
	ins, err := h.service.QueryHistory(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	utils.Success(w, ins, "查询告警历史成功")
}

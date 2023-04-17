package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/alert"
	"github.com/Duke1616/alertmanager-wechat-robot/utils"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
)

func (h *handler) QueryAlertInformation(r *restful.Request, w *restful.Response) {
	req := alert.NewQueryAlertRequestFromHTTP(r.Request)
	ins, err := h.service.QueryAlertInformation(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	utils.Success(w, ins, "查询告警历史成功")
}

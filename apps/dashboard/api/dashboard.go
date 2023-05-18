package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/dashboard"
	"github.com/Duke1616/alertmanager-wechat-robot/utils"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
)

func (h *handler) QueryDashboard(r *restful.Request, w *restful.Response) {
	//req := policy.NewQueryPolicyRequestFromHTTP(r.Request)
	req := &dashboard.Request{}
	ins, err := h.service.QueryAlertMessage(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	utils.Success(w, ins, "获取策略成功")
}

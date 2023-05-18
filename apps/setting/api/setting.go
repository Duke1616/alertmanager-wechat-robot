package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/setting"
	"github.com/Duke1616/alertmanager-wechat-robot/utils"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
)

func (h *handler) CreateSetting(r *restful.Request, w *restful.Response) {
	req := &setting.CreateSettingRequest{}

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.CreateSetting(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	utils.Success(w, set, "创建配置成功")
}

func (h *handler) DescribeSetting(r *restful.Request, w *restful.Response) {
	req := setting.NewDescribeSettingRequestById(r.PathParameter("id"))

	ins, err := h.service.DescribeSetting(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	utils.Success(w, ins, "查询配置成功")
}

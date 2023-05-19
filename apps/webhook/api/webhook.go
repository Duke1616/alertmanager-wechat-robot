package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/webhook"
	"github.com/emicklei/go-restful/v3"
	"github.com/golang/protobuf/jsonpb"
	"github.com/infraboard/mcube/http/restful/response"
)

func (h *handler) TransformToMarkdown(r *restful.Request, w *restful.Response) {
	// 读取body信息
	//data, _ := io.ReadAll(r.Request.Body)
	//fmt.Printf("data: %v\n", string(data))

	req := &webhook.TransData{
		Alerts: []*webhook.Alert{},
	}

	qs := r.Request.URL.Query()
	isIndex := qs.Get("filter")
	req.IsFilter = isIndex

	err := jsonpb.Unmarshal(r.Request.Body, req)
	if err != nil {
		h.log.Error(err)
	}

	set, err := h.service.TransformAlert(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/notifier"
	"github.com/emicklei/go-restful/v3"
	"github.com/golang/protobuf/jsonpb"
	"github.com/infraboard/mcube/http/restful/response"
)

func (h *handler) SendWechatRobot(r *restful.Request, w *restful.Response) {
	// 读取body信息
	//data, _ := io.ReadAll(r.Request.Body)
	//fmt.Printf("data: %v\n", string(data))

	name := notifier.NewQueryTargetRequestFromHTTP(r.Request, "name")

	req := &notifier.NotificationInfo{
		Name: name,
		Notification: &notifier.Notification{
			Alerts: []*notifier.Alert{},
		},
	}

	err := jsonpb.Unmarshal(r.Request.Body, req.Notification)
	if err != nil {
		//response.Failed(w, err)
		h.log.Error(err)
	}

	//if err := r.ReadEntity(req.Notification); err != nil {
	//	response.Failed(w, err)
	//	return
	//}

	set, err := h.service.SendWechatRobot(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

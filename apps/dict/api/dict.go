package api

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/dict"
	"github.com/Duke1616/alertmanager-wechat-robot/utils"
	"github.com/emicklei/go-restful/v3"
)

func (h *handler) RuleLabelType(r *restful.Request, w *restful.Response) {
	utils.Success(w, dict.PolicyLabelType, "获取策略类型成功")
}

func (h *handler) RuleActive(r *restful.Request, w *restful.Response) {
	utils.Success(w, dict.PolicyActive, "˙获取策略动作成功")
}

func (h *handler) UserType(r *restful.Request, w *restful.Response) {
	utils.Success(w, dict.UserType, "˙获取用户类型成功")
}

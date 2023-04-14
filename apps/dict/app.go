package dict

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/rule"
	"github.com/Duke1616/alertmanager-wechat-robot/utils"
)

const (
	AppName = "dict"
)

var (
	RuleLabelType = []utils.EnumDescribe{
		{Name: "组标签", Value: rule.LABEL_TYPE_GROUP.String(), Describe: "报警分组的label信息"},
		{Name: "通用标签", Value: rule.LABEL_TYPE_COMMON.String(), Describe: "通用的label信息"},
	}

	RuleActive = []utils.EnumDescribe{
		{Name: "@相关人", Value: rule.ACTIVE_NEWS.String(), Describe: "通知相关人进行报警"},
		{Name: "丢弃报警", Value: rule.ACTIVE_DROP.String(), Describe: "丢弃报警"},
		{Name: "状态转换", Value: rule.ACTIVE_STATUS_TRANS.String(), Describe: "firing转换resolved"},
	}
)

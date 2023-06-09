package dict

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/policy"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/user"
	"github.com/Duke1616/alertmanager-wechat-robot/utils"
)

const (
	AppName = "dict"
)

var (
	PolicyLabelType = []utils.EnumDescribe{
		{Name: "组标签", Value: policy.LABEL_TYPE_GROUP.String(), Describe: "报警分组的label信息"},
		{Name: "通用标签", Value: policy.LABEL_TYPE_COMMON.String(), Describe: "通用的label信息"},
	}

	PolicyActive = []utils.EnumDescribe{
		{Name: "@相关人", Value: policy.ACTIVE_NEWS.String(), Describe: "通知相关人进行报警"},
		{Name: "丢弃报警", Value: policy.ACTIVE_DROP.String(), Describe: "丢弃报警"},
		{Name: "状态转换", Value: policy.ACTIVE_STATUS_TRANS.String(), Describe: "firing转换resolved"},
	}

	UserType = []utils.EnumDescribe{
		{Name: "告警用户", Value: user.USER_TYPE_STAFF.String(), Describe: "员工用户、不可登陆"},
		{Name: "系统用户", Value: user.USER_TYPE_SYSTEM.String(), Describe: "系统创建用户、可以登陆"},
	}
)

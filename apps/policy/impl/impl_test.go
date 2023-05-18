package impl_test

import (
	"context"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/policy"
	"github.com/Duke1616/alertmanager-wechat-robot/conf"
	app "github.com/Duke1616/alertmanager-wechat-robot/register"
	"github.com/infraboard/mcube/logger/zap"

	// 注册所有服务
	_ "github.com/Duke1616/alertmanager-wechat-robot/apps/all"
)

var (
	impl policy.RPCServer
	ctx  = context.Background()
)

func init() {
	DevelopmentSetup()
	impl = app.GetInternalApp(policy.AppName).(policy.RPCServer)
}

func DevelopmentSetup() {
	// 初始化日志实例
	zap.DevelopmentSetup()

	// 初始化配置, 提前配置好/etc/unit_test.env
	err := conf.LoadConfigFromToml("/Users/draken/Desktop/alertmanager-wechat-robot/etc/config.toml")
	if err != nil {
		panic(err)
	}

	// 初始化全局app
	if err := app.InitAllApp(); err != nil {
		panic(err)
	}
}

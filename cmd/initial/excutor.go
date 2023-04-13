package initial

import (
	"context"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/target"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/user"
	"github.com/infraboard/mcube/logger/zap"
	"strings"

	// register
	app "github.com/Duke1616/alertmanager-wechat-robot/register"
	// 注册所有服务
	_ "github.com/Duke1616/alertmanager-wechat-robot/apps/all"
)

// NewExecutorFromCLI 初始化
func NewExecutorFromCLI() (*executor, error) {
	e := newExecutor()

	// debug 开关
	if debug {
		zap.SetLevel(zap.DebugLevel)
	} else {
		zap.SetLevel(zap.ErrorLevel)
	}

	err := survey.AskOne(
		&survey.Input{
			Message: "请输入企业微信地址:",
			Default: "https://qyapi.weixin.qq.com",
		},
		&e.wechatUrl,
		survey.WithValidator(survey.Required),
	)
	if err != nil {
		return nil, err
	}

	err = survey.AskOne(
		&survey.Input{
			Message: "请输入企业微信机器人密钥:",
			Default: "",
		},
		&e.wechatSecret,
		survey.WithValidator(survey.Required),
	)
	if err != nil {
		return nil, err
	}

	err = survey.AskOne(
		&survey.Input{
			Message: "请输入企业微信名称:",
			Default: "ops",
		},
		&e.wechatName,
		survey.WithValidator(survey.Required),
	)

	err = survey.AskOne(
		&survey.Input{
			Message: "请输入管理员用户名称:",
			Default: "admin",
		},
		&e.username,
		survey.WithValidator(survey.Required),
	)
	if err != nil {
		return nil, err
	}

	var repeatPass string
	err = survey.AskOne(
		&survey.Password{
			Message: "请输入管理员密码:",
		},
		&e.password,
		survey.WithValidator(survey.Required),
	)

	if err != nil {
		return nil, err
	}
	err = survey.AskOne(
		&survey.Password{
			Message: "再次输入管理员密码:",
		},
		&repeatPass,
		survey.WithValidator(survey.Required),
		survey.WithValidator(func(ans interface{}) error {
			if ans.(string) != e.password {
				return fmt.Errorf("两次输入的密码不一致")
			}
			return nil
		}),
	)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func newExecutor() *executor {
	return &executor{
		user:   app.GetGrpcApp(user.AppName).(user.RPCServer),
		target: app.GetGrpcApp(target.AppName).(target.RPCServer),
	}
}

type executor struct {
	wechatUrl    string
	wechatSecret string
	wechatName   string
	username     string
	password     string

	target target.RPCServer
	user   user.RPCServer
}

func (e *executor) InitTarget(ctx context.Context) error {
	req := &target.CreateTargetRequest{}
	req.Url = e.wechatUrl
	req.Secret = e.wechatSecret
	req.Name = e.wechatName
	ins, err := e.target.CreateTarget(ctx, req)
	if err != nil {
		return err
	}

	fmt.Printf("初始化企业微信机器人信息: %9s [成功]", ins.Spec.Name)
	fmt.Println()
	return nil
}

func (e *executor) InitAdminUser(ctx context.Context) error {
	req := &user.CreateUserRequest{}
	req.Password = strings.TrimSpace(e.password)
	req.Name = strings.TrimSpace(e.username)
	req.UserType = user.USER_TYPE_SYSTEM

	u, err := e.user.CreateUser(ctx, req)
	if err != nil {
		return err
	}
	fmt.Printf("初始化系统管理员: %9s [成功]", u.Spec.Name)
	fmt.Println()
	return nil
}

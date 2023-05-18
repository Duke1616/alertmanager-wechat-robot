package impl

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/history"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/notify"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/policy"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/webhook"
	app "github.com/Duke1616/alertmanager-wechat-robot/register"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	col *mongo.Collection
	log logger.Logger
	webhook.UnimplementedRPCServer

	policy  policy.RPCServer
	notify  notify.RPCServer
	history history.RPCServer
}

func (s *service) Config() error {
	s.log = zap.L().Named(s.Name())

	s.policy = app.GetGrpcApp("policy").(policy.RPCServer)
	s.notify = app.GetGrpcApp("notify").(notify.RPCServer)
	s.history = app.GetGrpcApp("history").(history.RPCServer)
	return nil
}

func (s *service) Name() string {
	return webhook.AppName
}

func (s *service) Registry(server *grpc.Server) {
	webhook.RegisterRPCServer(server, svr)
}

func init() {
	app.RegistryInternalApp(svr)
	app.RegistryGrpcApp(svr)
}

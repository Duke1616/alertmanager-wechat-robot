package impl

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/alert"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/notifier"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/rule"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/target"
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
	notifier.UnimplementedRPCServer

	target target.RPCServer
	alert  alert.RPCServer
	rule   rule.RPCServer
}

func (s *service) Config() error {
	s.log = zap.L().Named(s.Name())

	s.target = app.GetGrpcApp("target").(target.RPCServer)
	s.rule = app.GetGrpcApp("rule").(rule.RPCServer)
	s.alert = app.GetGrpcApp("alert").(alert.RPCServer)
	return nil
}

func (s *service) Name() string {
	return notifier.AppName
}

func (s *service) Registry(server *grpc.Server) {
	notifier.RegisterRPCServer(server, svr)
}

func init() {
	app.RegistryInternalApp(svr)
	app.RegistryGrpcApp(svr)
}

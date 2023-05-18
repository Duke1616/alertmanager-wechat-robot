package impl

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/dashboard"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/history"
	"github.com/Duke1616/alertmanager-wechat-robot/conf"
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
	dashboard.UnimplementedRPCServer

	history history.RPCServer
}

func (s *service) Config() error {
	s.log = zap.L().Named(s.Name())
	// 依赖MongoDB的DB对象
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	s.col = db.Collection(history.AppName)
	s.history = app.GetGrpcApp(history.AppName).(history.RPCServer)
	return nil
}

func (s *service) Name() string {
	return dashboard.AppName
}

func (s *service) Registry(server *grpc.Server) {
	dashboard.RegisterRPCServer(server, svr)
}

func init() {
	app.RegistryInternalApp(svr)
	app.RegistryGrpcApp(svr)
}

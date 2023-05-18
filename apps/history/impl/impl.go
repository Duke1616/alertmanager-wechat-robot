package impl

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/history"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/target"
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
	history.UnimplementedRPCServer

	target target.RPCServer
}

func (s *service) Config() error {
	// 依赖MongoDB的DB对象
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	// 获取一个Collection对象, 通过Collection对象 来进行CRUD
	s.col = db.Collection(s.Name())

	s.target = app.GetGrpcApp("target").(target.RPCServer)
	s.log = zap.L().Named(s.Name())

	return nil
}

func (s *service) Name() string {
	return history.AppName
}

func (s *service) Registry(server *grpc.Server) {
	history.RegisterRPCServer(server, svr)
}

func init() {
	app.RegistryInternalApp(svr)
	app.RegistryGrpcApp(svr)
}

package impl

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/rule"
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
	rule  *mongo.Collection
	group *mongo.Collection

	log logger.Logger
	rule.UnimplementedRPCServer
}

func (s *service) Config() error {
	s.log = zap.L().Named(s.Name())
	// 依赖MongoDB的DB对象
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	s.rule = db.Collection(s.Name())
	s.group = db.Collection("group")
	return nil
}

func (s *service) Name() string {
	return rule.AppName
}

func (s *service) Registry(server *grpc.Server) {
	rule.RegisterRPCServer(server, svr)
}

func init() {
	app.RegistryInternalApp(svr)
	app.RegistryGrpcApp(svr)
}

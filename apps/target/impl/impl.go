package impl

import (
	"context"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/user"

	"github.com/Duke1616/alertmanager-wechat-robot/apps/target"
	"github.com/Duke1616/alertmanager-wechat-robot/conf"
	app "github.com/Duke1616/alertmanager-wechat-robot/register"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	col *mongo.Collection
	log logger.Logger
	target.UnimplementedRPCServer

	user user.RPCServer
}

func (s *service) Config() error {
	// 依赖MongoDB的DB对象
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	dc := db.Collection(s.Name())
	indexes := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{Key: "spec.name", Value: bsonx.Int32(-1)},
			},
			Options: options.Index().SetUnique(true),
		},
	}

	_, err = dc.Indexes().CreateMany(context.Background(), indexes)
	if err != nil {
		return err
	}

	s.col = dc

	s.log = zap.L().Named(s.Name())
	s.user = app.GetGrpcApp(user.AppName).(user.RPCServer)
	return nil
}

func (s *service) Name() string {
	return target.AppName
}

func (s *service) Registry(server *grpc.Server) {
	target.RegisterRPCServer(server, svr)
}

func init() {
	app.RegistryInternalApp(svr)
	app.RegistryGrpcApp(svr)
}

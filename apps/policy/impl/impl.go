package impl

import (
	"context"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/filter"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/policy"
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
	policy.UnimplementedRPCServer

	target target.RPCServer
	filter filter.RPCServer
}

func (s *service) Config() error {
	s.log = zap.L().Named(s.Name())
	// 依赖MongoDB的DB对象
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	dc := db.Collection(s.Name())

	indexes := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "spec.priority", Value: bsonx.Int32(-1)}},
		},
		{
			Keys: bsonx.Doc{{Key: "spec.index", Value: bsonx.Int32(-1)}},
		},
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

	s.filter = app.GetGrpcApp("filter").(filter.RPCServer)
	s.target = app.GetGrpcApp("target").(target.RPCServer)
	s.col = db.Collection(s.Name())
	return nil
}

func (s *service) Name() string {
	return policy.AppName
}

func (s *service) Registry(server *grpc.Server) {
	policy.RegisterRPCServer(server, svr)
}

func init() {
	app.RegistryInternalApp(svr)
	app.RegistryGrpcApp(svr)
}

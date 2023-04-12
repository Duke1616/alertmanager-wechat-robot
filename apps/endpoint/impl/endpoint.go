package impl

import (
	"context"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/endpoint"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) RegistryEndpoint(ctx context.Context, req *endpoint.RegistryRequest) (*endpoint.RegistryResponse, error) {
	// 生产该服务的Endpoint
	endpoints := req.Endpoints()

	// 更新已有的记录
	news := make([]interface{}, 0, len(endpoints))
	for i := range endpoints {
		if err := s.col.FindOneAndReplace(ctx, bson.M{"_id": endpoints[i].Id}, endpoints[i]).Err(); err != nil {
			if err == mongo.ErrNoDocuments {
				news = append(news, endpoints[i])
			} else {
				return nil, err
			}
		}
	}

	// 插入新增记录
	if len(news) > 0 {
		if _, err := s.col.InsertMany(ctx, news); err != nil {
			return nil, exception.NewInternalServerError("inserted a service document error, %s", err)
		}
	}

	return nil, nil
}

package impl

import (
	"context"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/setting"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) CreateSetting(ctx context.Context, req *setting.CreateSettingRequest) (*setting.Setting, error) {
	payload, err := setting.NewSetting(req)

	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}
	if _, err = s.col.InsertOne(ctx, payload); err != nil {
		return nil, exception.NewInternalServerError("inserted a setting document error, %s", err)
	}

	return payload, nil
}

func (s *service) DescribeSetting(ctx context.Context, req *setting.DescribeSettingRequest) (*setting.Setting, error) {
	filter := bson.M{}
	switch req.DescribeBy {
	case setting.DESCRIBE_BY_ID:
		filter["_id"] = req.Id
	case setting.DESCRIBE_BY_NAME:
		filter["spec.name"] = req.Name
	}

	resp := setting.NewDefaultSetting()
	if err := s.col.FindOne(ctx, filter).Decode(resp); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("setting %s not found", req)
		}

		return nil, exception.NewInternalServerError("find setting %s error, %s", req.Id, err)
	}

	return resp, nil
}

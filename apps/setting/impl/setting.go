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

func (s *service) QuerySetting(ctx context.Context, req *setting.QuerySettingRequest) (*setting.SettingSet, error) {
	query, err := newQuerySettingRequest(req)
	if err != nil {
		return nil, err
	}

	s.log.Debugf("query setting filter: %s", query.FindFilter())
	resp, err := s.col.Find(ctx, query.FindFilter(), query.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find setting error, error is %s", err)
	}

	set := setting.NewSettingSet()
	// 循环插入数据
	for resp.Next(ctx) {
		ins := setting.NewSettingHistory()
		if err = resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode setting error, error is %s", err)
		}
		set.Add(ins)
	}

	// 计算数量
	count, err := s.col.CountDocuments(ctx, query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get setting count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

package impl

import (
	"context"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/target"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) CreateTarget(ctx context.Context, req *target.Target) (*target.Target, error) {
	t, err := target.NewTarget(req)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}
	if _, err := s.col.InsertOne(ctx, t); err != nil {
		return nil, exception.NewInternalServerError("inserted a target document error, %s", err)
	}

	return t, nil
}

func (s *service) DescribeTarget(ctx context.Context, req *target.DescribeTargetRequest) (*target.Target, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	filter := bson.M{}
	switch req.DescribeBy {
	case target.DESCRIBE_BY_ID:
		filter["_id"] = req.Id
	case target.DESCRIBE_BY_NAME:
		filter["name"] = req.Name
	}

	t := target.NewDefaultTarget()
	if err := s.col.FindOne(ctx, filter).Decode(t); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("target %s not found", req)
		}

		return nil, exception.NewInternalServerError("find target %s error, %s", req.Id, err)
	}

	return t, nil
}

func (s *service) QueryTarget(ctx context.Context, req *target.QueryTargetRequest) (*target.TargetSet, error) {
	query, err := newQueryTargetRequest(req)
	if err != nil {
		return nil, err
	}

	s.log.Debugf("query target filter: %s", query.FindFilter())
	resp, err := s.col.Find(ctx, query.FindFilter(), query.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find target error, error is %s", err)
	}

	set := target.NewTargetSet()
	// 循环插入数据
	for resp.Next(ctx) {
		ins := target.NewDefaultTarget()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode target error, error is %s", err)
		}
		set.Add(ins)
	}

	// 计算数量
	count, err := s.col.CountDocuments(ctx, query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get target count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

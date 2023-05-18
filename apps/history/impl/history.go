package impl

import (
	"context"
	"fmt"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/history"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/target"
	"github.com/infraboard/mcube/exception"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *service) SaveHistory(ctx context.Context, req *history.SaveHistoryRequest) (*emptypb.Empty, error) {
	// 组合数据
	t, err := s.target.DescribeTarget(ctx, target.NewDescribeTargetRequestById(req.Data.TargetId))
	if err != nil {
		s.log.Debug("告警通知target，没有找到")
		return nil, err
	}

	his := req.HistorySet(t.Spec.Name)

	fmt.Println(his)
	if _, err := s.col.InsertMany(ctx, his); err != nil {
		return nil, exception.NewInternalServerError("inserted a history document error, %s", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *service) QueryHistory(ctx context.Context, req *history.QueryHistoryRequest) (*history.HistorySet, error) {
	query, err := newQueryHistoryRequest(req)
	if err != nil {
		return nil, err
	}

	s.log.Debugf("query history filter: %s", query.FindFilter())
	resp, err := s.col.Find(ctx, query.FindFilter(), query.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find history error, error is %s", err)
	}

	set := history.NewHistorySet()
	// 循环插入数据
	for resp.Next(ctx) {
		ins := history.NewDefaultHistory()
		if err = resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode history error, error is %s", err)
		}
		set.Add(ins)
	}

	// 计算数量
	count, err := s.col.CountDocuments(ctx, query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get history count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

package impl

import (
	"context"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/filter"
	"github.com/infraboard/mcube/exception"
)

func (s *service) CreateFilter(ctx context.Context, req *filter.CreateFilterRequest) (*filter.Filter, error) {
	resp, err := filter.NewFilter(req)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	resp.AddMeta(req.Tags)

	if _, err = s.col.InsertOne(ctx, resp); err != nil {
		return nil, exception.NewInternalServerError("inserted a filter document error, %s", err)
	}

	return resp, nil
}

func (s *service) QueryTarget(ctx context.Context, req *filter.QueryFilterRequest) (*filter.FilterSet, error) {
	query, err := newQueryFilterRequest(req)
	if err != nil {
		return nil, err
	}

	s.log.Debugf("query filter : %s", query.FindFilter())
	resp, err := s.col.Find(ctx, query.FindFilter(), query.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find filter error, error is %s", err)
	}

	set := filter.NewFilterSet()
	// 循环插入数据
	for resp.Next(ctx) {
		ins := filter.NewDefaultFilter()
		if err = resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode filter error, error is %s", err)
		}
		set.Add(ins)
	}

	// 计算数量
	count, err := s.col.CountDocuments(ctx, query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get filter count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

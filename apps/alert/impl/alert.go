package impl

import (
	"context"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/alert"
	"github.com/infraboard/mcube/exception"
)

func (s *service) SaveAlertInformation(ctx context.Context, req *alert.SaveAlertRequest) (*alert.SaveAlertResponse, error) {
	// 查询用户名称
	//var users []string
	//
	//for _, wechatId := range req.Mention.Username {
	//	s.log.Info(wechatId)
	//	u, err := s.user.DescribeUser(ctx, user.NewDescribeUserRequestByWechatId(wechatId))
	//	if exception.IsNotFoundError(err) == true {
	//		s.log.Errorf("通过wechat_id: %s查询用户失败", wechatId)
	//	}
	//	if err != nil {
	//		s.log.Errorf("通过wechat_id: %s查询用户失败", wechatId)
	//		return nil, err
	//	}
	//	users = append(users, u.Spec.Name)
	//}

	// 组合数据
	ais := req.AlertInformationS()

	if _, err := s.col.InsertMany(ctx, ais); err != nil {
		return nil, exception.NewInternalServerError("inserted a rule document error, %s", err)
	}

	return &alert.SaveAlertResponse{Message: "ok"}, nil
}

func (s *service) QueryAlertInformation(ctx context.Context, req *alert.QueryAlertRequest) (*alert.AlertInformationSet, error) {
	query, err := newQueryAlertRequest(req)
	if err != nil {
		return nil, err
	}

	s.log.Debugf("query group filter: %s", query.FindFilter())
	resp, err := s.col.Find(ctx, query.FindFilter(), query.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find alert error, error is %s", err)
	}

	set := alert.NewAlertInformationSet()
	// 循环插入数据
	for resp.Next(ctx) {
		ins := alert.NewDefaultAlertInformation()
		if err = resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode alert error, error is %s", err)
		}
		set.Add(ins)
	}

	// 计算数量
	count, err := s.col.CountDocuments(ctx, query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get alert count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

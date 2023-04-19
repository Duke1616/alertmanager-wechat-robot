package impl

import (
	"context"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/target"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/user"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pb/request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.User, error) {
	u, err := user.NewUser(req)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	u.Spec.Password, err = user.NewHashedPassword(req.Password)
	if err != nil {
		return nil, err
	}

	queryReq := target.NewDefaultQueryTargetRequest()
	queryReq.TargetIds = req.TargetIds
	tSet, err := s.target.QueryTarget(ctx, queryReq)
	if err != nil {
		return nil, err
	}

	var ts []string
	for _, v := range tSet.Items {
		ts = append(ts, v.Spec.Name)
	}

	u.TargetNames = ts

	if _, err = s.col.InsertOne(ctx, u); err != nil {
		return nil, exception.NewInternalServerError("inserted a user document error, %s", err)
	}

	return u, nil
}

func (s *service) DescribeUser(ctx context.Context, req *user.DescribeUserRequest) (*user.User, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	filter := bson.M{}
	switch req.DescribeBy {
	case user.DESCRIBE_BY_ID:
		filter["_id"] = req.Id
	case user.DESCRIBE_BY_NAME:
		filter["spec.name"] = req.Name
	case user.DESCRIBE_BY_WECHAT_ID:
		filter["spec.wechat_id"] = req.WechatId
	}

	u := user.NewDefaultUser()
	if err := s.col.FindOne(ctx, filter).Decode(u); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("user %s not found", req)
		}

		return nil, exception.NewInternalServerError("find user %s error, %s", req.Id, err)
	}

	return u, nil
}

func (s *service) QueryUser(ctx context.Context, req *user.QueryUserRequest) (*user.UserSet, error) {
	query, err := newQueryUserRequest(req)
	if err != nil {
		return nil, err
	}

	s.log.Debugf("query user filter: %s", query.FindFilter())
	resp, err := s.col.Find(ctx, query.FindFilter(), query.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find user error, error is %s", err)
	}

	set := user.NewUserSet()
	// 循环插入数据
	for resp.Next(ctx) {
		ins := user.NewDefaultUser()
		if err = resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode user error, error is %s", err)
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

func (s *service) ValidateUser(ctx context.Context, req *user.ValidateRequest) (*user.User, error) {
	if req.Name == "" || req.Password == "" {
		return nil, exception.NewUnauthorized("Account or password does not exist")
	}

	// 检测用户的密码是否正确
	u, err := s.DescribeUser(ctx, user.NewDescribeUserRequestByName(req.Name))
	if err != nil {
		s.log.Errorf("Account or password does not exist")
		return nil, err
	}

	// 判断用户类型是否可以正常登陆
	if u.Spec.UserType != user.USER_TYPE_SYSTEM {
		s.log.Errorf("Account type refused login")
		return nil, exception.NewUnauthorized("Account type refused login")
	}

	err = user.CheckPassword(req.Password, u.Spec.Password)
	if err != nil {
		s.log.Errorf("Account or password does not exist")
		return nil, exception.NewUnauthorized("Account or password does not exist")
	}

	return u, nil
}

func (s *service) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.User, error) {
	ins, err := s.DescribeUser(ctx, user.NewDescribeUserRequestById(req.UserId))
	if err != nil {
		return nil, err
	}

	switch req.UpdateMode {
	case request.UpdateMode_PUT:
		ins.Update(req)
	case request.UpdateMode_PATCH:
		err = ins.Patch(req)
		if err != nil {
			return nil, exception.NewBadRequest("patch error, %s", err)
		}
	}

	if err := s.update(ctx, ins); err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.UserSet, error) {
	// 判断这些要删除的用户是否存在
	queryReq := user.NewDefaultQueryUserRequest()
	queryReq.UserIds = req.UserIds
	set, err := s.QueryUser(ctx, queryReq)
	if err != nil {
		return nil, err
	}

	var noExist []string
	for _, uid := range req.UserIds {
		if !set.HasUser(uid) {
			noExist = append(noExist, uid)
		}
	}
	if len(noExist) > 0 {
		return nil, exception.NewBadRequest("user %v not found", req.UserIds)
	}

	if err = s.delete(ctx, set); err != nil {
		return nil, err
	}
	return set, nil
}

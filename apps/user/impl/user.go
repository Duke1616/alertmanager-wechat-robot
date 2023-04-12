package impl

import (
	"context"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/user"
	"github.com/infraboard/mcube/exception"
)

func (s *service) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.User, error) {
	t, err := user.NewUser(req)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}
	if _, err := s.col.InsertOne(ctx, t); err != nil {
		return nil, exception.NewInternalServerError("inserted a user document error, %s", err)
	}

	return t, nil
}

func (s *service) DescribeUser(ctx context.Context, req *user.DescribeUserRequest) (*user.User, error) {
	return nil, nil
}

func (s *service) QueryUser(ctx context.Context, req *user.QueryUserRequest) (*user.User, error) {
	return nil, nil
}

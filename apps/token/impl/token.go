package impl

import (
	"context"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/token"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/user"
	"github.com/infraboard/mcube/exception"
)

func (s *service) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	t, err := token.NewToken(req)

	validate := user.NewValidate(req.Username, req.Password)
	u, err := s.user.ValidateUser(ctx, validate)
	if err != nil {
		return nil, err
	}

	t.Username = u.Spec.Name
	t.UserId = u.Id

	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	if _, err := s.col.InsertOne(ctx, t); err != nil {
		return nil, exception.NewInternalServerError("inserted a token document error, %s", err)
	}

	return t, nil
}

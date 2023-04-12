package impl

import (
	"context"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/token"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/user"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func (s *service) ValidateToken(ctx context.Context, req *token.ValidateTokenRequest) (*token.Token, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	tk, err := s.DescribeToken(ctx, token.NewDescribeTokenRequestByToken(req.AccessToken))

	if err != nil {
		return nil, exception.NewUnauthorized(err.Error())
	}

	return tk, nil
}

func (s *service) DescribeToken(ctx context.Context, req *token.DescribeTokenRequest) (*token.Token, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	filter := bson.M{}
	switch req.DescribeBy {
	case token.DESCRIBE_BY_ACCESS_TOKEN:
		filter["_id"] = req.AccessToken
	}

	tk := token.NewDefaultToken()
	if err := s.col.FindOne(ctx, filter).Decode(tk); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("user %s not found", req)
		}
		return nil, exception.NewInternalServerError("find token %s error, %s", req.AccessToken, err)
	}

	return tk, nil
}

func (s *service) RemoveToken(ctx context.Context, req *token.RemoveTokenRequest) (*token.Token, error) {
	return nil, nil
}

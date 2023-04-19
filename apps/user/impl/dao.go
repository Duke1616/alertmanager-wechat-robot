package impl

import (
	"context"
	"fmt"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/user"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newQueryUserRequest(req *user.QueryUserRequest) (*queryUserRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &queryUserRequest{
		QueryUserRequest: req}, nil
}

type queryUserRequest struct {
	*user.QueryUserRequest
}

func (r *queryUserRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryUserRequest) FindFilter() bson.M {
	filter := bson.M{}

	if len(r.TargetIds) > 0 {
		filter["spec.target_ids"] = bson.M{"$in": r.TargetIds}
	}
	if len(r.UserIds) > 0 {
		filter["_id"] = bson.M{"$in": r.UserIds}
	}

	return filter
}

func (s *service) update(ctx context.Context, ins *user.User) error {
	if _, err := s.col.UpdateByID(ctx, ins.Id, bson.M{"$set": ins}); err != nil {
		return exception.NewInternalServerError("update user(%s) document error, %s",
			ins.Id, err)
	}

	return nil
}

func (s *service) delete(ctx context.Context, set *user.UserSet) error {
	if set == nil || len(set.Items) == 0 {
		return fmt.Errorf("user is nil")
	}

	var result *mongo.DeleteResult
	var err error
	if len(set.Items) == 1 {
		result, err = s.col.DeleteMany(ctx, bson.M{"_id": set.UserIds()[0]})
	} else {
		result, err = s.col.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": set.UserIds()}})
	}

	if err != nil {
		return exception.NewInternalServerError("delete user(%s) error, %s", set, err)
	}

	if result.DeletedCount == 0 {
		return exception.NewNotFound("user %s not found", set)
	}

	return nil
}

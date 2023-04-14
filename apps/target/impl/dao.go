package impl

import (
	"context"
	"fmt"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/target"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newQueryTargetRequest(req *target.QueryTargetRequest) (*queryTargetRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &queryTargetRequest{
		QueryTargetRequest: req}, nil
}

type queryTargetRequest struct {
	*target.QueryTargetRequest
}

func (r *queryTargetRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryTargetRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Url != "" {
		filter["spec.url"] = r.Url
	}

	if len(r.TargetIds) > 1 {
		filter["_id"] = bson.M{"$in": r.TargetIds}
	}

	return filter
}

func (s *service) update(ctx context.Context, ins *target.Target) error {
	if _, err := s.col.UpdateByID(ctx, ins.Id, bson.M{"$set": ins}); err != nil {
		return exception.NewInternalServerError("update target(%s) document error, %s",
			ins.Id, err)
	}

	return nil
}

func (s *service) delete(ctx context.Context, set *target.TargetSet) error {
	if set == nil || len(set.Items) == 0 {
		return fmt.Errorf("target is nil")
	}

	var result *mongo.DeleteResult
	var err error
	if len(set.Items) == 1 {
		result, err = s.col.DeleteMany(ctx, bson.M{"_id": set.TargetIds()[0]})
	} else {
		result, err = s.col.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": set.TargetIds()}})
	}

	if err != nil {
		return exception.NewInternalServerError("delete target(%s) error, %s", set, err)
	}

	if result.DeletedCount == 0 {
		return exception.NewNotFound("target %s not found", set)
	}

	return nil
}

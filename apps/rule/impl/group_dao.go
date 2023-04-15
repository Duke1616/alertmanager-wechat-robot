package impl

import (
	"context"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/rule"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *service) saveGroups(ctx context.Context, rules *rule.Rules) error {
	groups := rules.GroupSet()

	// 更新group数据
	newGroup := make([]interface{}, 0, len(groups))
	for i := range groups {
		if err := s.group.FindOneAndReplace(ctx, bson.M{"_id": groups[i].Id}, groups[i]).Err(); err != nil {
			if err == mongo.ErrNoDocuments {
				newGroup = append(newGroup, groups[i])
			} else {
				return err
			}
		}
	}

	// 插入新增group记录
	if len(newGroup) > 0 {
		if _, err := s.group.InsertMany(ctx, newGroup); err != nil {
			return exception.NewInternalServerError("inserted a group document error, %s", err)
		}
	}

	return nil
}

func newQueryGroupRequest(req *rule.QueryGroupRequest) (*queryGroupRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &queryGroupRequest{
		QueryGroupRequest: req}, nil
}

type queryGroupRequest struct {
	*rule.QueryGroupRequest
}

func (r *queryGroupRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryGroupRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Name != "" {
		filter["name"] = r.Name
	}

	return filter
}

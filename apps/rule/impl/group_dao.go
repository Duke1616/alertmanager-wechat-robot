package impl

import (
	"context"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/rule"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/setting"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *service) saveGroups(ctx context.Context, rules *rule.Rules, setting *setting.Setting) ([]string, error) {
	groupSet := rules.GroupSet(setting)

	// 删除原有数据
	set, err := s.DeleteGroup(ctx, rule.NewDeleteGroupRequest(setting.Spec.Domain, setting.Spec.Namespace))
	if err != nil {
		return nil, err
	}

	var groupIds []string
	for _, v := range set.Items {
		groupIds = append(groupIds, v.Id)
	}

	// 插入新增group记录
	if _, err := s.group.InsertMany(ctx, groupSet); err != nil {
		return nil, exception.NewInternalServerError("inserted a group document error, %s", err)
	}

	return groupIds, nil
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

	if r.Domain != "" {
		filter["domain"] = r.Domain
	}

	if r.Namespace != "" {
		filter["namespace"] = r.Namespace
	}

	return filter
}

func (s *service) deleteGroup(ctx context.Context, set *rule.GroupSet) error {
	if set == nil || len(set.Items) == 0 {
		return nil
	}

	var result *mongo.DeleteResult
	var err error
	if len(set.Items) == 1 {
		result, err = s.group.DeleteMany(ctx, bson.M{"_id": set.GroupsIds()[0]})
	} else {
		result, err = s.group.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": set.GroupsIds()}})
	}

	if err != nil {
		return exception.NewInternalServerError("delete group(%s) error, %s", set, err)
	}

	if result.DeletedCount == 0 {
		return exception.NewNotFound("group %s not found", set)
	}

	return nil
}

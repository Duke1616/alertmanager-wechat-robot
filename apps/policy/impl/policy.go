package impl

import (
	"context"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/policy"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/target"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *service) CreatePolicy(ctx context.Context, req *policy.CreatePolicyRequest) (*policy.Policy, error) {
	r, err := policy.NewPolicy(req)

	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	if req.Active == policy.ACTIVE_NEWS {
		t, err := s.target.DescribeTarget(ctx, target.NewDescribeTargetRequestById(req.TargetId))
		if err != nil {
			return nil, err
		}

		r.TargetName = t.Spec.Name
		//r.Spec.Mention.Username =
	}

	if _, err = s.col.InsertOne(ctx, r); err != nil {
		return nil, exception.NewInternalServerError("inserted a policy document error, %s", err)
	}

	return r, nil
}

func (s *service) DescribePolicy(ctx context.Context, req *policy.DescribePolicyRequest) (*policy.Policy, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	filter := bson.M{}
	switch req.DescribeBy {
	case policy.DESCRIBE_BY_id:
		filter["_id"] = req.Id
	}

	r := policy.NewDefaultPolicy()
	if err := s.col.FindOne(ctx, filter).Decode(r); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("policy %s not found", req)
		}

		return nil, exception.NewInternalServerError("find policy %s error, %s", req.Id, err)
	}

	return r, nil
}

func (s *service) QueryPolicy(ctx context.Context, req *policy.QueryPolicyRequest) (*policy.PolicySet, error) {
	query, err := newQueryPolicyRequest(req)
	if err != nil {
		return nil, err
	}

	s.log.Debugf("query policy filter: %s", query.FindFilter())
	resp, err := s.col.Find(ctx, query.FindFilter(), query.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find policy error, error is %s", err)
	}

	set := policy.NewPolicySet()
	// 循环插入数据
	for resp.Next(ctx) {
		ins := policy.NewDefaultPolicy()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode policy error, error is %s", err)
		}
		set.Add(ins)
	}

	// 计算数量
	count, err := s.col.CountDocuments(ctx, query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get policy count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

func (s *service) DeletePolicy(ctx context.Context, req *policy.DeletePolicyRequest) (*policy.PolicySet, error) {
	// 判断这些要删除的用户是否存在
	queryReq := policy.NewDefaultQueryPolicyRequest()
	queryReq.PolicyIds = req.PolicyIds
	set, err := s.QueryPolicy(ctx, queryReq)
	if err != nil {
		return nil, err
	}

	var noExist []string
	for _, uid := range req.PolicyIds {
		if !set.HasPolicy(uid) {
			noExist = append(noExist, uid)
		}
	}
	if len(noExist) > 0 {
		return nil, exception.NewBadRequest("policy %v not found", req.PolicyIds)
	}

	if err = s.delete(ctx, set); err != nil {
		return nil, err
	}
	return set, nil
}

func (s *service) UpdatePolicy(ctx context.Context, req *policy.UpdatePolicyRequest) (*policy.Policy, error) {
	return nil, nil
}

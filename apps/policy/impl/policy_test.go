package impl_test

import (
	"github.com/Duke1616/alertmanager-wechat-robot/apps/policy"
	"testing"
)

func TestCreatePolicy(t *testing.T) {
	req := policy.NewCreateRequest()
	req.PolicyType = policy.POLICY_TYPE_RADIO
	set, err := impl.CreatePolicy(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(set)
}

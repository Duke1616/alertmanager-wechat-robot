package impl

import (
	"context"
	"github.com/Duke1616/alertmanager-wechat-robot/apps/alert"
	"github.com/infraboard/mcube/exception"
)

func (s *service) SaveAlertInformation(ctx context.Context, req *alert.SaveAlertRequest) (*alert.SaveAlertResponse, error) {
	ais := req.AlertInformationS()

	if _, err := s.col.InsertMany(ctx, ais); err != nil {
		return nil, exception.NewInternalServerError("inserted a rule document error, %s", err)
	}

	return &alert.SaveAlertResponse{Message: "ok"}, nil
}

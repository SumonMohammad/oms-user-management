package service

import (
	"context"
	pb "gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"

)

type HealthService interface {
  HealthCheck(ctx context.Context) (*pb.HealthCheckResponse, error)
}

type HealthReceiver struct {
	*OmsUserManagementService
}

func (ms *OmsUserManagementService) Health() HealthService {
	return &HealthReceiver{
		ms,
	}
}


func (v *HealthReceiver) HealthCheck(ctx context.Context) (*pb.HealthCheckResponse, error) {
    v.log.Info(ctx, "Service health check handler")
    err := v.db.Health().Ping(ctx)
    if err != nil {
      return nil, err
    }
    return &pb.HealthCheckResponse{
        Status: "ok",
    }, nil
}


package server

import (
	"context"
	"github.com/gogo/protobuf/types"
	"gitlab.techetronventures.com/core/backend/pkg/log"
	"gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/service"
	pb "gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
)

type OmsUserManagementServer struct {
	log     *log.Logger
	service service.Service
}

func New(logger *log.Logger, service service.Service) (srv *OmsUserManagementServer) {
	srv = new(OmsUserManagementServer)
	srv.service = service
	srv.log = logger.Named("server")
	return
}

// HealthHandler handles the health-check API call
func (v *OmsUserManagementServer) HealthCheck(ctx context.Context, empty *types.Empty) (*pb.HealthCheckResponse, error) {
	v.log.Info(ctx, "Health check handler triggered")
	return v.service.Health().HealthCheck(ctx)
}

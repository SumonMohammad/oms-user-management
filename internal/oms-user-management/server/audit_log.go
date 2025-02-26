package server

import (
	"context"
	"gitlab.techetronventures.com/core/grpctest/pkg/grpc"
)

func (s *GrpctestServer) CreateAuditLog(ctx context.Context, req *grpc.CreateAuditLogRequest) (*grpc.CreateAuditLogResponse, error) {
	res, err := s.service.AuditLog().CreateAuditLog(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) GetAuditLogs(ctx context.Context, req *grpc.GetAuditLogsRequest) (*grpc.GetAuditLogsResponse, error) {
	res, err := s.service.AuditLog().GetAuditLogs(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) GetAuditLogById(ctx context.Context, req *grpc.GetAuditLogByIdRequest) (*grpc.GetAuditLogByIdResponse, error) {
	res, err := s.service.AuditLog().GetAuditLogById(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

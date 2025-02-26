package server

import (
	"context"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
)

func (s *OmsUserManagementServer) CreateAuditLog(ctx context.Context, req *grpc.CreateAuditLogRequest) (*grpc.CreateAuditLogResponse, error) {
	res, err := s.service.AuditLog().CreateAuditLog(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetAuditLogs(ctx context.Context, req *grpc.GetAuditLogsRequest) (*grpc.GetAuditLogsResponse, error) {
	res, err := s.service.AuditLog().GetAuditLogs(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetAuditLogById(ctx context.Context, req *grpc.GetAuditLogByIdRequest) (*grpc.GetAuditLogByIdResponse, error) {
	res, err := s.service.AuditLog().GetAuditLogById(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

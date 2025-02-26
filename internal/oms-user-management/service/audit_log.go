package service

import (
	"context"
	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/status"

	//"fmt"
	model "gitlab.techetronventures.com/core/grpctest/internal/grpctest/models"
	//pg "gitlab.techetronventures.com/core/grpctest/internal/grpctest/models/pg"
	"gitlab.techetronventures.com/core/grpctest/pkg/grpc"
	"go.uber.org/zap"
)

type AuditLogs interface {
	CreateAuditLog(ctx context.Context, req *grpc.CreateAuditLogRequest) (*grpc.CreateAuditLogResponse, error)
	GetAuditLogs(ctx context.Context, req *grpc.GetAuditLogsRequest) (*grpc.GetAuditLogsResponse, error)
	GetAuditLogById(ctx context.Context, req *grpc.GetAuditLogByIdRequest) (*grpc.GetAuditLogByIdResponse, error)
}

type AuditLog struct {
	service *GrpctestService
}

type AuditLogReceiver struct {
	*GrpctestService
}

func (ms *GrpctestService) AuditLog() AuditLogs {
	return &AuditLogReceiver{
		ms,
	}
}

func (s *AuditLogReceiver) CreateAuditLog(ctx context.Context, req *grpc.CreateAuditLogRequest) (*grpc.CreateAuditLogResponse, error) {

	auditLog := &model.AuditLog{
		IPAddress: req.IpAddress,
	}

	// Call the database method to create the trader
	err := s.db.AuditLog().CreateAuditLog(ctx, auditLog)
	if err != nil {
		msg := "failed to create audit log"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.CreateAuditLogResponse{
		Code: 0,
	}, nil
}

func (s *AuditLogReceiver) GetAuditLogs(ctx context.Context, req *grpc.GetAuditLogsRequest) (*grpc.GetAuditLogsResponse, error) {

	res, count, err := s.db.AuditLog().GetAuditLogs(ctx, req)
	if err != nil {
		msg := "failed to get broker admins"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}
	auditLogs := []*grpc.GetAuditLogsResponseAuditLogList{}

	for _, item := range res {

		auditLog := &grpc.GetAuditLogsResponseAuditLogList{
			Id: item.ID,
		}
		auditLogs = append(auditLogs, auditLog)
	}
	return &grpc.GetAuditLogsResponse{
		AuditLogs: auditLogs,
		PaginationResponse: &grpc.PaginationInfoResponse{
			TotalRecordCount: int32(count),
		},
	}, nil
}

func (s *AuditLogReceiver) GetAuditLogById(ctx context.Context, req *grpc.GetAuditLogByIdRequest) (*grpc.GetAuditLogByIdResponse, error) {
	// Call the database layer to fetch the trader by ID or email
	auditLog, err := s.db.AuditLog().GetAuditLogById(ctx, req)
	if err != nil {
		msg := "failed to get audit log by id"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Map the result from the database layer to the GRPC response
	response := &grpc.GetAuditLogByIdResponse{
		Id: auditLog.Id,
	}

	return response, nil
}

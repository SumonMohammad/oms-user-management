package pg

import (
	"context"
	"github.com/uptrace/bun"
	"gitlab.techetronventures.com/core/backend/pkg/log"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	//"time"
)

func (db *DB) AuditLog() model.AuditLogs {
	return &AuditLog{
		IDB: db.DB,
		log: db.log,
	}
}

func (db *Tx) AuditLog() model.AuditLogs {
	return &AuditLog{
		IDB: db.Tx,
		log: db.log,
	}
}

type AuditLog struct {
	bun.IDB
	log *log.Logger
}

func (s *AuditLog) CreateAuditLog(ctx context.Context, auditLog *model.AuditLog) error {
	_, err := s.NewInsert().Model(auditLog).
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *AuditLog) GetAuditLogs(ctx context.Context, req *grpc.GetAuditLogsRequest) ([]*model.AuditLog, int64, error) {
	auditLogs := []*model.AuditLog{}
	offset := (req.PaginationRequest.PageToken - 1) * req.PaginationRequest.PageSize
	query := s.NewSelect().
		Model((*model.AuditLog)(nil)).
		Where("deleted_at IS NULL")
	count, err := query.Limit(int(req.PaginationRequest.PageSize)).
		Offset(int(offset)).
		Order("created_at DESC").
		ScanAndCount(ctx, &auditLogs)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, 0, err
	}

	return auditLogs, int64(count), nil
}

func (s *AuditLog) GetAuditLogById(ctx context.Context, req *grpc.GetAuditLogByIdRequest) (*grpc.GetAuditLogByIdResponse, error) {
	auditLog := &model.AuditLog{}
	query := s.NewSelect().
		Model(auditLog).
		Where("deleted_at IS NULL").
		Where("attempt_by_id = ?", auditLog.AttemptByID)
	err := query.Scan(ctx)
	if err != nil {
		msg := "Failed to fetch audit log"
		return nil, status.Error(codes.Internal, msg)
	}

	// Map the result to the response format
	response := &grpc.GetAuditLogByIdResponse{
		Id: auditLog.ID,
	}

	return response, nil
}

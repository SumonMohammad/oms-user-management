package pg

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
	"gitlab.techetronventures.com/core/backend/pkg/log"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"go.uber.org/zap"
	"time"
	//"google.golang.org/grpc/status"
	//"time"
)

func (db *DB) TWS() model.Tws {
	return &TWS{
		IDB: db.DB,
		log: db.log,
	}
}

func (db *Tx) TWS() model.Tws {
	return &TWS{
		IDB: db.Tx,
		log: db.log,
	}
}

type TWS struct {
	bun.IDB
	log *log.Logger
}

func (s *TWS) CreateTws(ctx context.Context, tws *model.TWS) error {
	_, err := s.NewInsert().Model(tws).Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *TWS) UpdateTws(ctx context.Context, tws *model.TWS) error {
	// Set the UpdatedAt field
	tws.UpdatedAt = time.Now()

	// Execute the update query
	result, err := s.NewUpdate().Model(tws).
		Where("id = ?", tws.ID).
		ExcludeColumn("created_at").
		//Column("updated_at").
		Exec(ctx)

	if err != nil {
		s.log.Error(ctx, "Failed to execute update query", zap.Error(err))
		return err
	}

	// Get the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		s.log.Error(ctx, "Failed to get rows affected", zap.Error(err))
		return err
	}

	// Check if no rows were affected (indicating no matching ID)
	if rowsAffected == 0 {
		msg := fmt.Sprintf("Tws with ID %d does not exist", tws.ID)
		s.log.Warn(ctx, msg)
		return fmt.Errorf(msg)
	}

	return nil
}

func (s *TWS) GetTws(ctx context.Context, req *grpc.GetTwsRequest) ([]*model.TWS, int64, error) {
	twsList := []*model.TWS{}
	offset := (req.PaginationRequest.PageToken - 1) * req.PaginationRequest.PageSize
	query := s.NewSelect().
		Model((*model.TWS)(nil)).
		Where("deleted_at IS NULL")
	count, err := query.Limit(int(req.PaginationRequest.PageSize)).
		Offset(int(offset)).
		Order("created_at DESC").
		ScanAndCount(ctx, &twsList)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, 0, err
	}

	return twsList, int64(count), nil
}

func (s *TWS) DeleteTws(ctx context.Context, tws *model.TWS) error {
	tws.DeletedAt = time.Now()
	_, err := s.NewUpdate().
		Model(tws).
		Column("deleted_at").
		Where("id = ?", tws.ID).
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

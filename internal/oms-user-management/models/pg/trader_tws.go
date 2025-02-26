package pg

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
	"gitlab.techetronventures.com/core/backend/pkg/log"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"time"
	//"google.golang.org/grpc/status"
	//"time"
)

func (db *DB) TraderTws() model.TradersTws {
	return &TraderTws{
		IDB: db.DB,
		log: db.log,
	}
}

func (db *Tx) TraderTws() model.TradersTws {
	return &TraderTws{
		IDB: db.Tx,
		log: db.log,
	}
}

type TraderTws struct {
	bun.IDB
	log *log.Logger
}

func (s *TraderTws) CreateTraderTws(ctx context.Context, traderTws *model.TraderTws) error {
	_, err := s.NewInsert().Model(traderTws).
		//Where("deleted_at IS NOT NULL").
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *TraderTws) GetTradersTws(ctx context.Context, req *grpc.GetTradersTwsMapRequest) ([]*model.TraderTws, int64, error) {
	tradersTws := []*model.TraderTws{}
	offset := (req.PaginationRequest.PageToken - 1) * req.PaginationRequest.PageSize
	query := s.NewSelect().
		Model((*model.TraderTws)(nil)).
		Where("deleted_at IS NULL")
	count, err := query.Limit(int(req.PaginationRequest.PageSize)).
		Offset(int(offset)).
		Order("created_at DESC").
		ScanAndCount(ctx, &tradersTws)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, 0, err
	}

	return tradersTws, int64(count), nil
}

func (s *TraderTws) DeleteTraderTws(ctx context.Context, traderTws *model.TraderTws) error {
	traderTws.DeletedAt = time.Now()
	_, err := s.NewUpdate().
		Model(traderTws).
		Column("deleted_at").
		Where("tws_id = ?", traderTws.TwsId).
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *TraderTws) UpdateTraderTws(ctx context.Context, traderTws *model.TraderTws) error {
	traderTws.UpdatedAt = time.Now()

	// Check if the record exists
	exists, err := s.NewSelect().Model(traderTws).
		Where("id = ?", traderTws.ID).
		Exists(ctx)
	if err != nil {
		return fmt.Errorf("failed to check existence: %w", err)
	}

	if !exists {
		msg := "No trader TWS record found with the given TWS ID"
		return fmt.Errorf(msg)
	}

	// Update the record
	_, err = s.NewUpdate().Model(traderTws).
		Where("id = ?", traderTws.ID).
		ExcludeColumn("created_at").
		Exec(ctx)
	return err
}

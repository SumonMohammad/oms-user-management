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

func (db *DB) TraderTeam() model.TraderTeams {
	return &TraderTeam{
		IDB: db.DB,
		log: db.log,
	}
}

func (db *Tx) TraderTeam() model.TraderTeams {
	return &TraderTeam{
		IDB: db.Tx,
		log: db.log,
	}
}

type TraderTeam struct {
	bun.IDB
	log *log.Logger
}

func (s *TraderTeam) CreateTeam(ctx context.Context, traderTeam *model.TraderTeam) error {
	_, err := s.NewInsert().Model(traderTeam).
		//Where("deleted_at IS NOT NULL").
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *TraderTeam) GetTeams(ctx context.Context, req *grpc.GetTraderTeamRequest) ([]*model.TraderTeam, int64, error) {
	tradersTeam := []*model.TraderTeam{}
	offset := (req.PaginationRequest.PageToken - 1) * req.PaginationRequest.PageSize
	query := s.NewSelect().
		Model((*model.TraderTeam)(nil)).
		Where("deleted_at IS NULL")
	count, err := query.Limit(int(req.PaginationRequest.PageSize)).
		Offset(int(offset)).
		Order("created_at DESC").
		ScanAndCount(ctx, &tradersTeam)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, 0, err
	}

	return tradersTeam, int64(count), nil
}

func (s *TraderTeam) DeleteTeam(ctx context.Context, traderTeam *model.TraderTeam) error {
	traderTeam.DeletedAt = time.Now()
	_, err := s.NewUpdate().
		Model(traderTeam).
		Column("deleted_at").
		Where("id = ?", traderTeam.ID).
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *TraderTeam) UpdateTeam(ctx context.Context, traderTeam *model.TraderTeam) error {
	traderTeam.UpdatedAt = time.Now()

	// Check if the record exists
	exists, err := s.NewSelect().Model(traderTeam).
		Where("id = ?", traderTeam.ID).
		Exists(ctx)
	if err != nil {
		return fmt.Errorf("failed to check existence: %w", err)
	}

	if !exists {
		msg := "No trader TWS record found with the given TWS ID"
		return fmt.Errorf(msg)
	}

	// Update the record
	_, err = s.NewUpdate().Model(traderTeam).
		Where("id = ?", traderTeam.ID).
		ExcludeColumn("created_at").
		Exec(ctx)
	return err
}

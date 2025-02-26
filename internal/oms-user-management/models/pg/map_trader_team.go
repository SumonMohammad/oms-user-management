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

func (db *DB) MapTraderTeam() model.MapTraderTeams {
	return &MapTraderTeam{
		IDB: db.DB,
		log: db.log,
	}
}

func (db *Tx) MapTraderTeam() model.MapTraderTeams {
	return &MapTraderTeam{
		IDB: db.Tx,
		log: db.log,
	}
}

type MapTraderTeam struct {
	bun.IDB
	log *log.Logger
}

func (s *MapTraderTeam) CreateMapTeam(ctx context.Context, mapTeam *model.MapTraderTeam) error {
	_, err := s.NewInsert().Model(mapTeam).
		//Where("deleted_at IS NOT NULL").
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *MapTraderTeam) GetMappedTeams(ctx context.Context, req *grpc.GetTraderMapRequest) ([]*model.MapTraderTeam, int64, error) {
	mappedTeams := []*model.MapTraderTeam{}
	offset := (req.PaginationRequest.PageToken - 1) * req.PaginationRequest.PageSize
	query := s.NewSelect().
		Model((*model.MapTraderTeam)(nil)).
		Where("deleted_at IS NULL")
	count, err := query.Limit(int(req.PaginationRequest.PageSize)).
		Offset(int(offset)).
		Order("created_at DESC").
		ScanAndCount(ctx, &mappedTeams)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, 0, err
	}

	return mappedTeams, int64(count), nil
}

func (s *MapTraderTeam) DeleteMappedTeam(ctx context.Context, mapTeam *model.MapTraderTeam) error {
	mapTeam.DeletedAt = time.Now()
	_, err := s.NewUpdate().
		Model(mapTeam).
		Column("deleted_at").
		Where("id = ?", mapTeam.ID).
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *MapTraderTeam) UpdateMapTeam(ctx context.Context, mapTeam *model.MapTraderTeam) error {
	mapTeam.UpdatedAt = time.Now()

	// Check if the record exists
	exists, err := s.NewSelect().Model(mapTeam).
		Where("id = ?", mapTeam.ID).
		Exists(ctx)
	if err != nil {
		return fmt.Errorf("failed to check existence: %w", err)
	}

	if !exists {
		msg := "No trader TWS record found with the given ID"
		return fmt.Errorf(msg)
	}

	// Update the record
	_, err = s.NewUpdate().Model(mapTeam).
		Where("id = ?", mapTeam.ID).
		ExcludeColumn("created_at").
		Exec(ctx)
	return err
}

package pg

import (
	"context"
	"github.com/uptrace/bun"
	"gitlab.techetronventures.com/core/backend/pkg/log"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
	//"time"
)

func (db *DB) MapUserRole() model.MapUserRoles {
	return &MapUserRole{
		IDB: db.DB,
		log: db.log,
	}
}

func (db *Tx) MapUserRole() model.MapUserRoles {
	return &MapUserRole{
		IDB: db.Tx,
		log: db.log,
	}
}

type MapUserRole struct {
	bun.IDB
	log *log.Logger
}

func (s *MapUserRole) CreateMapUserRole(ctx context.Context, mapUserRole *model.MapUserRole) error {
	_, err := s.NewInsert().Model(mapUserRole).
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *MapUserRole) UpdateMapUserRole(ctx context.Context, mapUserRole *model.MapUserRole) error {
	mapUserRole.UpdatedAt = time.Now()
	_, err := s.NewUpdate().Model(mapUserRole).
		Where("id = ?", mapUserRole.ID).
		ExcludeColumn("created_at").
		Column("updated_at").
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *MapUserRole) GetMapUserRoles(ctx context.Context, req *grpc.GetMapUserRolesRequest) ([]*model.MapUserRole, int64, error) {
	mapUserRoles := []*model.MapUserRole{}
	offset := (req.PaginationRequest.PageToken - 1) * req.PaginationRequest.PageSize
	query := s.NewSelect().
		Model((*model.MapUserRole)(nil)).
		Where("deleted_at IS NULL")
	count, err := query.Limit(int(req.PaginationRequest.PageSize)).
		Offset(int(offset)).
		Order("created_at DESC").
		ScanAndCount(ctx, &mapUserRoles)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, 0, err
	}

	return mapUserRoles, int64(count), nil
}

func (s *MapUserRole) GetMapUserRoleById(ctx context.Context, req *grpc.GetMapUserRoleByIdRequest) (*model.MapUserRole, error) {
	mapUserRole := &model.MapUserRole{}
	query := s.NewSelect().Model(mapUserRole).Where("deleted_at IS NULL").
		Where("id = ?", req.Id)

	// Execute the query
	err := query.Scan(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to fetch mapUserRole")
	}

	return mapUserRole, nil
}

func (s *MapUserRole) DeleteMapUserRole(ctx context.Context, mapUserRole *model.MapUserRole) error {
	mapUserRole.DeletedAt = time.Now()
	_, err := s.NewUpdate().
		Model(mapUserRole).
		Column("deleted_at").
		Where("id = ?", mapUserRole.ID).
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

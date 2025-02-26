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

func (db *DB) Permission() model.Permissions {
	return &Permission{
		IDB: db.DB,
		log: db.log,
	}
}

func (db *Tx) Permission() model.Permissions {
	return &Permission{
		IDB: db.Tx,
		log: db.log,
	}
}

type Permission struct {
	bun.IDB
	log *log.Logger
}

func (s *Permission) CreatePermission(ctx context.Context, Permission *model.Permission) error {
	_, err := s.NewInsert().Model(Permission).
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *Permission) UpdatePermission(ctx context.Context, permission *model.Permission) error {
	permission.UpdatedAt = time.Now()
	_, err := s.NewUpdate().Model(permission).
		Where("id = ?", permission.ID).
		ExcludeColumn("created_at").
		Column("updated_at").
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *Permission) GetPermissions(ctx context.Context, req *grpc.GetPermissionsRequest) ([]*model.Permission, int64, error) {
	permissions := []*model.Permission{}
	offset := (req.PaginationRequest.PageToken - 1) * req.PaginationRequest.PageSize
	query := s.NewSelect().
		Model((*model.Permission)(nil)).
		Where("deleted_at IS NULL")
	count, err := query.Limit(int(req.PaginationRequest.PageSize)).
		Offset(int(offset)).
		Order("created_at DESC").
		ScanAndCount(ctx, &permissions)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, 0, err
	}

	return permissions, int64(count), nil
}

func (s *Permission) GetPermissionById(ctx context.Context, req *grpc.GetPermissionByIdRequest) (*model.Permission, error) {
	permission := &model.Permission{}
	query := s.NewSelect().Model(permission).Where("deleted_at IS NULL").
		Where("id = ?", req.Id)

	// Execute the query
	err := query.Scan(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to fetch Permission")
	}

	return permission, nil
}

func (s *Permission) DeletePermission(ctx context.Context, permission *model.Permission) error {
	permission.DeletedAt = time.Now()
	_, err := s.NewUpdate().
		Model(permission).
		Column("deleted_at").
		Where("id = ?", permission.ID).
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

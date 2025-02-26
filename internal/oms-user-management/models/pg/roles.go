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

func (db *DB) Role() model.Roles {
	return &Role{
		IDB: db.DB,
		log: db.log,
	}
}

func (db *Tx) Role() model.Roles {
	return &Role{
		IDB: db.Tx,
		log: db.log,
	}
}

type Role struct {
	bun.IDB
	log *log.Logger
}

func (s *Role) CreateRole(ctx context.Context, role *model.Role) error {
	_, err := s.NewInsert().Model(role).
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *Role) UpdateRole(ctx context.Context, role *model.Role) error {
	role.UpdatedAt = time.Now()
	_, err := s.NewUpdate().Model(role).
		Where("id = ?", role.ID).
		ExcludeColumn("created_at").
		Column("updated_at").
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *Role) GetRoles(ctx context.Context, req *grpc.GetRolesRequest) ([]*model.Role, int64, error) {
	roles := []*model.Role{}
	offset := (req.PaginationRequest.PageToken - 1) * req.PaginationRequest.PageSize
	query := s.NewSelect().
		Model((*model.Role)(nil)).
		Where("deleted_at IS NULL")
	count, err := query.Limit(int(req.PaginationRequest.PageSize)).
		Offset(int(offset)).
		Order("created_at DESC").
		ScanAndCount(ctx, &roles)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, 0, err
	}

	return roles, int64(count), nil
}

func (s *Role) GetRoleById(ctx context.Context, req *grpc.GetRoleByIdRequest) (*model.Role, error) {
	role := &model.Role{}
	query := s.NewSelect().Model(role).Where("deleted_at IS NULL").
		Where("id = ?", req.Id)

	// Execute the query
	err := query.Scan(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to fetch Role")
	}

	return role, nil
}

func (s *Role) DeleteRole(ctx context.Context, role *model.Role) error {
	role.DeletedAt = time.Now()
	_, err := s.NewUpdate().
		Model(role).
		Column("deleted_at").
		Where("id = ?", role.ID).
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

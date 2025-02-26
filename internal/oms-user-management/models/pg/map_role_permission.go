package pg

import (
	"context"
	"github.com/uptrace/bun"
	"gitlab.techetronventures.com/core/backend/pkg/log"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"time"
	//"time"
)

func (db *DB) MapRolePermission() model.MapRolePermissions {
	return &MapRolePermission{
		IDB: db.DB,
		log: db.log,
	}
}

func (db *Tx) MapRolePermission() model.MapRolePermissions {
	return &MapRolePermission{
		IDB: db.Tx,
		log: db.log,
	}
}

type MapRolePermission struct {
	bun.IDB
	log *log.Logger
}

func (s *MapRolePermission) CreateMapRolePermission(ctx context.Context, rolePermission *model.MapRolePermission) error {
	_, err := s.NewInsert().Model(rolePermission).
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *MapRolePermission) UpdateMapRolePermission(ctx context.Context, rolePermission *model.MapRolePermission) error {
	rolePermission.UpdatedAt = time.Now()
	_, err := s.NewUpdate().Model(rolePermission).
		Where("id = ?", rolePermission.ID).
		ExcludeColumn("created_at").
		Column("updated_at").
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *MapRolePermission) GetMapRolePermissions(ctx context.Context, req *grpc.GetMapRolePermissionsRequest) ([]*model.MapRolePermission, int64, error) {
	rolePermissions := []*model.MapRolePermission{}
	offset := (req.PaginationRequest.PageToken - 1) * req.PaginationRequest.PageSize
	query := s.NewSelect().
		Model((*model.MapRolePermission)(nil)).
		Where("deleted_at IS NULL")
	count, err := query.Limit(int(req.PaginationRequest.PageSize)).
		Offset(int(offset)).
		Order("created_at DESC").
		ScanAndCount(ctx, &rolePermissions)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, 0, err
	}

	return rolePermissions, int64(count), nil
}

func (s *MapRolePermission) GetMapRolePermissionById(ctx context.Context, req *grpc.GetMapRolePermissionByIdRequest) ([]*model.GetPermissionsByRole, error) {
	var results []*model.GetPermissionsByRole
	err := s.NewSelect().
		TableExpr("map_role_permission AS mrp").
		Join("JOIN permissions AS p ON mrp.permission_id = p.id").
		ColumnExpr("mrp.role_id, p.id AS permission_id, p.name, mrp.is_enabled").
		Where("mrp.role_id = ?", req.RoleId).
		Scan(ctx, &results)

	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}
	return results, nil
}

func (s *MapRolePermission) DeleteMapRolePermission(ctx context.Context, rolePermission *model.MapRolePermission) error {
	rolePermission.DeletedAt = time.Now()
	_, err := s.NewUpdate().
		Model(rolePermission).
		Column("deleted_at").
		Where("id = ?", rolePermission.ID).
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

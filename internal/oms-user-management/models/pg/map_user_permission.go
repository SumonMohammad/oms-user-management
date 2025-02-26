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

func (db *DB) MapUserPermission() model.MapUserPermissions {
	return &MapUserPermission{
		IDB: db.DB,
		log: db.log,
	}
}

func (db *Tx) MapUserPermission() model.MapUserPermissions {
	return &MapUserPermission{
		IDB: db.Tx,
		log: db.log,
	}
}

type MapUserPermission struct {
	bun.IDB
	log *log.Logger
}

func (s *MapUserPermission) CreateMapUserPermission(ctx context.Context, mapUserPermission *model.MapUserPermission) error {
	_, err := s.NewInsert().Model(mapUserPermission).
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *MapUserPermission) UpdateMapUserPermission(ctx context.Context, mapUserPermission *model.MapUserPermission) error {
	mapUserPermission.UpdatedAt = time.Now()
	_, err := s.NewUpdate().Model(mapUserPermission).
		Where("user_id = ? AND permission_id = ?", mapUserPermission.UserID, mapUserPermission.PermissionID).
		ExcludeColumn("created_at").
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

//func (s *MapUserPermission) GetUserPermissionsByUserId(ctx context.Context, req *grpc.GetUserPermissionsByUserIdRequest) ([]*model.UserPermissions, error) {
//	var results []*model.UserPermissions
//
//	// Subquery: Permissions from roles
//	rolePermissions := s.NewSelect().
//		TableExpr("permissions AS p").
//		Join("JOIN map_role_permission AS rp ON p.id = rp.permission_id").
//		Join("JOIN map_user_role AS ur ON rp.role_id = ur.role_id").
//		ColumnExpr("p.name, p.id AS permission_id").
//		Where("ur.user_id = ?", req.UserId)
//
//	// Subquery: Direct user permissions
//	userPermissions := s.NewSelect().
//		TableExpr("permissions AS p").
//		Join("JOIN map_user_permission AS up ON p.id = up.permission_id").
//		ColumnExpr("p.name, p.id AS permission_id").
//		Where("up.user_id = ?", req.UserId)
//
//	// Combine with UNION ALL
//	unionQuery := s.NewSelect().
//		ColumnExpr("name, permission_id").
//		TableExpr("(?) AS role_perms", rolePermissions).
//		UnionAll(
//			s.NewSelect().
//				ColumnExpr("name, permission_id").
//				TableExpr("(?) AS user_perms", userPermissions),
//		)
//
//	// Wrap the union in a subquery and provide an alias
//	finalQuery := s.NewSelect().
//		ColumnExpr("name, permission_id").
//		TableExpr("(?) AS permissions_union", unionQuery)
//
//	// Execute the final query
//	err := finalQuery.Scan(ctx, &results)
//	if err != nil {
//		return nil, err
//	}
//
//	return results, nil
//}

func (s *MapUserPermission) DeleteMapUserPermission(ctx context.Context, mapUserPermission *model.MapUserPermission) error {
	mapUserPermission.DeletedAt = time.Now()
	_, err := s.NewUpdate().
		Model(mapUserPermission).
		Column("deleted_at").
		Where("user_id = ? AND permission_id = ?", mapUserPermission.UserID, mapUserPermission.PermissionID).
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *MapUserPermission) GetUserPermissionsByUserId(ctx context.Context, req *grpc.GetUserPermissionsByUserIdRequest) ([]*model.UserPermissions, error) {
	var results []*model.UserPermissions

	// Subquery: Permissions from roles
	rolePermissions := s.NewSelect().
		TableExpr("permissions AS p").
		Join("JOIN map_role_permission AS rp ON p.id = rp.permission_id").
		Join("JOIN map_user_role AS ur ON rp.role_id = ur.role_id").
		ColumnExpr("p.name, p.id AS permission_id").
		Where("ur.user_id = ?", req.UserId)

	// Subquery: Direct user permissions
	userPermissions := s.NewSelect().
		TableExpr("permissions AS p").
		Join("JOIN map_user_permission AS up ON p.id = up.permission_id").
		ColumnExpr("p.name, p.id AS permission_id").
		Where("up.deleted_at = NULL").
		Where("up.user_id = ?", req.UserId)

	// Combine with UNION ALL
	unionQuery := s.NewSelect().
		ColumnExpr("name, permission_id").
		TableExpr("(?) AS combined_permissions", rolePermissions).
		UnionAll(
			s.NewSelect().
				ColumnExpr("name, permission_id").
				TableExpr("(?) AS user_perms", userPermissions),
		)

	// Final query: Exclude revoked permissions
	finalQuery := s.NewSelect().
		ColumnExpr("name, permission_id").
		TableExpr("(?) AS permissions_union", unionQuery).
		Where("permission_id NOT IN (SELECT permission_id FROM map_user_permission WHERE user_id = ? AND is_revoked = TRUE )", req.UserId)

	// Execute the final query
	err := finalQuery.Scan(ctx, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

package model

import (
	"context"
	"github.com/uptrace/bun"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"time"
	//"time"
)

type MapRolePermissions interface {
	CreateMapRolePermission(ctx context.Context, rolePermission *MapRolePermission) error
	UpdateMapRolePermission(ctx context.Context, rolePermission *MapRolePermission) error
	GetMapRolePermissions(ctx context.Context, req *grpc.GetMapRolePermissionsRequest) ([]*MapRolePermission, int64, error)
	GetMapRolePermissionById(ctx context.Context, req *grpc.GetMapRolePermissionByIdRequest) ([]*GetPermissionsByRole, error)
	DeleteMapRolePermission(ctx context.Context, rolePermission *MapRolePermission) error
}

type MapRolePermission struct {
	bun.BaseModel `bun:"table:map_role_permission"`

	ID           int64     `json:"id" bun:"id,pk,autoincrement"`
	PermissionID int64     `json:"permission_id" bun:"permission_id"`
	RoleID       int64     `json:"role_id" bun:"role_id"`
	IsEnabled    bool      `json:"is_enabled" bun:"is_enabled"`
	CreatedAt    time.Time `json:"created_at" bun:"created_at,default:current_timestamp"`
	UpdatedAt    time.Time `json:"updated_at" bun:"updated_at,nullzero"`
	DeletedAt    time.Time `json:"-" bun:"deleted_at,nullzero"`
}

type GetPermissionsByRole struct {
	PermissionID int64  `json:"permission_id" bun:"permission_id"`
	Name         string `json:"name" bun:"name"`
	RoleID       int64  `json:"role_id" bun:"role_id"`
	IsEnabled    bool   `json:"is_enabled" bun:"is_enabled"`
}

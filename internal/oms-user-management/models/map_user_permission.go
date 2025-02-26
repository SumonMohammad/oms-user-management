package model

import (
	"context"
	"github.com/uptrace/bun"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"time"
	//"time"
)

type MapUserPermissions interface {
	CreateMapUserPermission(ctx context.Context, mapUserPermission *MapUserPermission) error
	UpdateMapUserPermission(ctx context.Context, mapUserPermission *MapUserPermission) error
	GetUserPermissionsByUserId(ctx context.Context, req *grpc.GetUserPermissionsByUserIdRequest) ([]*UserPermissions, error)
	DeleteMapUserPermission(ctx context.Context, mapUserPermission *MapUserPermission) error
}

type MapUserPermission struct {
	bun.BaseModel `bun:"table:map_user_permission"`

	ID           int64     `json:"id" bun:"id,pk,autoincrement"`
	PermissionID int64     `json:"permission_id" bun:"permission_id"`
	UserID       int64     `json:"user_id" bun:"user_id"`
	IsEnabled    bool      `json:"is_enabled" bun:"is_enabled"`
	IsRevoked    bool      `json:"is_revoked" bun:"is_revoked,default:false"`
	CreatedAt    time.Time `json:"created_at" bun:"created_at,default:current_timestamp"`
	UpdatedAt    time.Time `json:"updated_at" bun:"updated_at,nullzero"`
	DeletedAt    time.Time `json:"-" bun:"deleted_at,nullzero"`
}

type UserPermissions struct {
	PermissionID int64  `json:"permission_id" bun:"permission_id"`
	Name         string `json:"name" bun:"name"`
}

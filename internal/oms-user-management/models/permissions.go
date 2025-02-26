package model

import (
	"context"
	"github.com/uptrace/bun"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"time"
	//"time"
)

type Permissions interface {
	CreatePermission(ctx context.Context, permission *Permission) error
	UpdatePermission(ctx context.Context, permission *Permission) error
	GetPermissions(ctx context.Context, req *grpc.GetPermissionsRequest) ([]*Permission, int64, error)
	GetPermissionById(ctx context.Context, req *grpc.GetPermissionByIdRequest) (*Permission, error)
	DeletePermission(ctx context.Context, permission *Permission) error
}

type Permission struct {
	bun.BaseModel `bun:"table:permissions"`

	ID          int64     `json:"id" bun:"id,pk,autoincrement"`
	Name        string    `json:"name" bun:"name"`
	Description string    `json:"description" bun:"description"`
	IsEnabled   bool      `json:"is_enabled" bun:"is_enabled"`
	CreatedAt   time.Time `json:"created_at" bun:"created_at,default:current_timestamp"`
	UpdatedAt   time.Time `json:"updated_at" bun:"updated_at,nullzero"`
	DeletedAt   time.Time `json:"-" bun:"deleted_at,nullzero"`
}

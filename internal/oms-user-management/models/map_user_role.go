package model

import (
	"context"
	"github.com/uptrace/bun"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"time"
	//"time"
)

type MapUserRoles interface {
	CreateMapUserRole(ctx context.Context, mapUserRole *MapUserRole) error
	UpdateMapUserRole(ctx context.Context, mapUserRole *MapUserRole) error
	GetMapUserRoles(ctx context.Context, req *grpc.GetMapUserRolesRequest) ([]*MapUserRole, int64, error)
	GetMapUserRoleById(ctx context.Context, req *grpc.GetMapUserRoleByIdRequest) (*MapUserRole, error)
	DeleteMapUserRole(ctx context.Context, mapUserRole *MapUserRole) error
}

type MapUserRole struct {
	bun.BaseModel `bun:"table:map_user_role"`

	ID        int64     `json:"id" bun:"id,pk,autoincrement"`
	UserId    int64     `json:"user_id" bun:"user_id"`
	RoleId    int64     `json:"role_id" bun:"role_id"`
	IsEnabled bool      `json:"is_enabled" bun:"is_enabled"`
	CreatedAt time.Time `json:"created_at" bun:"created_at,default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" bun:"updated_at,nullzero"`
	DeletedAt time.Time `json:"-" bun:"deleted_at,nullzero"`
}

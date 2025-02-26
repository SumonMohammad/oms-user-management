package model

import (
	"context"
	"github.com/uptrace/bun"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"time"
	//"time"
)

type Roles interface {
	CreateRole(ctx context.Context, role *Role) error
	UpdateRole(ctx context.Context, role *Role) error
	GetRoles(ctx context.Context, req *grpc.GetRolesRequest) ([]*Role, int64, error)
	GetRoleById(ctx context.Context, req *grpc.GetRoleByIdRequest) (*Role, error)
	DeleteRole(ctx context.Context, role *Role) error
}

type Role struct {
	bun.BaseModel `bun:"table:roles"`

	ID          int64     `json:"id" bun:"id,pk,autoincrement"`
	RoleName    string    `json:"role_name" bun:"role_name"`
	Description string    `json:"description" bun:"description"`
	IsEnabled   bool      `json:"is_enabled" bun:"is_enabled"`
	CreatedAt   time.Time `json:"created_at" bun:"created_at,default:current_timestamp"`
	UpdatedAt   time.Time `json:"updated_at" bun:"updated_at,nullzero"`
	DeletedAt   time.Time `json:"-" bun:"deleted_at,nullzero"`
}

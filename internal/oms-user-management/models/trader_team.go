package model

import (
	"context"
	"github.com/uptrace/bun"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"time"
	//"time"
)

type TraderTeams interface {
	CreateTeam(ctx context.Context, traderTeam *TraderTeam) error
	UpdateTeam(ctx context.Context, traderTeam *TraderTeam) error
	GetTeams(ctx context.Context, req *grpc.GetTraderTeamRequest) ([]*TraderTeam, int64, error)
	DeleteTeam(ctx context.Context, traderTeam *TraderTeam) error
}

type TraderTeam struct {
	bun.BaseModel `bun:"table:trader_team"`

	ID          int64     `json:"id" bun:"id,pk,autoincrement"`
	Name        string    `json:"name" bun:"name"`
	Description string    `json:"description" bun:"description"`
	Status      string    `json:"status" bun:"status"`
	IsEnabled   bool      `json:"is_enabled" bun:"is_enabled"`
	IsDeleted   bool      `json:"is_deleted" bun:"is_deleted"`
	CreatedAt   time.Time `json:"created_at" bun:"created_at,default:current_timestamp"`
	UpdatedAt   time.Time `json:"updated_at" bun:"updated_at,nullzero"`
	DeletedAt   time.Time `json:"-" bun:"deleted_at,nullzero"`
}

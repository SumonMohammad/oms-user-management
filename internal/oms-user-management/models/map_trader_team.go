package model

import (
	"context"
	"github.com/uptrace/bun"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"time"
	//"time"
)

type MapTraderTeams interface {
	CreateMapTeam(ctx context.Context, mapTeam *MapTraderTeam) error
	UpdateMapTeam(ctx context.Context, mapTeam *MapTraderTeam) error
	GetMappedTeams(ctx context.Context, req *grpc.GetTraderMapRequest) ([]*MapTraderTeam, int64, error)
	DeleteMappedTeam(ctx context.Context, mapTeam *MapTraderTeam) error
}

type MapTraderTeam struct {
	bun.BaseModel `bun:"table:map_trader_team"`

	ID        int64     `json:"id" bun:"id,pk,autoincrement"`
	TeamId    int64     `json:"team-id" bun:"team_id"`
	TraderId  int64     `json:"trader-id" bun:"trader_id"`
	Status    string    `json:"status" bun:"status"`
	IsEnabled bool      `json:"is_enabled" bun:"is_enabled"`
	CreatedAt time.Time `json:"created_at" bun:"created_at,default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" bun:"updated_at,nullzero"`
	DeletedAt time.Time `json:"-" bun:"deleted_at,nullzero"`
}

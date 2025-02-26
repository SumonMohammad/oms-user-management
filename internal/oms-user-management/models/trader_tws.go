package model

import (
	"context"
	"github.com/uptrace/bun"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"time"
)

type TradersTws interface {
	CreateTraderTws(ctx context.Context, traderTws *TraderTws) error
	GetTradersTws(ctx context.Context, req *grpc.GetTradersTwsMapRequest) ([]*TraderTws, int64, error)
	UpdateTraderTws(ctx context.Context, traderTws *TraderTws) error
	DeleteTraderTws(ctx context.Context, traderTws *TraderTws) error
}

type TraderTws struct {
	bun.BaseModel `bun:"table:map_trader_tws"`

	ID        int64     `json:"id" bun:"id,pk,autoincrement"`
	TwsId     int64     `json:"tws_id" bun:"tws_id"`
	TraderId  int64     `json:"trader_id" bun:"trader_id"`
	IsEnabled bool      `json:"is_enabled" bun:"is_enabled"`
	Status    string    `json:"status" bun:"status"`
	IsDeleted bool      `json:"is_deleted" bun:"is_deleted"`
	ExpireAt  time.Time `json:"expire_at" bun:"expire_at,nullzero"`
	CreatedAt time.Time `json:"created_at" bun:"created_at,default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" bun:"updated_at,nullzero"`
	DeletedAt time.Time `json:"-" bun:"deleted_at,nullzero"`
}

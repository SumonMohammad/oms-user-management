package model

import (
	"context"
	"github.com/uptrace/bun"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"time"
	//"time"
)

type Tws interface {
	CreateTws(ctx context.Context, tws *TWS) error
	UpdateTws(ctx context.Context, tws *TWS) error
	GetTws(ctx context.Context, req *grpc.GetTwsRequest) ([]*TWS, int64, error)
	DeleteTws(ctx context.Context, tws *TWS) error
}

// trader_work_station
type TWS struct {
	bun.BaseModel `bun:"table:tws"`

	ID        int64     `json:"id" bun:"id,pk,autoincrement"`
	TwsCode   string    `json:"tws_code" bun:"tws_code"`
	IsEnabled bool      `json:"is_enabled" bun:"is_enabled"`
	Status    string    `json:"status" bun:"status"`
	IsActive  bool      `json:"is_active" bun:"is_active"`
	IsDeleted bool      `json:"is_deleted" bun:"is_deleted"`
	ExpireAt  time.Time `json:"-" bun:"expire_at,nullzero"`
	CreatedAt time.Time `json:"created_at" bun:"created_at,default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" bun:"updated_at,nullzero"`
	DeletedAt time.Time `json:"-" bun:"deleted_at,nullzero"`
}

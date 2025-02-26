package model

import (
	"context"
	"github.com/uptrace/bun"
	"time"
	//"time"
)

type Traders interface {
	UpdateTrader(ctx context.Context, trader *Trader) (int64, error)
	CreateTrader(ctx context.Context, trader *Trader) error
	GetTraderById(ctx context.Context, userId int64) (*UserWithTraderType, error)
	//GetInvestors(ctx context.Context, page, limit int32) ([]*OmsUserWithInvestorModel, int64, error)
}

type Trader struct {
	bun.BaseModel `bun:"table:traders"`

	ID            int64     `json:"id" bun:"id,pk,autoincrement"`
	UserId        int64     `json:"user_id" bun:"user_id"`
	Status        string    `json:"status" bun:"status"`
	BranchId      int64     `json:"branch_id" bun:"branch_id"`
	IsActive      bool      `json:"is_active" bun:"is_active"`
	ReadOnly      bool      `json:"read_only" bun:"read_only"`
	LicenceNumber string    `json:"licence_number" bun:"licence_number"`
	CanTrade      bool      `json:"can_trade" bun:"can_trade"`
	IsDeleted     bool      `json:"is_deleted" bun:"is_deleted"`
	ExpireAt      time.Time `json:"-" bun:"expire_at,nullzero"`
	CreatedAt     time.Time `json:"created_at" bun:"created_at,default:current_timestamp"`
	UpdatedAt     time.Time `json:"updated_at" bun:"updated_at,nullzero"`
	DeletedAt     time.Time `json:"-" bun:"deleted_at,nullzero"`
}

type UserWithTraderType struct {
	TraderUserId  int64  `json:"trader_user_id" bun:"trader_user_id"`
	UserName      string `json:"user_name" bun:"user_name"`
	EmailAddress  string `json:"email_address" bun:"email_address"`
	PhoneNumber   string `json:"phone_number" bun:"phone_number"`
	CountryCode   string `json:"country_code" bun:"country_code"`
	ReadOnly      bool   `json:"read_only" bun:"read_only"`
	BranchId      int64  `json:"branch_id" bun:"branch_id"`
	IsActive      bool   `json:"is_active" bun:"is_active"`
	CanTrade      bool   `json:"can_trade" bun:"can_trade"`
	LicenceNumber string `json:"licence_number" bun:"licence_number"`
}

package model

import (
	"context"
	"github.com/uptrace/bun"
	"time"
	//"time"
)

type Investors interface {
	UpdateInvestor(ctx context.Context, investor *Investor) (int64, error)
	CreateInvestor(ctx context.Context, investor *Investor) error
	GetInvestorById(ctx context.Context, userId int64) (*UserWithInvestorType, error)
	GetInvestors(ctx context.Context, page, limit int32) ([]*UserWithInvestorType, int64, error)
}

type Investor struct {
	bun.BaseModel `bun:"table:investors"`

	ID              int64     `json:"id" bun:"id,pk,autoincrement"`
	UserId          int64     `json:"user_id" bun:"user_id"`
	PrimaryTwsId    int64     `json:"primary_tws_id" bun:"primary_tws_id"`
	SecondaryTwsId  int64     `json:"secondary_tws_id" bun:"secondary_tws_id"`
	BoAccountNumber string    `json:"bo_account_number" bun:"bo_account_number"`
	Status          string    `json:"status" bun:"status"`
	CanTrade        bool      `json:"can_trade" bun:"can_trade"`
	ReadOnly        bool      `json:"read_only" bun:"read_only"`
	IsDeleted       bool      `json:"is_deleted" bun:"is_deleted"`
	ClientCode      string    `json:"client_code" bun:"client_code"`
	ExpireAt        time.Time `json:"-" bun:"expire_at,nullzero"`
	CreatedAt       time.Time `json:"created_at" bun:"created_at,default:current_timestamp"`
	UpdatedAt       time.Time `json:"updated_at" bun:"updated_at,nullzero"`
	DeletedAt       time.Time `json:"-" bun:"deleted_at,nullzero"`
}

type UserWithInvestorType struct {
	InvestorUserId  int64  `json:"investor_user_id" bun:"investor_user_id"`
	UserName        string `json:"user_name" bun:"user_name"`
	EmailAddress    string `json:"email_address" bun:"email_address"`
	PhoneNumber     string `json:"phone_number" bun:"phone_number"`
	CountryCode     string `json:"country_code" bun:"country_code"`
	PrimaryTwsId    int64  `json:"primary_tws_id" bun:"primary_tws_id"`
	ClientCode      string `json:"client_code" bun:"client_code"`
	BoAccountNumber string `json:"bo_account_number" bun:"bo_account_number"`
}

package model

import (
	"context"
	"github.com/uptrace/bun"
	"time"
	//"time"
)

type BrokerAdmins interface {
	UpdateBrokerAdmin(ctx context.Context, brokerAdmin *BrokerAdmin) (int64, error)
	CreateBrokerAdmin(ctx context.Context, brokerAdmin *BrokerAdmin) error
	GetBrokerAdminById(ctx context.Context, userId int64) (*UserWithBrokerAdminType, error)
	//GetBrokerAdmins(ctx context.Context, page, limit int32) ([]*UserWithBrokerAdminType, int64, error)
}

type BrokerAdmin struct {
	bun.BaseModel `bun:"table:broker_admin"`

	ID             int64     `json:"id" bun:"id,pk,autoincrement"`
	UserId         int64     `json:"user_id" bun:"user_id"`
	BranchId       int64     `json:"branch_id" bun:"branch_id"`
	CanTrade       bool      `json:"can_trade" bun:"can_trade"`
	ReadOnly       bool      `json:"read_only" bun:"read_only"`
	IsIsolatedUser bool      `json:"is_isolated_user" bun:"is_isolated_user"`
	Status         string    `json:"status" bun:"status"`
	IsDeleted      bool      `json:"is_deleted" bun:"is_deleted"`
	ExpireAt       time.Time `json:"-" bun:"expire_at,nullzero"`
	CreatedAt      time.Time `json:"created_at" bun:"created_at,default:current_timestamp"`
	UpdatedAt      time.Time `json:"updated_at" bun:"updated_at,nullzero"`
	DeletedAt      time.Time `json:"-" bun:"deleted_at,nullzero"`
}

type UserWithBrokerAdminType struct {
	BrokerUserId   int64  `json:"broker_user_id" bun:"broker_user_id"`
	UserName       string `json:"user_name" bun:"user_name"`
	EmailAddress   string `json:"email_address" bun:"email_address"`
	PhoneNumber    string `json:"phone_number" bun:"phone_number"`
	CountryCode    string `json:"country_code" bun:"country_code"`
	BranchId       int64  `json:"branch_id" bun:"branch_id"`
	CanTrade       bool   `json:"can_trade" bun:"can_trade"`
	ReadOnly       bool   `json:"read_only" bun:"read_only"`
	IsIsolatedUser bool   `json:"is_isolated_user" bun:"is_isolated_user"`
	IsDeleted      bool   `json:"is_deleted" bun:"is_deleted"`
}

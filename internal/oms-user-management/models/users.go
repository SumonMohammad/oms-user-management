package model

import (
	"context"
	"github.com/uptrace/bun"
	"time"
	//"time"
)

type Users interface {
	CreateInvestor(ctx context.Context, user *User) (int64, error)
	UpdateInvestor(ctx context.Context, user *User) error
	CreateTrader(ctx context.Context, user *User) (int64, error)
	UpdateTrader(ctx context.Context, user *User) error
	CreateBrokerAdmin(ctx context.Context, user *User) (int64, error)
	UpdateBrokerAdmin(ctx context.Context, user *User) error
	CreateEmployee(ctx context.Context, user *User) (int64, error)
	UpdateEmployee(ctx context.Context, user *User) error
}

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID           int64     `json:"id" bun:"id,pk,autoincrement"`
	AuthId       int64     `json:"auth_id" bun:"auth_id"`
	UserName     string    `json:"user_name" bun:"user_name"`
	UserType     string    `json:"user_type" bun:"user_type"`
	EmailAddress string    `json:"email_address" bun:"email_address"`
	PhoneNumber  string    `json:"phone_number" bun:"phone_number"`
	CountryCode  string    `json:"country_code" bun:"country_code"`
	CanLogin     bool      `json:"can_login" bun:"can_login"`
	Nid          string    `json:"nid" bun:"nid"`
	IsVerified   bool      `json:"is_verified" bun:"is_verified"`
	IsEnabled    bool      `json:"is_enabled" bun:"is_enabled"`
	LastLogin    time.Time `json:"-" bun:"last_login,nullzero"`
	ExpireAt     time.Time `json:"-" bun:"expire_at,nullzero"`
	CreatedAt    time.Time `json:"created_at" bun:"created_at,default:current_timestamp"`
	UpdatedAt    time.Time `json:"updated_at" bun:"updated_at,nullzero"`
	DeletedAt    time.Time `json:"-" bun:"deleted_at,nullzero"`
}

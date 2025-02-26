package model

import (
	"context"
	"github.com/uptrace/bun"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"time"
	//"time"
)

type Branches interface {
	CreateBranch(ctx context.Context, branch *Branch) error
	UpdateBranch(ctx context.Context, branch *Branch) error
	GetBranches(ctx context.Context, req *grpc.GetBranchesRequest) ([]*Branch, int64, error)
	GetBranchById(ctx context.Context, req *grpc.GetBranchByIdRequest) (*grpc.GetBranchByIdResponse, error)
	DeleteBranch(ctx context.Context, branch *Branch) error
}

type Branch struct {
	bun.BaseModel `bun:"table:branches"`

	ID              int64     `json:"id" bun:"id,pk,autoincrement"`
	BrokerHouseId   int64     `json:"broker_house_id" bun:"broker_house_id"`
	BranchName      string    `json:"branch_name" bun:"branch_name"`
	ShortName       string    `json:"short_name" bun:"short_name"`
	BranchType      string    `json:"branch_type" bun:"branch_type"`
	Description     string    `json:"description" bun:"description"`
	Address         string    `json:"address" bun:"address"`
	PhoneNumber     string    `json:"phone_number" bun:"phone_number"`
	CountryCode     string    `json:"country-code" bun:"country_code"`
	TelephoneNumber string    `json:"telephone_number" bun:"telephone_number"`
	EmailAddress    string    `json:"email_address" bun:"email_address"`
	ValidCurrency   string    `json:"valid_currency" bun:"valid_currency"`
	Status          string    `json:"status" bun:"status"`
	IsActive        bool      `json:"is_active" bun:"is_active"`
	IsEnabled       bool      `json:"is_enabled" bun:"is_enabled"`
	ExpireAt        time.Time `json:"-" bun:"expire_at,nullzero"`
	CreatedAt       time.Time `json:"created_at" bun:"created_at,default:current_timestamp"`
	UpdatedAt       time.Time `json:"updated_at" bun:"updated_at,nullzero"`
	DeletedAt       time.Time `json:"-" bun:"deleted_at,nullzero"`
}

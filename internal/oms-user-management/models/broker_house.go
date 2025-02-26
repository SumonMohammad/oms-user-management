package model

import (
	"context"
	"github.com/uptrace/bun"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"time"
	//"time"
)

type BrokerHouses interface {
	CreateBrokerHouse(ctx context.Context, brokerHouse *BrokerHouse) error
	UpdateBrokerHouse(ctx context.Context, brokerHouse *BrokerHouse) error
	GetBrokerHouses(ctx context.Context, req *grpc.GetBrokerHousesRequest) ([]*BrokerHouse, int64, error)
	GetBrokerHouseById(ctx context.Context, req *grpc.GetBrokerHouseByIdRequest) (*grpc.GetBrokerHouseByIdResponse, error)
	DeleteBrokerHouse(ctx context.Context, brokerHouse *BrokerHouse) error
}

type BrokerHouse struct {
	bun.BaseModel `bun:"table:broker_houses"`

	ID              int64     `json:"id" bun:"id,pk,autoincrement"`
	BrokerHouseName string    `json:"broker_house_name" bun:"broker_house_name"`
	ShortName       string    `json:"short_name" bun:"short_name"`
	Description     string    `json:"description" bun:"description"`
	Address         string    `json:"address" bun:"address"`
	PhoneNumber     string    `json:"phone_number" bun:"phone_number"`
	CountryCode     string    `json:"country-code" bun:"country_code"`
	TelephoneNumber string    `json:"telephone_number" bun:"telephone_number"`
	EmailAddress    string    `json:"email_address" bun:"email_address"`
	ValidCurrency   string    `json:"valid_currency" bun:"valid_currency"`
	Status          string    `json:"status" bun:"status"`
	IsEnabled       bool      `json:"is_enabled" bun:"is_enabled"`
	ExpireAt        time.Time `json:"-" bun:"expire_at,nullzero"`
	CreatedAt       time.Time `json:"created_at" bun:"created_at,default:current_timestamp"`
	UpdatedAt       time.Time `json:"updated_at" bun:"updated_at,nullzero"`
	DeletedAt       time.Time `json:"-" bun:"deleted_at,nullzero"`
}

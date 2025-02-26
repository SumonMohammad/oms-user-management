package pg

import (
	"context"
	"github.com/uptrace/bun"
	"gitlab.techetronventures.com/core/backend/pkg/log"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"time"
	//"google.golang.org/grpc/status"
	//"time"
)

func (db *DB) Trader() model.Traders {
	return &Trader{
		IDB: db.DB,
		log: db.log,
	}
}

func (db *Tx) Trader() model.Traders {
	return &Trader{
		IDB: db.Tx,
		log: db.log,
	}
}

type Trader struct {
	bun.IDB
	log *log.Logger
}

func (s *Trader) CreateTrader(ctx context.Context, trader *model.Trader) error {
	_, err := s.NewInsert().Model(trader).
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *Trader) UpdateTrader(ctx context.Context, trader *model.Trader) (int64, error) {
	var id int64
	trader.UpdatedAt = time.Now()
	err := s.NewUpdate().Model(trader).
		Where("user_id = ?", trader.UserId).
		ExcludeColumn("created_at").
		Returning("user_id").
		Scan(ctx, &id)

	if err != nil {
		s.log.Error(ctx, err.Error())
		return 0, err
	}
	return id, err
}

func (s *Trader) GetTraderById(ctx context.Context, userId int64) (*model.UserWithTraderType, error) {
	userInfoWithTraderType := new(model.UserWithTraderType)

	err := s.NewSelect().Model(&model.Investor{}).
		TableExpr("users AS u").
		ColumnExpr("u.id AS trader_user_id, u.user_name, u.email_address, u.phone_number, u.country_code").
		ColumnExpr("t.branch_id, t.read_only, t.can_trade, t.is_active, t.is_deleted, t.licence_number").
		Join("JOIN traders AS t ON t.user_id = u.id").
		Where("t.user_id = ?", userId). // Fix: Use "t" alias instead of "traders"
		Scan(ctx, userInfoWithTraderType)

	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}
	return userInfoWithTraderType, nil
}

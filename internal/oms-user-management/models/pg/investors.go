package pg

import (
	"context"
	"github.com/uptrace/bun"
	"gitlab.techetronventures.com/core/backend/pkg/log"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"go.uber.org/zap"
	"time"
	//"google.golang.org/grpc/status"
	//"time"
)

func (db *DB) Investor() model.Investors {
	return &Investor{
		IDB: db.DB,
		log: db.log,
	}
}

func (db *Tx) Investor() model.Investors {
	return &Investor{
		IDB: db.Tx,
		log: db.log,
	}
}

type Investor struct {
	bun.IDB
	log *log.Logger
}

func (s *Investor) CreateInvestor(ctx context.Context, investor *model.Investor) error {
	_, err := s.NewInsert().Model(investor).
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *Investor) UpdateInvestor(ctx context.Context, investor *model.Investor) (int64, error) {
	var id int64
	investor.UpdatedAt = time.Now()
	err := s.NewUpdate().Model(investor).
		Where("user_id = ?", investor.UserId).
		ExcludeColumn("created_at").
		Returning("user_id").
		Scan(ctx, &id)

	if err != nil {
		s.log.Error(ctx, err.Error())
		return 0, err
	}
	return id, err
}

func (s *Investor) GetInvestorById(ctx context.Context, userId int64) (*model.UserWithInvestorType, error) {
	investorAndOmsUserInfo := new(model.UserWithInvestorType)

	err := s.NewSelect().Model(&model.Investor{}).
		TableExpr("users AS u").
		ColumnExpr("u.id AS investor_user_id, u.user_name, u.email_address, u.phone_number, u.country_code").
		ColumnExpr("i.primary_tws_id, i.client_code, i.bo_account_number").
		Join("JOIN investors AS i ON i.user_id = u.id").
		Where("i.user_id = ?", userId). // Fix: Use "i" alias instead of "investors"
		Scan(ctx, investorAndOmsUserInfo)

	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}
	return investorAndOmsUserInfo, nil
}

func (s *Investor) GetInvestors(ctx context.Context, page, limit int32) ([]*model.UserWithInvestorType, int64, error) {
	var results []*model.UserWithInvestorType
	offset := (page - 1) * limit

	count, err := s.NewSelect().
		TableExpr("users AS u").
		Join("INNER JOIN investors AS i ON i.user_id = u.id").
		ColumnExpr("u.id AS investor_user_id").
		Column("u.user_name", "u.email_address", "u.phone_number", "u.country_code").
		Column("i.primary_tws_id", "i.client_code", "i.bo_account_number").
		Order("u.created_at DESC").
		Limit(int(limit)).
		Offset(int(offset)).
		ScanAndCount(ctx, &results)

	if err != nil {
		s.log.Error(ctx, "Failed to get OMS Users with Investors", zap.Error(err))
		return nil, 0, err
	}

	return results, int64(count), nil
}

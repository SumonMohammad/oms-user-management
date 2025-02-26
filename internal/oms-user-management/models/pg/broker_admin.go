package pg

import (
	"context"
	"github.com/uptrace/bun"
	"gitlab.techetronventures.com/core/backend/pkg/log"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"time"
	//"time"
)

func (db *DB) BrokerAdmin() model.BrokerAdmins {
	return &BrokerAdmin{
		IDB: db.DB,
		log: db.log,
	}
}

func (db *Tx) BrokerAdmin() model.BrokerAdmins {
	return &BrokerAdmin{
		IDB: db.Tx,
		log: db.log,
	}
}

type BrokerAdmin struct {
	bun.IDB
	log *log.Logger
}

func (s *BrokerAdmin) CreateBrokerAdmin(ctx context.Context, brokerAdmin *model.BrokerAdmin) error {
	_, err := s.NewInsert().Model(brokerAdmin).
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *BrokerAdmin) UpdateBrokerAdmin(ctx context.Context, brokerAdmin *model.BrokerAdmin) (int64, error) {
	var id int64
	brokerAdmin.UpdatedAt = time.Now()
	err := s.NewUpdate().Model(brokerAdmin).
		Where("user_id = ?", brokerAdmin.UserId).
		ExcludeColumn("created_at").
		Returning("user_id").
		Scan(ctx, &id)

	if err != nil {
		s.log.Error(ctx, err.Error())
		return 0, err
	}
	return id, err
}

func (s *BrokerAdmin) GetBrokerAdminById(ctx context.Context, userId int64) (*model.UserWithBrokerAdminType, error) {
	userWithBrokerAdminType := new(model.UserWithBrokerAdminType)

	err := s.NewSelect().Model(&model.BrokerAdmin{}).
		TableExpr("users AS u").
		ColumnExpr("u.id AS employee_user_id, u.user_name, u.email_address, u.phone_number, u.country_code").
		ColumnExpr("ba.branch_id, ba.can_trade, ba.read_only, ba.is_isolated_user, ba.is_deleted").
		Join("JOIN broker_admin AS ba ON ba.user_id = u.id").
		Where("ba.user_id = ?", userId). // Fix: Use "e" alias instead of "employee"
		Scan(ctx, userWithBrokerAdminType)

	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}
	return userWithBrokerAdminType, nil
}

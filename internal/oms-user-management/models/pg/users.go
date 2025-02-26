package pg

import (
	"context"
	"github.com/uptrace/bun"
	"gitlab.techetronventures.com/core/backend/pkg/log"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"time"
)

func (db *DB) User() model.Users {
	return &User{
		IDB: db.DB,
		log: db.log,
	}
}

func (db *Tx) User() model.Users {
	return &User{
		IDB: db.Tx,
		log: db.log,
	}
}

type User struct {
	bun.IDB
	log *log.Logger
}

func (s *User) CreateInvestor(ctx context.Context, user *model.User) (int64, error) {
	var id int64
	err := s.NewInsert().Model(user).Returning("id").Scan(ctx, &id)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return 0, err
	}
	return id, err
}

func (s *User) UpdateInvestor(ctx context.Context, user *model.User) error {
	user.UpdatedAt = time.Now()
	_, err := s.NewUpdate().Model(user).
		Where("id = ?", user.ID).
		ExcludeColumn("created_at").
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *User) CreateTrader(ctx context.Context, user *model.User) (int64, error) {
	var id int64
	err := s.NewInsert().Model(user).Returning("id").Scan(ctx, &id)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return 0, err
	}
	return id, err
}

func (s *User) UpdateTrader(ctx context.Context, user *model.User) error {
	user.UpdatedAt = time.Now()
	_, err := s.NewUpdate().Model(user).
		Where("id = ?", user.ID).
		ExcludeColumn("created_at").
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *User) CreateEmployee(ctx context.Context, user *model.User) (int64, error) {
	var id int64
	err := s.NewInsert().Model(user).Returning("id").Scan(ctx, &id)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return 0, err
	}
	return id, err
}

func (s *User) UpdateEmployee(ctx context.Context, user *model.User) error {
	user.UpdatedAt = time.Now()
	_, err := s.NewUpdate().Model(user).
		Where("id = ?", user.ID).
		ExcludeColumn("created_at").
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *User) CreateBrokerAdmin(ctx context.Context, user *model.User) (int64, error) {
	var id int64
	err := s.NewInsert().Model(user).Returning("id").Scan(ctx, &id)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return 0, err
	}
	return id, err
}

func (s *User) UpdateBrokerAdmin(ctx context.Context, user *model.User) error {
	user.UpdatedAt = time.Now()
	_, err := s.NewUpdate().Model(user).
		Where("id = ?", user.ID).
		ExcludeColumn("created_at").
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

package pg

import (
	"context"
	"github.com/uptrace/bun"
	"gitlab.techetronventures.com/core/backend/pkg/log"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"time"
	//"time"
)

func (db *DB) Employee() model.Employees {
	return &Employee{
		IDB: db.DB,
		log: db.log,
	}
}

func (db *Tx) Employee() model.Employees {
	return &Employee{
		IDB: db.Tx,
		log: db.log,
	}
}

type Employee struct {
	bun.IDB
	log *log.Logger
}

func (s *Employee) CreateEmployee(ctx context.Context, employee *model.Employee) error {
	_, err := s.NewInsert().Model(employee).
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *Employee) UpdateEmployee(ctx context.Context, employee *model.Employee) (int64, error) {
	var id int64
	employee.UpdatedAt = time.Now()
	err := s.NewUpdate().Model(employee).
		Where("user_id = ?", employee.UserId).
		ExcludeColumn("created_at").
		Returning("user_id").
		Scan(ctx, &id)

	if err != nil {
		s.log.Error(ctx, err.Error())
		return 0, err
	}
	return id, err
}

func (s *Employee) GetEmployeeById(ctx context.Context, userId int64) (*model.UserWithEmployeeType, error) {
	userWithEmployeeType := new(model.UserWithEmployeeType)

	err := s.NewSelect().Model(&model.Employee{}).
		TableExpr("users AS u").
		ColumnExpr("u.id AS employee_user_id, u.user_name, u.email_address, u.phone_number, u.country_code").
		ColumnExpr("e.branch_id, e.designation").
		Join("JOIN employees AS e ON e.user_id = u.id").
		Where("e.user_id = ?", userId). // Fix: Use "e" alias instead of "employee"
		Scan(ctx, userWithEmployeeType)

	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}
	return userWithEmployeeType, nil
}

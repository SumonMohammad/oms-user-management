package model

import (
	"context"
	"github.com/uptrace/bun"
	"time"
	//"time"
)

type Employees interface {
	UpdateEmployee(ctx context.Context, employee *Employee) (int64, error)
	CreateEmployee(ctx context.Context, employee *Employee) error
	GetEmployeeById(ctx context.Context, userId int64) (*UserWithEmployeeType, error)
	//GetEmployees(ctx context.Context, page, limit int32) ([]*UserWithEmployeeType, int64, error)
}

type Employee struct {
	bun.BaseModel `bun:"table:employees"`

	ID          int64     `json:"id" bun:"id,pk,autoincrement"`
	UserId      int64     `json:"user_id" bun:"user_id"`
	BranchId    int64     `json:"branch_id" bun:"branch_id"`
	Designation string    `json:"designation" bun:"designation"`
	Description string    `json:"description" bun:"description"`
	ExpireAt    time.Time `json:"-" bun:"expire_at,nullzero"`
	CreatedAt   time.Time `json:"created_at" bun:"created_at,default:current_timestamp"`
	UpdatedAt   time.Time `json:"updated_at" bun:"updated_at,nullzero"`
	DeletedAt   time.Time `json:"-" bun:"deleted_at,nullzero"`
}

type UserWithEmployeeType struct {
	EmployeeUserId int64  `json:"employee_user_id" bun:"employee_user_id"`
	UserName       string `json:"user_name" bun:"user_name"`
	EmailAddress   string `json:"email_address" bun:"email_address"`
	PhoneNumber    string `json:"phone_number" bun:"phone_number"`
	CountryCode    string `json:"country_code" bun:"country_code"`
	BranchId       int64  `json:"branch_id" bun:"branch_id"`
	Designation    string `json:"designation" bun:"designation"`
}

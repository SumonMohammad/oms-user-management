package service

import (
	"context"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"go.uber.org/zap"
)

type Employees interface {
	CreateEmployee(ctx context.Context, req *grpc.CreateEmployeeRequest) (*grpc.CreateEmployeeResponse, error)
	UpdateEmployee(ctx context.Context, req *grpc.UpdateEmployeeRequest) (*grpc.UpdateEmployeeResponse, error)
	GetEmployeeById(ctx context.Context, userId int64) (*grpc.GetEmployeeByIdResponse, error)
	//GetInvestors(ctx context.Context, page int32, limit int32) (*grpc.GetInvestorsResponse, error)
}

type Employee struct {
	service *OmsUserManagementService
}

type EmployeeReceiver struct {
	*OmsUserManagementService
}

func (ms *OmsUserManagementService) Employee() Employees {
	return &EmployeeReceiver{
		ms,
	}
}

func (s *OmsUserManagementService) CreateEmployee(ctx context.Context, req *grpc.CreateEmployeeRequest) (*grpc.CreateEmployeeResponse, error) {
	var userId int64
	var err error
	err = s.db.InTx(ctx, func(ctx context.Context, tx model.Repository) error {

		// Create OMSUser struct
		user := &model.User{
			UserName:     req.UserName,
			AuthId:       req.AuthId,
			UserType:     req.UserType.String(),
			EmailAddress: req.EmailAddress,
			PhoneNumber:  req.PhoneNumber,
			CountryCode:  req.CountryCode,
			CanLogin:     req.CanLogin,
			Nid:          req.Nid,
			IsEnabled:    req.IsEnabled,
			IsVerified:   req.IsVerified,
		}

		// Insert OMSUser and retrieve the generated ID
		userId, err = tx.User().CreateEmployee(ctx, user)
		if err != nil {
			msg := "Failed to create user"
			s.log.Error(ctx, msg, zap.Error(err))
			return err
		}

		// Create Investor struct
		employee := &model.Employee{
			UserId:      userId, // Assign the newly created OMS user ID
			BranchId:    req.BranchId,
			Designation: req.Designation,
			Description: req.Description,
		}

		err = tx.Employee().CreateEmployee(ctx, employee)
		if err != nil {
			msg := "Failed to create employee"
			s.log.Error(ctx, msg, zap.Error(err))
			return err
		}

		return err
	})

	return &grpc.CreateEmployeeResponse{
		Code: 0,
	}, err
}

func (s *OmsUserManagementService) UpdateEmployee(ctx context.Context, req *grpc.UpdateEmployeeRequest) (*grpc.UpdateEmployeeResponse, error) {

	var userId int64
	var err error
	err = s.db.InTx(ctx, func(ctx context.Context, tx model.Repository) error {

		employee := &model.Employee{
			UserId:      userId, // Assign the newly created OMS user ID
			BranchId:    req.BranchId,
			Designation: req.Designation,
			Description: req.Description,
		}

		userId, err = tx.Employee().UpdateEmployee(ctx, employee)
		if err != nil {
			msg := "Failed to update Employee"
			s.log.Error(ctx, msg, zap.Error(err))
			return err
		}

		user := &model.User{
			UserName:     req.UserName,
			AuthId:       req.AuthId,
			UserType:     req.UserType.String(),
			EmailAddress: req.EmailAddress,
			PhoneNumber:  req.PhoneNumber,
			CountryCode:  req.CountryCode,
			CanLogin:     req.CanLogin,
			Nid:          req.Nid,
			IsEnabled:    req.IsEnabled,
			IsVerified:   req.IsVerified,
		}

		err = tx.User().UpdateEmployee(ctx, user)
		if err != nil {
			msg := "Failed to update user"
			s.log.Error(ctx, msg, zap.Error(err))
			return err
		}

		return nil
	})

	return &grpc.UpdateEmployeeResponse{
		Code: 0,
	}, err
}

func (s *OmsUserManagementService) GetEmployeeById(ctx context.Context, userId int64) (*grpc.GetEmployeeByIdResponse, error) {
	userInfoWithEmployeeType, err := s.db.Employee().GetEmployeeById(ctx, userId)
	if err != nil {
		s.log.Error(ctx, "Failed to get employee", zap.Error(err))
		return nil, err
	}

	return &grpc.GetEmployeeByIdResponse{
		UserId:       userInfoWithEmployeeType.EmployeeUserId, // Now accessible
		UserName:     userInfoWithEmployeeType.UserName,
		EmailAddress: userInfoWithEmployeeType.EmailAddress,
		PhoneNumber:  userInfoWithEmployeeType.PhoneNumber,
		CountryCode:  userInfoWithEmployeeType.CountryCode,
		BranchId:     userInfoWithEmployeeType.BranchId,
		Designation:  userInfoWithEmployeeType.Designation,
	}, err
}

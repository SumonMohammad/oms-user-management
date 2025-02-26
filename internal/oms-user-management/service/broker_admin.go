package service

import (
	"context"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"go.uber.org/zap"
)

type BrokerAdmins interface {
	CreateBrokerAdmin(ctx context.Context, req *grpc.CreateBrokerAdminRequest) (*grpc.CreateBrokerAdminResponse, error)
	UpdateBrokerAdmin(ctx context.Context, req *grpc.UpdateBrokerAdminRequest) (*grpc.UpdateBrokerAdminResponse, error)
	GetBrokerAdminById(ctx context.Context, userId int64) (*grpc.GetBrokerAdminByIdResponse, error)
	//GetInvestors(ctx context.Context, page int32, limit int32) (*grpc.GetInvestorsResponse, error)
}

type BrokerAdmin struct {
	service *OmsUserManagementService
}

type BrokerAdminReceiver struct {
	*OmsUserManagementService
}

func (ms *OmsUserManagementService) BrokerAdmin() BrokerAdmins {
	return &BrokerAdminReceiver{
		ms,
	}
}

func (s *OmsUserManagementService) CreateBrokerAdmin(ctx context.Context, req *grpc.CreateBrokerAdminRequest) (*grpc.CreateBrokerAdminResponse, error) {
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
		userId, err = tx.User().CreateBrokerAdmin(ctx, user)
		if err != nil {
			msg := "Failed to create user"
			s.log.Error(ctx, msg, zap.Error(err))
			return err
		}

		// Create Investor struct
		brokerAdmin := &model.BrokerAdmin{
			UserId:         userId, // Assign the newly created OMS user ID
			BranchId:       req.BranchId,
			CanTrade:       req.CanTrade,
			IsIsolatedUser: req.IsIsolatedUser,
			ReadOnly:       req.ReadOnly,
			IsDeleted:      req.IsDeleted,
			Status:         req.Status.String(),
		}

		err = tx.BrokerAdmin().CreateBrokerAdmin(ctx, brokerAdmin)
		if err != nil {
			msg := "Failed to create BrokerAdmin"
			s.log.Error(ctx, msg, zap.Error(err))
			return err
		}

		return err
	})

	return &grpc.CreateBrokerAdminResponse{
		Code: 0,
	}, err
}

func (s *OmsUserManagementService) UpdateBrokerAdmin(ctx context.Context, req *grpc.UpdateBrokerAdminRequest) (*grpc.UpdateBrokerAdminResponse, error) {

	var userId int64
	var err error
	err = s.db.InTx(ctx, func(ctx context.Context, tx model.Repository) error {

		brokerAdmin := &model.BrokerAdmin{
			UserId:         userId, // Assign the newly created OMS user ID
			BranchId:       req.BranchId,
			CanTrade:       req.CanTrade,
			IsIsolatedUser: req.IsIsolatedUser,
			ReadOnly:       req.ReadOnly,
			IsDeleted:      req.IsDeleted,
			Status:         req.Status.String(),
		}

		userId, err = tx.BrokerAdmin().UpdateBrokerAdmin(ctx, brokerAdmin)
		if err != nil {
			msg := "Failed to update BrokerAdmin"
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

		err = tx.User().UpdateBrokerAdmin(ctx, user)
		if err != nil {
			msg := "Failed to update user"
			s.log.Error(ctx, msg, zap.Error(err))
			return err
		}

		return nil
	})

	return &grpc.UpdateBrokerAdminResponse{
		Code: 0,
	}, err
}

func (s *OmsUserManagementService) GetBrokerAdminById(ctx context.Context, userId int64) (*grpc.GetBrokerAdminByIdResponse, error) {
	userInfoWithBrokerAdminType, err := s.db.BrokerAdmin().GetBrokerAdminById(ctx, userId)
	if err != nil {
		s.log.Error(ctx, "Failed to get BrokerAdmin", zap.Error(err))
		return nil, err
	}

	return &grpc.GetBrokerAdminByIdResponse{
		UserId:         userInfoWithBrokerAdminType.BrokerUserId, // Now accessible
		UserName:       userInfoWithBrokerAdminType.UserName,
		EmailAddress:   userInfoWithBrokerAdminType.EmailAddress,
		PhoneNumber:    userInfoWithBrokerAdminType.PhoneNumber,
		CountryCode:    userInfoWithBrokerAdminType.CountryCode,
		BranchId:       userInfoWithBrokerAdminType.BranchId,
		CanTrade:       userInfoWithBrokerAdminType.CanTrade,
		IsIsolatedUser: userInfoWithBrokerAdminType.IsIsolatedUser,
	}, err
}

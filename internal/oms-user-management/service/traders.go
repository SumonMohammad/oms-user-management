package service

import (
	"context"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"go.uber.org/zap"
)

type Traders interface {
	CreateTrader(ctx context.Context, req *grpc.CreateTraderRequest) (*grpc.CreateTraderResponse, error)
	UpdateTrader(ctx context.Context, req *grpc.UpdateTraderRequest) (*grpc.UpdateTraderResponse, error)
	GetTraderById(ctx context.Context, userId int64) (*grpc.GetTraderByIdResponse, error)
	//GetInvestors(ctx context.Context, page int32, limit int32) (*grpc.GetInvestorsResponse, error)
}

type Trader struct {
	service *OmsUserManagementService
}

type TraderReceiver struct {
	*OmsUserManagementService
}

func (ms *OmsUserManagementService) Trader() Traders {
	return &TraderReceiver{
		ms,
	}
}

func (s *OmsUserManagementService) CreateTrader(ctx context.Context, req *grpc.CreateTraderRequest) (*grpc.CreateTraderResponse, error) {
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
		userId, err = tx.User().CreateTrader(ctx, user)
		if err != nil {
			msg := "Failed to create Trader"
			s.log.Error(ctx, msg, zap.Error(err))
			return err
		}

		// Create Investor struct
		trader := &model.Trader{
			UserId:        userId, // Assign the newly created OMS user ID
			BranchId:      req.BranchId,
			IsActive:      req.IsActive,
			IsDeleted:     req.IsDeleted,
			ReadOnly:      req.ReadOnly,
			Status:        req.Status.String(),
			CanTrade:      req.CanTrade,
			LicenceNumber: req.LicenceNumber,
		}

		err = tx.Trader().CreateTrader(ctx, trader)
		if err != nil {
			msg := "Failed to create Trader"
			s.log.Error(ctx, msg, zap.Error(err))
			return err
		}

		return err
	})

	return &grpc.CreateTraderResponse{
		Code: 0,
	}, err
}

func (s *OmsUserManagementService) UpdateTrader(ctx context.Context, req *grpc.UpdateTraderRequest) (*grpc.UpdateTraderResponse, error) {

	var userId int64
	var err error
	err = s.db.InTx(ctx, func(ctx context.Context, tx model.Repository) error {

		trader := &model.Trader{
			UserId:        userId, // Assign the newly created OMS user ID
			BranchId:      req.BranchId,
			IsActive:      req.IsActive,
			IsDeleted:     req.IsDeleted,
			ReadOnly:      req.ReadOnly,
			Status:        req.Status.String(),
			CanTrade:      req.CanTrade,
			LicenceNumber: req.LicenceNumber,
		}

		userId, err = tx.Trader().UpdateTrader(ctx, trader)
		if err != nil {
			msg := "Failed to update Trader"
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

		err = tx.User().UpdateTrader(ctx, user)
		if err != nil {
			msg := "Failed to update user"
			s.log.Error(ctx, msg, zap.Error(err))
			return err
		}

		return nil
	})

	return &grpc.UpdateTraderResponse{
		Code: 0,
	}, err
}

func (s *OmsUserManagementService) GetTraderById(ctx context.Context, userId int64) (*grpc.GetTraderByIdResponse, error) {
	userInfoWithTraderType, err := s.db.Trader().GetTraderById(ctx, userId)
	if err != nil {
		s.log.Error(ctx, "Failed to get Trader", zap.Error(err))
		return nil, err
	}

	return &grpc.GetTraderByIdResponse{
		UserId:        userInfoWithTraderType.TraderUserId, // Now accessible
		UserName:      userInfoWithTraderType.UserName,
		EmailAddress:  userInfoWithTraderType.EmailAddress,
		PhoneNumber:   userInfoWithTraderType.PhoneNumber,
		CountryCode:   userInfoWithTraderType.CountryCode,
		BranchId:      userInfoWithTraderType.BranchId,
		LicenceNumber: userInfoWithTraderType.LicenceNumber,
		CanTrade:      userInfoWithTraderType.CanTrade,
		IsActive:      userInfoWithTraderType.IsActive,
		ReadOnly:      userInfoWithTraderType.ReadOnly,
	}, err
}

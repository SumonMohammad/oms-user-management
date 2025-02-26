package service

import (
	"context"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"go.uber.org/zap"
)

type Investors interface {
	CreateInvestor(ctx context.Context, req *grpc.CreateInvestorRequest) (*grpc.CreateInvestorResponse, error)
	UpdateInvestor(ctx context.Context, req *grpc.UpdateInvestorRequest) (*grpc.UpdateInvestorResponse, error)
	GetInvestorById(ctx context.Context, userId int64) (*grpc.GetInvestorByIdResponse, error)
	GetInvestors(ctx context.Context, page int32, limit int32) (*grpc.GetInvestorsResponse, error)
}

type Investor struct {
	service *OmsUserManagementService
}

type InvestorReceiver struct {
	*OmsUserManagementService
}

func (ms *OmsUserManagementService) Investor() Investors {
	return &InvestorReceiver{
		ms,
	}
}

func (s *OmsUserManagementService) CreateInvestor(ctx context.Context, req *grpc.CreateInvestorRequest) (*grpc.CreateInvestorResponse, error) {
	var userId int64
	var err error
	err = s.db.InTx(ctx, func(ctx context.Context, tx model.Repository) error {

		// Create OMSUser struct
		omsUser := &model.User{
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
		userId, err = tx.User().CreateInvestor(ctx, omsUser)
		if err != nil {
			msg := "Failed to create OMS user"
			s.log.Error(ctx, msg, zap.Error(err))
			return err
		}

		// Create Investor struct
		investor := &model.Investor{
			UserId:          userId, // Assign the newly created OMS user ID
			PrimaryTwsId:    req.PrimaryTwsId,
			SecondaryTwsId:  req.SecondaryTwsId,
			ClientCode:      req.ClientCode,
			BoAccountNumber: req.BoAccountNumber,
			Status:          req.Status.String(),
			CanTrade:        req.CanTrade,
			ReadOnly:        req.ReadOnly,
			IsDeleted:       req.IsDeleted,
		}

		err = tx.Investor().CreateInvestor(ctx, investor)
		if err != nil {
			msg := "Failed to create Investor"
			s.log.Error(ctx, msg, zap.Error(err))
			return err
		}

		return err
	})

	return &grpc.CreateInvestorResponse{
		UserId:          userId,
		UserName:        req.UserName,
		PrimaryTwsId:    req.PrimaryTwsId,
		ClientCode:      req.ClientCode,
		BoAccountNumber: req.BoAccountNumber,
	}, err
}

func (s *OmsUserManagementService) UpdateInvestor(ctx context.Context, req *grpc.UpdateInvestorRequest) (*grpc.UpdateInvestorResponse, error) {

	var userId int64
	var err error
	err = s.db.InTx(ctx, func(ctx context.Context, tx model.Repository) error {

		investor := &model.Investor{
			UserId:          req.UserId,
			PrimaryTwsId:    req.PrimaryTwsId,
			SecondaryTwsId:  req.SecondaryTwsId,
			ClientCode:      req.ClientCode,
			BoAccountNumber: req.BoAccountNumber,
			Status:          req.Status.String(),
			CanTrade:        req.CanTrade,
			ReadOnly:        req.ReadOnly,
			IsDeleted:       req.IsDeleted,
		}

		userId, err = tx.Investor().UpdateInvestor(ctx, investor)
		if err != nil {
			msg := "Failed to update Investor"
			s.log.Error(ctx, msg, zap.Error(err))
			return err
		}

		omsUser := &model.User{
			ID:           userId, // Using user_id from Investor table
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

		err = tx.User().UpdateInvestor(ctx, omsUser)
		if err != nil {
			msg := "Failed to update OMS user"
			s.log.Error(ctx, msg, zap.Error(err))
			return err
		}

		return nil
	})

	return &grpc.UpdateInvestorResponse{
		Code: 0,
	}, err
}

func (s *OmsUserManagementService) GetInvestorById(ctx context.Context, userId int64) (*grpc.GetInvestorByIdResponse, error) {
	OmsUserWithInvestorInfo, err := s.db.Investor().GetInvestorById(ctx, userId)
	if err != nil {
		s.log.Error(ctx, "Failed to get OMS user", zap.Error(err))
		return nil, err
	}

	return &grpc.GetInvestorByIdResponse{
		UserId:          OmsUserWithInvestorInfo.InvestorUserId, // Now accessible
		UserName:        OmsUserWithInvestorInfo.UserName,
		EmailAddress:    OmsUserWithInvestorInfo.EmailAddress,
		PhoneNumber:     OmsUserWithInvestorInfo.PhoneNumber,
		CountryCode:     OmsUserWithInvestorInfo.CountryCode,
		PrimaryTwsId:    OmsUserWithInvestorInfo.PrimaryTwsId,
		ClientCode:      OmsUserWithInvestorInfo.ClientCode,
		BoAccountNumber: OmsUserWithInvestorInfo.BoAccountNumber,
	}, err
}

func (s *OmsUserManagementService) GetInvestors(ctx context.Context, page, limit int32) (*grpc.GetInvestorsResponse, error) {
	var usersWithInvestors []*grpc.UserWithInvestorType

	// Fetch investors with pagination
	results, totalCount, err := s.db.Investor().GetInvestors(ctx, page, limit)
	if err != nil {
		s.log.Error(ctx, "Failed to get OMS users with Investors", zap.Error(err))
		return nil, err
	}

	// Map database results to gRPC response struct
	for _, user := range results {
		usersWithInvestors = append(usersWithInvestors, &grpc.UserWithInvestorType{
			UserId:          user.InvestorUserId,
			UserName:        user.UserName,
			EmailAddress:    user.EmailAddress,
			PhoneNumber:     user.PhoneNumber,
			CountryCode:     user.CountryCode,
			PrimaryTwsId:    user.PrimaryTwsId,
			ClientCode:      user.ClientCode,
			BoAccountNumber: user.BoAccountNumber,
		})
	}

	return &grpc.GetInvestorsResponse{
		Investors:  usersWithInvestors,
		TotalCount: totalCount,
	}, nil
}

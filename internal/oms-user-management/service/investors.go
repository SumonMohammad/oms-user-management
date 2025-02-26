package service

import (
	"context"
	//"fmt"
	model "gitlab.techetronventures.com/core/grpctest/internal/grpctest/models"
	//pg "gitlab.techetronventures.com/core/grpctest/internal/grpctest/models/pg"
	"gitlab.techetronventures.com/core/grpctest/pkg/grpc"
	"go.uber.org/zap"
)

type Investors interface {
	CreateInvestor(ctx context.Context, req *grpc.CreateInvestorRequest) (*grpc.CreateInvestorResponse, error)
	UpdateInvestor(ctx context.Context, req *grpc.UpdateInvestorRequest) (*grpc.UpdateInvestorResponse, error)
	GetInvestorByIdOrEmail(ctx context.Context, req *grpc.GetInvestorByIdOrEmailRequest) (*grpc.GetInvestorByIdOrEmailResponse, error)
	GetInvestors(ctx context.Context, req *grpc.GetInvestorsRequest) (*grpc.GetInvestorsResponse, error)
}

type Investor struct {
	service *GrpctestService
}

type InvestorReceiver struct {
	*GrpctestService
}

func (ms *GrpctestService) Investor() Investors {
	return &InvestorReceiver{
		ms,
	}
}

func (s *InvestorReceiver) CreateInvestor(ctx context.Context, req *grpc.CreateInvestorRequest) (*grpc.CreateInvestorResponse, error) {

	investor := &model.Investor{
		UserId:          req.UserId,
		PrimaryTwsId:    req.PrimaryTwsId,
		SecondaryTwsId:  req.SecondaryTwsId,
		BoAccountNumber: req.BoAccountNumber,
		Status:          req.Status.String(),
		CanTrade:        req.CanTrade,
		ReadOnly:        req.ReadOnly,
		IsEnabled:       req.IsEnabled,
		IsDeleted:       req.IsDeleted,
	}

	// Call the database method to create the trader
	err := s.db.Investor().CreateInvestor(ctx, investor)
	if err != nil {
		msg := "failed to create investor"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.CreateInvestorResponse{
		Code: 0,
	}, nil
}

func (s *InvestorReceiver) UpdateInvestor(ctx context.Context, req *grpc.UpdateInvestorRequest) (*grpc.UpdateInvestorResponse, error) {

	investor := &model.Investor{
		ID:              req.Id,
		UserId:          req.UserId,
		PrimaryTwsId:    req.PrimaryTwsId,
		SecondaryTwsId:  req.SecondaryTwsId,
		BoAccountNumber: req.BoAccountNumber,
		Status:          req.Status.String(),
		CanTrade:        req.CanTrade,
		ReadOnly:        req.ReadOnly,
		IsEnabled:       req.IsEnabled,
		IsDeleted:       req.IsDeleted,
	}

	// Call the database method to create the trader
	err := s.db.Investor().UpdateInvestor(ctx, investor)
	if err != nil {
		msg := "failed to update trader"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.UpdateInvestorResponse{
		Code: 0,
	}, nil
}

func (s *InvestorReceiver) GetInvestorByIdOrEmail(ctx context.Context, req *grpc.GetInvestorByIdOrEmailRequest) (*grpc.GetInvestorByIdOrEmailResponse, error) {
	// Call the database layer to fetch the trader by ID or email

	investor, err := s.db.Investor().GetInvestorByIdOrEmail(ctx, req)

	if err != nil {
		msg := "Failed to fetch investor by ID or email"
		s.log.Error(ctx, msg, zap.Error(err))
	}
	return &grpc.GetInvestorByIdOrEmailResponse{
		Id:              investor.Id,
		UserId:          investor.UserId,
		PrimaryTwsId:    investor.PrimaryTwsId,
		SecondaryTwsId:  investor.SecondaryTwsId,
		BoAccountNumber: investor.BoAccountNumber,
		Status:          investor.Status,
		CanTrade:        investor.CanTrade,
		ReadOnly:        investor.ReadOnly,
		IsEnabled:       investor.IsEnabled,
		IsDeleted:       investor.IsDeleted,
	}, nil
}

func (s *InvestorReceiver) GetInvestors(ctx context.Context, req *grpc.GetInvestorsRequest) (*grpc.GetInvestorsResponse, error) {

	res, count, err := s.db.Investor().GetInvestors(ctx, req)
	if err != nil {
		msg := "failed to fetch investors"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	investors := []*grpc.GetInvestorsResponse_Investor{}
	for _, item := range res {
		var investorStatus grpc.Status

		if item.Status == "PENDING" {
			investorStatus = grpc.Status_PENDING
		} else {
			investorStatus = grpc.Status_ACTIVE
		}

		investor := &grpc.GetInvestorsResponse_Investor{
			Id:              item.ID,
			UserId:          item.UserId,
			PrimaryTwsId:    item.PrimaryTwsId,
			SecondaryTwsId:  item.SecondaryTwsId,
			BoAccountNumber: item.BoAccountNumber,
			Status:          investorStatus,
			CanTrade:        item.CanTrade,
			ReadOnly:        item.ReadOnly,
			IsEnabled:       item.IsEnabled,
			IsDeleted:       item.IsDeleted,
		}
		investors = append(investors, investor)
	}

	return &grpc.GetInvestorsResponse{
		Investors: investors,
		PaginationResponse: &grpc.PaginationInfoResponse{
			TotalRecordCount: int32(count),
		},
	}, nil

}

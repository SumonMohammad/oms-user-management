package service

import (
	"context"
	model "gitlab.techetronventures.com/core/grpctest/internal/grpctest/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.techetronventures.com/core/grpctest/pkg/grpc"
	"go.uber.org/zap"
)

type Traders interface {
	CreateTrader(ctx context.Context, req *grpc.CreateTraderRequest) (*grpc.CreateTraderResponse, error)
	UpdateTrader(ctx context.Context, req *grpc.UpdateTraderRequest) (*grpc.UpdateTraderResponse, error)
	GetTraders(ctx context.Context, req *grpc.GetTradersRequest) (*grpc.GetTradersResponse, error)
	GetTraderByIdOrEmail(ctx context.Context, req *grpc.GetTraderByIdOrEmailRequest) (*grpc.GetTraderByIdOrEmailResponse, error)
}
type Trader struct {
	service *GrpctestService
}

type TraderReceiver struct {
	*GrpctestService
}

func (ms *GrpctestService) Trader() Traders {
	return &TraderReceiver{
		ms,
	}
}

func (s *TraderReceiver) CreateTrader(ctx context.Context, req *grpc.CreateTraderRequest) (*grpc.CreateTraderResponse, error) {
	// Map the incoming gRPC request to the Trader model

	trader := &model.Trader{
		UserId:        req.UserId,
		LicenceNumber: req.LicenceNumber,
		IsActive:      req.IsActive,
		ReadOnly:      req.ReadOnly,
		TwsId:         req.TwsId,
		IsEnabled:     req.IsEnabled,
		OfficeId:      req.OfficeId,
		Status:        req.Status.String(),
		CanTrade:      req.CanTrade,
		IsDeleted:     req.IsDeleted,
	}

	// Call the database method to create the trader
	err := s.db.Trader().CreateTrader(ctx, trader)
	if err != nil {
		msg := "failed to create trader"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.CreateTraderResponse{
		Code: 0,
	}, nil
}

func (s *TraderReceiver) UpdateTrader(ctx context.Context, req *grpc.UpdateTraderRequest) (*grpc.UpdateTraderResponse, error) {

	trader := &model.Trader{
		ID:            req.Id,
		UserId:        req.UserId,
		LicenceNumber: req.LicenceNumber,
		IsActive:      req.IsActive,
		ReadOnly:      req.ReadOnly,
		TwsId:         req.TwsId,
		IsEnabled:     req.IsEnabled,
		OfficeId:      req.OfficeId,
		Status:        req.Status.String(),
		CanTrade:      req.CanTrade,
		IsDeleted:     req.IsDeleted,
	}

	// Call the database method to create the trader
	err := s.db.Trader().UpdateTrader(ctx, trader)
	if err != nil {
		msg := "failed to update trader"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.UpdateTraderResponse{
		Code: 0,
	}, nil
}

func (s *TraderReceiver) GetTraders(ctx context.Context, req *grpc.GetTradersRequest) (*grpc.GetTradersResponse, error) {

	res, count, err := s.db.Trader().GetTraders(ctx, req)
	if err != nil {
		msg := "failed to fetch traders"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	traders := []*grpc.GetTradersResponse_Trader{}
	for _, item := range res {

		var traderStatus grpc.Status

		if item.Status == "PENDING" {
			traderStatus = grpc.Status_PENDING
		} else {
			traderStatus = grpc.Status_ACTIVE
		}

		trader := &grpc.GetTradersResponse_Trader{
			Id:            item.ID,
			UserId:        item.UserId,
			LicenceNumber: item.LicenceNumber,
			IsActive:      item.IsActive,
			ReadOnly:      item.ReadOnly,
			TwsId:         item.TwsId,
			IsEnabled:     item.IsEnabled,
			OfficeId:      item.OfficeId,
			Status:        traderStatus,
			CanTrade:      item.CanTrade,
			IsDeleted:     item.IsDeleted,
		}
		traders = append(traders, trader)
	}

	return &grpc.GetTradersResponse{
		Traders: traders,
		PaginationResponse: &grpc.PaginationInfoResponse{
			TotalRecordCount: int32(count),
		},
	}, nil

}

func (s *TraderReceiver) GetTraderByIdOrEmail(ctx context.Context, req *grpc.GetTraderByIdOrEmailRequest) (*grpc.GetTraderByIdOrEmailResponse, error) {
	// Call the database layer to fetch the trader by ID or email
	trader, err := s.db.Trader().GetTraderByIdOrEmail(ctx, req)
	if err != nil {
		// Handle error if no trader is found or any other error occurs
		s.log.Error(ctx, "Failed to fetch trader by ID or email", zap.Error(err))
		if err.Error() == "Trader not found" {
			return nil, status.Error(codes.NotFound, "Trader not found")
		}
		return nil, status.Error(codes.Internal, "Failed to fetch trader")
	}

	// Map the result from the database layer to the GRPC response
	response := &grpc.GetTraderByIdOrEmailResponse{
		Id:            trader.Id,
		UserId:        trader.UserId,
		LicenceNumber: trader.LicenceNumber,
		IsActive:      trader.IsActive,
		ReadOnly:      trader.ReadOnly,
		TwsId:         trader.TwsId,
		IsEnabled:     trader.IsEnabled,
		OfficeId:      trader.OfficeId,
		Status:        trader.Status,
		CanTrade:      trader.CanTrade,
		IsDeleted:     trader.IsDeleted,
	}

	return response, nil
}

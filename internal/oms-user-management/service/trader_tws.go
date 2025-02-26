package service

import (
	"context"
	model "gitlab.techetronventures.com/core/grpctest/internal/grpctest/models"
	"gitlab.techetronventures.com/core/grpctest/pkg/grpc"
	"go.uber.org/zap"
)

type TradersTws interface {
	CreateTraderTws(ctx context.Context, req *grpc.CreateTraderTwsMapRequest) (*grpc.CreateTraderTwsMapResponse, error)
	GetTradersTws(ctx context.Context, req *grpc.GetTradersTwsMapRequest) (*grpc.GetTradersTwsMapResponse, error)
	UpdateTraderTws(ctx context.Context, req *grpc.UpdateTraderTwsMapRequest) (*grpc.UpdateTraderTwsMapResponse, error)
	DeleteTraderTws(ctx context.Context, req *grpc.DeleteTraderTwsMapRequest) (*grpc.DeleteTraderTwsMapResponse, error)
}
type TraderTws struct {
	service *GrpctestService
}

type TraderTwsReceiver struct {
	*GrpctestService
}

func (ms *GrpctestService) TraderTws() TradersTws {
	return &TraderTwsReceiver{
		ms,
	}
}

func (s *TraderTwsReceiver) CreateTraderTws(ctx context.Context, req *grpc.CreateTraderTwsMapRequest) (*grpc.CreateTraderTwsMapResponse, error) {
	// Map the incoming gRPC request to the Trader model

	traderTws := &model.TraderTws{
		TwsId:     req.TwsId,
		TraderId:  req.TraderId,
		Status:    req.Status.String(),
		IsEnabled: req.IsEnabled,
		IsDeleted: req.IsDeleted,
	}
	// Call the database method to create the trader
	err := s.db.TraderTws().CreateTraderTws(ctx, traderTws)
	if err != nil {
		msg := "failed to create trader_tws"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.CreateTraderTwsMapResponse{
		Code: 0,
	}, nil
}

func (s *TraderTwsReceiver) GetTradersTws(ctx context.Context, req *grpc.GetTradersTwsMapRequest) (*grpc.GetTradersTwsMapResponse, error) {

	res, count, err := s.db.TraderTws().GetTradersTws(ctx, req)
	if err != nil {
		msg := "failed to fetch traders_tws"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	tradersTws := []*grpc.GetTradersTwsMapResponse_TraderTws{}
	for _, item := range res {

		var twsStatus grpc.TwsStatus
		if item.Status == "ASSIGNED" {
			twsStatus = grpc.TwsStatus_ASSIGNED
		} else {
			twsStatus = grpc.TwsStatus_UNASSIGNED
		}

		traderTws := &grpc.GetTradersTwsMapResponse_TraderTws{
			Id:        item.ID,
			TwsId:     item.TwsId,
			TraderId:  item.TraderId,
			Status:    twsStatus,
			IsDeleted: item.IsDeleted,
			IsEnabled: item.IsEnabled,
		}
		tradersTws = append(tradersTws, traderTws)
	}

	return &grpc.GetTradersTwsMapResponse{
		TradersTws: tradersTws,
		PaginationResponse: &grpc.PaginationInfoResponse{
			TotalRecordCount: int32(count),
		},
	}, nil

}

func (s *TraderTwsReceiver) UpdateTraderTws(ctx context.Context, req *grpc.UpdateTraderTwsMapRequest) (*grpc.UpdateTraderTwsMapResponse, error) {

	traderTws := &model.TraderTws{
		TwsId:     req.TwsId,
		TraderId:  req.TraderId,
		Status:    req.Status.String(),
		IsEnabled: req.IsEnabled,
		IsDeleted: req.IsDeleted,
	}

	// Call database update
	err := s.db.TraderTws().UpdateTraderTws(ctx, traderTws)
	if err != nil {
		msg := "failed to update trader tws map info"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return success response
	return &grpc.UpdateTraderTwsMapResponse{
		Code: 0,
	}, nil
}

func (s *TraderTwsReceiver) DeleteTraderTws(ctx context.Context, req *grpc.DeleteTraderTwsMapRequest) (*grpc.DeleteTraderTwsMapResponse, error) {
	// Map the incoming gRPC request to the TraderTws model
	traderTws := &model.TraderTws{
		ID: req.Id,
	}

	// Call the database method to delete the trader
	err := s.db.TraderTws().DeleteTraderTws(ctx, traderTws)
	if err != nil {
		msg := "failed to delete trader tws"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.DeleteTraderTwsMapResponse{
		Code: 0,
	}, nil
}

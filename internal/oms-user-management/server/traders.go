package server

import (
	"context"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
)

func (s *OmsUserManagementServer) CreateTrader(ctx context.Context, req *grpc.CreateTraderRequest) (*grpc.CreateTraderResponse, error) {
	res, err := s.service.Trader().CreateTrader(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) UpdateTrader(ctx context.Context, req *grpc.UpdateTraderRequest) (*grpc.UpdateTraderResponse, error) {
	res, err := s.service.Trader().UpdateTrader(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetTraderById(ctx context.Context, req *grpc.GetTraderByIdRequest) (*grpc.GetTraderByIdResponse, error) {
	res, err := s.service.Trader().GetTraderById(ctx, req.UserId)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

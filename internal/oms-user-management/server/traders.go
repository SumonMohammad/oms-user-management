package server

import (
	"context"
	"gitlab.techetronventures.com/core/grpctest/pkg/grpc"
)

func (s *GrpctestServer) CreateTrader(ctx context.Context, req *grpc.CreateTraderRequest) (*grpc.CreateTraderResponse, error) {
	res, err := s.service.Trader().CreateTrader(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) UpdateTrader(ctx context.Context, req *grpc.UpdateTraderRequest) (*grpc.UpdateTraderResponse, error) {
	res, err := s.service.Trader().UpdateTrader(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) GetTraders(ctx context.Context, req *grpc.GetTradersRequest) (*grpc.GetTradersResponse, error) {
	res, err := s.service.Trader().GetTraders(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) GetTraderByIdOrEmail(ctx context.Context, req *grpc.GetTraderByIdOrEmailRequest) (*grpc.GetTraderByIdOrEmailResponse, error) {
	res, err := s.service.Trader().GetTraderByIdOrEmail(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

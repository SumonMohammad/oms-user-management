// trader_tws
package server

import (
	"context"
	"gitlab.techetronventures.com/core/grpctest/pkg/grpc"
)

func (s *GrpctestServer) CreateTraderTws(ctx context.Context, req *grpc.CreateTraderTwsMapRequest) (*grpc.CreateTraderTwsMapResponse, error) {
	res, err := s.service.TraderTws().CreateTraderTws(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) GetTradersTws(ctx context.Context, req *grpc.GetTradersTwsMapRequest) (*grpc.GetTradersTwsMapResponse, error) {
	res, err := s.service.TraderTws().GetTradersTws(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) UpdateTraderTws(ctx context.Context, req *grpc.UpdateTraderTwsMapRequest) (*grpc.UpdateTraderTwsMapResponse, error) {
	res, err := s.service.TraderTws().UpdateTraderTws(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) DeleteTraderTws(ctx context.Context, req *grpc.DeleteTraderTwsMapRequest) (*grpc.DeleteTraderTwsMapResponse, error) {
	res, err := s.service.TraderTws().DeleteTraderTws(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

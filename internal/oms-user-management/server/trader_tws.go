// trader_tws
package server

import (
	"context"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
)

func (s *OmsUserManagementServer) CreateTraderTws(ctx context.Context, req *grpc.CreateTraderTwsMapRequest) (*grpc.CreateTraderTwsMapResponse, error) {
	res, err := s.service.TraderTws().CreateTraderTws(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetTradersTws(ctx context.Context, req *grpc.GetTradersTwsMapRequest) (*grpc.GetTradersTwsMapResponse, error) {
	res, err := s.service.TraderTws().GetTradersTws(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) UpdateTraderTws(ctx context.Context, req *grpc.UpdateTraderTwsMapRequest) (*grpc.UpdateTraderTwsMapResponse, error) {
	res, err := s.service.TraderTws().UpdateTraderTws(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) DeleteTraderTws(ctx context.Context, req *grpc.DeleteTraderTwsMapRequest) (*grpc.DeleteTraderTwsMapResponse, error) {
	res, err := s.service.TraderTws().DeleteTraderTws(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

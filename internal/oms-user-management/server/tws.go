package server

import (
	"context"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
)

func (s *OmsUserManagementServer) CreateTws(ctx context.Context, req *grpc.CreateTwsRequest) (*grpc.CreateTwsResponse, error) {
	res, err := s.service.TWS().CreateTws(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) UpdateTws(ctx context.Context, req *grpc.UpdateTwsRequest) (*grpc.UpdateTwsResponse, error) {
	res, err := s.service.TWS().UpdateTws(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetTws(ctx context.Context, req *grpc.GetTwsRequest) (*grpc.GetTwsResponse, error) {
	res, err := s.service.TWS().GetTws(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) DeleteTws(ctx context.Context, req *grpc.DeleteTwsRequest) (*grpc.DeleteTwsResponse, error) {
	res, err := s.service.TWS().DeleteTws(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

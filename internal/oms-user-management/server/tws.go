package server

import (
	"context"
	"gitlab.techetronventures.com/core/grpctest/pkg/grpc"
)

func (s *GrpctestServer) CreateTws(ctx context.Context, req *grpc.CreateTwsRequest) (*grpc.CreateTwsResponse, error) {
	res, err := s.service.TWS().CreateTws(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) UpdateTws(ctx context.Context, req *grpc.UpdateTwsRequest) (*grpc.UpdateTwsResponse, error) {
	res, err := s.service.TWS().UpdateTws(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) GetTws(ctx context.Context, req *grpc.GetTwsRequest) (*grpc.GetTwsResponse, error) {
	res, err := s.service.TWS().GetTws(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) DeleteTws(ctx context.Context, req *grpc.DeleteTwsRequest) (*grpc.DeleteTwsResponse, error) {
	res, err := s.service.TWS().DeleteTws(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

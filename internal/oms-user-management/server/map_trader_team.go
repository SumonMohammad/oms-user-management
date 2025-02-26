package server

import (
	"context"
	"gitlab.techetronventures.com/core/grpctest/pkg/grpc"
)

func (s *GrpctestServer) CreateMapTeam(ctx context.Context, req *grpc.CreateTraderMapRequest) (*grpc.CreateTraderMapResponse, error) {
	res, err := s.service.MapTraderTeam().CreateMapTeam(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) UpdateMapTeam(ctx context.Context, req *grpc.UpdateTraderMapRequest) (*grpc.UpdateTraderMapResponse, error) {
	res, err := s.service.MapTraderTeam().UpdateMapTeam(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) GetMappedTeams(ctx context.Context, req *grpc.GetTraderMapRequest) (*grpc.GetTraderMapResponse, error) {
	res, err := s.service.MapTraderTeam().GetMappedTeams(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) DeleteMappedTeam(ctx context.Context, req *grpc.DeleteTraderMapRequest) (*grpc.DeleteTraderMapResponse, error) {
	res, err := s.service.MapTraderTeam().DeleteMappedTeam(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

package server

import (
	"context"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
)

func (s *OmsUserManagementServer) CreateMapTeam(ctx context.Context, req *grpc.CreateTraderMapRequest) (*grpc.CreateTraderMapResponse, error) {
	res, err := s.service.MapTraderTeam().CreateMapTeam(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) UpdateMapTeam(ctx context.Context, req *grpc.UpdateTraderMapRequest) (*grpc.UpdateTraderMapResponse, error) {
	res, err := s.service.MapTraderTeam().UpdateMapTeam(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetMappedTeams(ctx context.Context, req *grpc.GetTraderMapRequest) (*grpc.GetTraderMapResponse, error) {
	res, err := s.service.MapTraderTeam().GetMappedTeams(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) DeleteMappedTeam(ctx context.Context, req *grpc.DeleteTraderMapRequest) (*grpc.DeleteTraderMapResponse, error) {
	res, err := s.service.MapTraderTeam().DeleteMappedTeam(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

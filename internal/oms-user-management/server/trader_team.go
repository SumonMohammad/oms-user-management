// trader_tws
package server

import (
	"context"
	"gitlab.techetronventures.com/core/grpctest/pkg/grpc"
)

func (s *GrpctestServer) CreateTeam(ctx context.Context, req *grpc.CreateTraderTeamRequest) (*grpc.CreateTraderTeamResponse, error) {
	res, err := s.service.TraderTeam().CreateTeam(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) GetTeams(ctx context.Context, req *grpc.GetTraderTeamRequest) (*grpc.GetTraderTeamResponse, error) {
	res, err := s.service.TraderTeam().GetTeams(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) UpdateTeam(ctx context.Context, req *grpc.UpdateTraderTeamRequest) (*grpc.UpdateTraderTeamResponse, error) {
	res, err := s.service.TraderTeam().UpdateTeam(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) DeleteTeam(ctx context.Context, req *grpc.DeleteTraderTeamRequest) (*grpc.DeleteTraderTeamResponse, error) {
	res, err := s.service.TraderTeam().DeleteTeam(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

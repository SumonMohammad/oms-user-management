package service

import (
	"context"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"go.uber.org/zap"
)

type TraderTeams interface {
	CreateTeam(ctx context.Context, req *grpc.CreateTraderTeamRequest) (*grpc.CreateTraderTeamResponse, error)
	GetTeams(ctx context.Context, req *grpc.GetTraderTeamRequest) (*grpc.GetTraderTeamResponse, error)
	UpdateTeam(ctx context.Context, req *grpc.UpdateTraderTeamRequest) (*grpc.UpdateTraderTeamResponse, error)
	DeleteTeam(ctx context.Context, req *grpc.DeleteTraderTeamRequest) (*grpc.DeleteTraderTeamResponse, error)
}
type TraderTeam struct {
	service *OmsUserManagementService
}

type TraderTeamReceiver struct {
	*OmsUserManagementService
}

func (ms *OmsUserManagementService) TraderTeam() TraderTeams {
	return &TraderTeamReceiver{
		ms,
	}
}

func (s *TraderTeamReceiver) CreateTeam(ctx context.Context, req *grpc.CreateTraderTeamRequest) (*grpc.CreateTraderTeamResponse, error) {
	// Map the incoming gRPC request to the Trader model

	traderTeam := &model.TraderTeam{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status.String(),
		IsEnabled:   req.IsEnabled,
		IsDeleted:   req.IsDeleted,
	}
	// Call the database method to create the trader
	err := s.db.TraderTeam().CreateTeam(ctx, traderTeam)
	if err != nil {
		msg := "failed to create trader_tws"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.CreateTraderTeamResponse{
		Code: 0,
	}, nil
}

func (s *TraderTeamReceiver) GetTeams(ctx context.Context, req *grpc.GetTraderTeamRequest) (*grpc.GetTraderTeamResponse, error) {

	res, count, err := s.db.TraderTeam().GetTeams(ctx, req)
	if err != nil {
		msg := "failed to fetch trader team"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	tradersTeam := []*grpc.GetTraderTeamResponse_TraderTeam{}
	for _, item := range res {

		var traderStatus grpc.Status
		if item.Status == "ACTIVE" {
			traderStatus = grpc.Status_ACTIVE
		} else {
			traderStatus = grpc.Status_PENDING
		}

		traderTeam := &grpc.GetTraderTeamResponse_TraderTeam{
			Id:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Status:      traderStatus,
			IsEnabled:   item.IsEnabled,
			IsDeleted:   item.IsDeleted,
		}
		tradersTeam = append(tradersTeam, traderTeam)
	}

	return &grpc.GetTraderTeamResponse{
		TradersTeam: tradersTeam,
		PaginationResponse: &grpc.PaginationInfoResponse{
			TotalRecordCount: int32(count),
		},
	}, nil

}

func (s *TraderTeamReceiver) UpdateTeam(ctx context.Context, req *grpc.UpdateTraderTeamRequest) (*grpc.UpdateTraderTeamResponse, error) {

	traderTeam := &model.TraderTeam{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status.String(),
		IsEnabled:   req.IsEnabled,
		IsDeleted:   req.IsDeleted,
	}

	// Call database update
	err := s.db.TraderTeam().UpdateTeam(ctx, traderTeam)
	if err != nil {
		msg := "failed to update trader team info"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return success response
	return &grpc.UpdateTraderTeamResponse{
		Code: 0,
	}, nil
}

func (s *TraderTeamReceiver) DeleteTeam(ctx context.Context, req *grpc.DeleteTraderTeamRequest) (*grpc.DeleteTraderTeamResponse, error) {
	// Map the incoming gRPC request to the TraderTws model
	traderTeam := &model.TraderTeam{
		ID: req.Id,
	}

	// Call the database method to delete the trader
	err := s.db.TraderTeam().DeleteTeam(ctx, traderTeam)
	if err != nil {
		msg := "failed to delete trader team"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.DeleteTraderTeamResponse{
		Code: 0,
	}, nil
}

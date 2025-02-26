package service

import (
	"context"
	//"fmt"
	model "gitlab.techetronventures.com/core/grpctest/internal/grpctest/models"
	//pg "gitlab.techetronventures.com/core/grpctest/internal/grpctest/models/pg"
	"gitlab.techetronventures.com/core/grpctest/pkg/grpc"
	"go.uber.org/zap"
)

type MapTraderTeams interface {
	CreateMapTeam(ctx context.Context, req *grpc.CreateTraderMapRequest) (*grpc.CreateTraderMapResponse, error)
	UpdateMapTeam(ctx context.Context, req *grpc.UpdateTraderMapRequest) (*grpc.UpdateTraderMapResponse, error)
	GetMappedTeams(ctx context.Context, req *grpc.GetTraderMapRequest) (*grpc.GetTraderMapResponse, error)
	DeleteMappedTeam(ctx context.Context, req *grpc.DeleteTraderMapRequest) (*grpc.DeleteTraderMapResponse, error)
}

type MapTraderTeam struct {
	service *GrpctestService
}

type MapTraderReceiver struct {
	*GrpctestService
}

func (ms *GrpctestService) MapTraderTeam() MapTraderTeams {
	return &MapTraderReceiver{
		ms,
	}
}

func (s *MapTraderReceiver) CreateMapTeam(ctx context.Context, req *grpc.CreateTraderMapRequest) (*grpc.CreateTraderMapResponse, error) {

	mapTraderTeam := &model.MapTraderTeam{
		TeamId:    req.TeamId,
		TraderId:  req.TraderId,
		IsEnabled: req.IsEnabled,
		Status:    req.Status.String(),
	}

	// Call the database method to create the trader
	err := s.db.MapTraderTeam().CreateMapTeam(ctx, mapTraderTeam)
	if err != nil {
		msg := "failed to create trader team map"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.CreateTraderMapResponse{
		Code: 0,
	}, nil
}

func (s *MapTraderReceiver) UpdateMapTeam(ctx context.Context, req *grpc.UpdateTraderMapRequest) (*grpc.UpdateTraderMapResponse, error) {

	mapTraderTeam := &model.MapTraderTeam{
		ID:        req.Id,
		TeamId:    req.TeamId,
		TraderId:  req.TraderId,
		IsEnabled: req.IsEnabled,
		Status:    req.Status.String(),
	}

	// Call the database method to create the trader
	err := s.db.MapTraderTeam().UpdateMapTeam(ctx, mapTraderTeam)
	if err != nil {
		msg := "failed to update trader team map"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.UpdateTraderMapResponse{
		Code: 0,
	}, nil
}
func (s *MapTraderReceiver) GetMappedTeams(ctx context.Context, req *grpc.GetTraderMapRequest) (*grpc.GetTraderMapResponse, error) {

	res, count, err := s.db.MapTraderTeam().GetMappedTeams(ctx, req)
	if err != nil {
		msg := "failed to get trader's team map info"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}
	mappedTraders := []*grpc.GetTraderMapResponse_TraderMap{}

	for _, item := range res {

		var status grpc.Status
		if item.Status == "ACTIVE" {
			status = grpc.Status_ACTIVE
		} else {
			status = grpc.Status_PENDING
		}

		mappedTrader := &grpc.GetTraderMapResponse_TraderMap{
			Id:        item.ID,
			TeamId:    item.TeamId,
			TraderId:  item.TraderId,
			IsEnabled: item.IsEnabled,
			Status:    status,
		}
		mappedTraders = append(mappedTraders, mappedTrader)
	}
	return &grpc.GetTraderMapResponse{
		MappedTrader: mappedTraders,
		PaginationResponse: &grpc.PaginationInfoResponse{
			TotalRecordCount: int32(count),
		},
	}, nil
}

func (s *MapTraderReceiver) DeleteMappedTeam(ctx context.Context, req *grpc.DeleteTraderMapRequest) (*grpc.DeleteTraderMapResponse, error) {
	mappedTrader := &model.MapTraderTeam{
		ID: req.Id,
	}

	// Call the database method to create the trader
	err := s.db.MapTraderTeam().DeleteMappedTeam(ctx, mappedTrader)
	if err != nil {
		msg := "failed to delete mapped trader team"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.DeleteTraderMapResponse{
		Code: 0,
	}, nil
}

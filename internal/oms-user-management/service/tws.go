package service

import (
	"context"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"go.uber.org/zap"
	"time"
)

type Tws interface {
	CreateTws(ctx context.Context, req *grpc.CreateTwsRequest) (*grpc.CreateTwsResponse, error)
	UpdateTws(ctx context.Context, req *grpc.UpdateTwsRequest) (*grpc.UpdateTwsResponse, error)
	GetTws(ctx context.Context, req *grpc.GetTwsRequest) (*grpc.GetTwsResponse, error)
	DeleteTws(ctx context.Context, req *grpc.DeleteTwsRequest) (*grpc.DeleteTwsResponse, error)
}

type TWS struct {
	service *OmsUserManagementService
}

type TwsReceiver struct {
	*OmsUserManagementService
}

func (ms *OmsUserManagementService) TWS() Tws {
	return &TwsReceiver{
		ms,
	}
}

func (s *TwsReceiver) CreateTws(ctx context.Context, req *grpc.CreateTwsRequest) (*grpc.CreateTwsResponse, error) {

	tws := &model.TWS{
		TwsCode:   req.TwsCode,
		IsActive:  req.IsActive,
		IsEnabled: req.IsEnabled,
		Status:    req.Status.String(),
		IsDeleted: req.IsDeleted,
	}

	// Call the database method to create the trader
	err := s.db.TWS().CreateTws(ctx, tws)
	if err != nil {
		msg := "failed to create tws"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.CreateTwsResponse{
		Code: 0,
	}, nil
}

func (s *TwsReceiver) UpdateTws(ctx context.Context, req *grpc.UpdateTwsRequest) (*grpc.UpdateTwsResponse, error) {

	tws := &model.TWS{
		ID:        req.Id,
		TwsCode:   req.TwsCode,
		IsActive:  req.IsActive,
		IsEnabled: req.IsEnabled,
		Status:    req.Status.String(),
		IsDeleted: req.IsDeleted,
		UpdatedAt: time.Now(),
	}

	// Call the database method to create the trader
	err := s.db.TWS().UpdateTws(ctx, tws)
	if err != nil {
		msg := "failed to update tws"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.UpdateTwsResponse{
		Code: 0,
	}, nil
}

func (s *TwsReceiver) GetTws(ctx context.Context, req *grpc.GetTwsRequest) (*grpc.GetTwsResponse, error) {

	res, count, err := s.db.TWS().GetTws(ctx, req)
	if err != nil {
		msg := "failed to fetch tws"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	twsList := []*grpc.GetTwsResponseTwsList{}
	for _, item := range res {
		var twsStatus grpc.TwsStatus

		if item.Status == "ASSIGNED" {
			twsStatus = grpc.TwsStatus_ASSIGNED
		} else {
			twsStatus = grpc.TwsStatus_UNASSIGNED
		}

		tws := &grpc.GetTwsResponseTwsList{
			Id:        item.ID,
			TwsCode:   item.TwsCode,
			IsActive:  item.IsActive,
			IsEnabled: item.IsEnabled,
			IsDeleted: item.IsDeleted,
			Status:    twsStatus,
		}
		twsList = append(twsList, tws)
	}

	return &grpc.GetTwsResponse{
		Tws: twsList,
		PaginationResponse: &grpc.PaginationInfoResponse{
			TotalRecordCount: int32(count),
		},
	}, nil

}

func (s *TwsReceiver) DeleteTws(ctx context.Context, req *grpc.DeleteTwsRequest) (*grpc.DeleteTwsResponse, error) {

	tws := &model.TWS{
		ID: req.Id,
	}

	// Call the database method to create the trader
	err := s.db.TWS().DeleteTws(ctx, tws)
	if err != nil {
		msg := "failed to delete tws"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.DeleteTwsResponse{
		Code: 0,
	}, nil
}

package service

import (
	"context"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BrokerHouses interface {
	CreateBrokerHouse(ctx context.Context, req *grpc.CreateBrokerHouseRequest) (*grpc.CreateBrokerHouseResponse, error)
	UpdateBrokerHouse(ctx context.Context, req *grpc.UpdateBrokerHouseRequest) (*grpc.UpdateBrokerHouseResponse, error)
	GetBrokerHouses(ctx context.Context, req *grpc.GetBrokerHousesRequest) (*grpc.GetBrokerHousesResponse, error)
	DeleteBrokerHouse(ctx context.Context, req *grpc.DeleteBrokerHouseRequest) (*grpc.DeleteBrokerHouseResponse, error)
	GetBrokerHouseById(ctx context.Context, req *grpc.GetBrokerHouseByIdRequest) (*grpc.GetBrokerHouseByIdResponse, error)
}

type BrokerHouse struct {
	service *OmsUserManagementService
}

type BrokerHouseReceiver struct {
	*OmsUserManagementService
}

func (ms *OmsUserManagementService) BrokerHouse() BrokerHouses {
	return &BrokerHouseReceiver{
		ms,
	}
}

func (s *BrokerHouseReceiver) CreateBrokerHouse(ctx context.Context, req *grpc.CreateBrokerHouseRequest) (*grpc.CreateBrokerHouseResponse, error) {

	brokerHouse := &model.BrokerHouse{
		BrokerHouseName: req.BrokerHouseName,
		ShortName:       req.ShortName,
		Description:     req.Description,
		Address:         req.Address,
		PhoneNumber:     req.PhoneNumber,
		TelephoneNumber: req.TelephoneNumber,
		EmailAddress:    req.EmailAddress,
		ValidCurrency:   req.ValidCurrency,
		Status:          req.Status,
		IsEnabled:       req.IsEnabled,
		CountryCode:     req.CountryCode,
	}

	// Call the database method to create the trader
	err := s.db.BrokerHouse().CreateBrokerHouse(ctx, brokerHouse)
	if err != nil {
		msg := "failed to create Broker House"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.CreateBrokerHouseResponse{
		Code: 0,
	}, nil
}

func (s *BrokerHouseReceiver) UpdateBrokerHouse(ctx context.Context, req *grpc.UpdateBrokerHouseRequest) (*grpc.UpdateBrokerHouseResponse, error) {

	brokerHouse := &model.BrokerHouse{
		ID:              req.Id,
		BrokerHouseName: req.BrokerHouseName,
		ShortName:       req.ShortName,
		Description:     req.Description,
		Address:         req.Address,
		PhoneNumber:     req.PhoneNumber,
		TelephoneNumber: req.TelephoneNumber,
		EmailAddress:    req.EmailAddress,
		ValidCurrency:   req.ValidCurrency,
		Status:          req.Status,
		IsEnabled:       req.IsEnabled,
		CountryCode:     req.CountryCode,
	}

	// Call the database method to create the trader
	err := s.db.BrokerHouse().UpdateBrokerHouse(ctx, brokerHouse)
	if err != nil {
		msg := "failed to update Broker House"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.UpdateBrokerHouseResponse{
		Code: 0,
	}, nil
}
func (s *BrokerHouseReceiver) GetBrokerHouses(ctx context.Context, req *grpc.GetBrokerHousesRequest) (*grpc.GetBrokerHousesResponse, error) {

	res, count, err := s.db.BrokerHouse().GetBrokerHouses(ctx, req)
	if err != nil {
		msg := "failed to get Broker House"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}
	brokerHouses := []*grpc.GetBrokerHousesResponseBrokerHouseList{}

	for _, item := range res {

		brokerHouse := &grpc.GetBrokerHousesResponseBrokerHouseList{
			Id:              item.ID,
			BrokerHouseName: item.BrokerHouseName,
			ShortName:       item.ShortName,
			Description:     item.Description,
			Address:         item.Address,
			PhoneNumber:     item.PhoneNumber,
			TelephoneNumber: item.TelephoneNumber,
			EmailAddress:    item.EmailAddress,
			ValidCurrency:   item.ValidCurrency,
			Status:          item.Status,
			IsEnabled:       item.IsEnabled,
			CountryCode:     item.CountryCode,
		}
		brokerHouses = append(brokerHouses, brokerHouse)
	}
	return &grpc.GetBrokerHousesResponse{
		BrokerHouses: brokerHouses,
		PaginationResponse: &grpc.PaginationInfoResponse{
			TotalRecordCount: int32(count),
		},
	}, nil
}

func (s *BrokerHouseReceiver) DeleteBrokerHouse(ctx context.Context, req *grpc.DeleteBrokerHouseRequest) (*grpc.DeleteBrokerHouseResponse, error) {
	brokerHouse := &model.BrokerHouse{
		ID: req.Id,
	}

	// Call the database method to create the trader
	err := s.db.BrokerHouse().DeleteBrokerHouse(ctx, brokerHouse)
	if err != nil {
		msg := "failed to delete Broker House"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.DeleteBrokerHouseResponse{
		Code: 0,
	}, nil
}

func (s *BrokerHouseReceiver) GetBrokerHouseById(ctx context.Context, req *grpc.GetBrokerHouseByIdRequest) (*grpc.GetBrokerHouseByIdResponse, error) {
	// Call the database layer to fetch the trader by ID or email
	brokerHouse, err := s.db.BrokerHouse().GetBrokerHouseById(ctx, req)
	if err != nil {
		// Handle error if no trader is found or any other error occurs
		s.log.Error(ctx, "Failed to fetch Broker House by ID", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to fetch Broker House")
	}

	// Map the result from the database layer to the GRPC response
	response := &grpc.GetBrokerHouseByIdResponse{
		Id:              brokerHouse.Id,
		BrokerHouseName: brokerHouse.BrokerHouseName,
		ShortName:       brokerHouse.ShortName,
		Description:     brokerHouse.Description,
		Address:         brokerHouse.Address,
		PhoneNumber:     brokerHouse.PhoneNumber,
		TelephoneNumber: brokerHouse.TelephoneNumber,
		EmailAddress:    brokerHouse.EmailAddress,
		ValidCurrency:   brokerHouse.ValidCurrency,
		Status:          brokerHouse.Status,
		IsEnabled:       brokerHouse.IsEnabled,
		CountryCode:     brokerHouse.CountryCode,
	}

	return response, nil
}

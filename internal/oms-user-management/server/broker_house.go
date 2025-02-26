package server

import (
	"context"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
)

func (s *OmsUserManagementServer) CreateBrokerHouse(ctx context.Context, req *grpc.CreateBrokerHouseRequest) (*grpc.CreateBrokerHouseResponse, error) {
	res, err := s.service.BrokerHouse().CreateBrokerHouse(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) UpdateBrokerHouse(ctx context.Context, req *grpc.UpdateBrokerHouseRequest) (*grpc.UpdateBrokerHouseResponse, error) {
	res, err := s.service.BrokerHouse().UpdateBrokerHouse(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetBrokerHouses(ctx context.Context, req *grpc.GetBrokerHousesRequest) (*grpc.GetBrokerHousesResponse, error) {
	res, err := s.service.BrokerHouse().GetBrokerHouses(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) DeleteBrokerHouse(ctx context.Context, req *grpc.DeleteBrokerHouseRequest) (*grpc.DeleteBrokerHouseResponse, error) {
	res, err := s.service.BrokerHouse().DeleteBrokerHouse(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetBrokerHouseById(ctx context.Context, req *grpc.GetBrokerHouseByIdRequest) (*grpc.GetBrokerHouseByIdResponse, error) {
	res, err := s.service.BrokerHouse().GetBrokerHouseById(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

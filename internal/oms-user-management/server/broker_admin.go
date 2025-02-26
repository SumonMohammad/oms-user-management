package server

import (
	"context"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
)

func (s *OmsUserManagementServer) CreateBrokerAdmin(ctx context.Context, req *grpc.CreateBrokerAdminRequest) (*grpc.CreateBrokerAdminResponse, error) {
	res, err := s.service.BrokerAdmin().CreateBrokerAdmin(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) UpdateBrokerAdmin(ctx context.Context, req *grpc.UpdateBrokerAdminRequest) (*grpc.UpdateBrokerAdminResponse, error) {
	res, err := s.service.BrokerAdmin().UpdateBrokerAdmin(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetBrokerAdminById(ctx context.Context, req *grpc.GetBrokerAdminByIdRequest) (*grpc.GetBrokerAdminByIdResponse, error) {
	res, err := s.service.BrokerAdmin().GetBrokerAdminById(ctx, req.UserId)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

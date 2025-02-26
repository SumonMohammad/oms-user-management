package server

import (
	"context"
	"gitlab.techetronventures.com/core/grpctest/pkg/grpc"
)

func (s *GrpctestServer) CreateBrokerAdmin(ctx context.Context, req *grpc.CreateBrokerAdminRequest) (*grpc.CreateBrokerAdminResponse, error) {
	res, err := s.service.BrokerAdmin().CreateBrokerAdmin(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) UpdateBrokerAdmin(ctx context.Context, req *grpc.UpdateBrokerAdminRequest) (*grpc.UpdateBrokerAdminResponse, error) {
	res, err := s.service.BrokerAdmin().UpdateBrokerAdmin(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) GetBrokerAdmins(ctx context.Context, req *grpc.GetBrokerAdminsRequest) (*grpc.GetBrokerAdminsResponse, error) {
	res, err := s.service.BrokerAdmin().GetBrokerAdmins(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) DeleteBrokerAdmin(ctx context.Context, req *grpc.DeleteBrokerAdminRequest) (*grpc.DeleteBrokerAdminResponse, error) {
	res, err := s.service.BrokerAdmin().DeleteBrokerAdmin(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) GetBrokerAdminByIdOrUserName(ctx context.Context, req *grpc.GetBrokerAdminByIdOrUserNameRequest) (*grpc.GetBrokerAdminByIdOrUserNameResponse, error) {
	res, err := s.service.BrokerAdmin().GetBrokerAdminByIdOrUserName(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

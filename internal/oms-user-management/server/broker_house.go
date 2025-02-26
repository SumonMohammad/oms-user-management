package server

import (
	"context"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
)

func (s *OmsUserManagementServer) CreateOffice(ctx context.Context, req *grpc.CreateOfficeRequest) (*grpc.CreateOfficeResponse, error) {
	res, err := s.service.Office().CreateOffice(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) UpdateOffice(ctx context.Context, req *grpc.UpdateOfficeRequest) (*grpc.UpdateOfficeResponse, error) {
	res, err := s.service.Office().UpdateOffice(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetOffices(ctx context.Context, req *grpc.GetOfficesRequest) (*grpc.GetOfficesResponse, error) {
	res, err := s.service.Office().GetOffices(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) DeleteOffice(ctx context.Context, req *grpc.DeleteOfficeRequest) (*grpc.DeleteOfficeResponse, error) {
	res, err := s.service.Office().DeleteOffice(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetOfficeByIdOrEmail(ctx context.Context, req *grpc.GetOfficeByIdOrEmailRequest) (*grpc.GetOfficeByIdOrEmailResponse, error) {
	res, err := s.service.Office().GetOfficeByIdOrEmail(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

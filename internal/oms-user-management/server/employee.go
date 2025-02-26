package server

import (
	"context"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
)

func (s *OmsUserManagementServer) CreateEmployee(ctx context.Context, req *grpc.CreateEmployeeRequest) (*grpc.CreateEmployeeResponse, error) {
	res, err := s.service.Employee().CreateEmployee(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) UpdateEmployee(ctx context.Context, req *grpc.UpdateEmployeeRequest) (*grpc.UpdateEmployeeResponse, error) {
	res, err := s.service.Employee().UpdateEmployee(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetEmployeeById(ctx context.Context, req *grpc.GetEmployeeByIdRequest) (*grpc.GetEmployeeByIdResponse, error) {
	res, err := s.service.Employee().GetEmployeeById(ctx, req.UserId)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

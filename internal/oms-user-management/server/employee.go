package server

import (
	"context"
	"gitlab.techetronventures.com/core/grpctest/pkg/grpc"
)

func (s *GrpctestServer) CreateEmployee(ctx context.Context, req *grpc.CreateEmployeeRequest) (*grpc.CreateEmployeeResponse, error) {
	res, err := s.service.Employee().CreateEmployee(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) UpdateEmployee(ctx context.Context, req *grpc.UpdateEmployeeRequest) (*grpc.UpdateEmployeeResponse, error) {
	res, err := s.service.Employee().UpdateEmployee(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) GetEmployees(ctx context.Context, req *grpc.GetEmployeesRequest) (*grpc.GetEmployeesResponse, error) {
	res, err := s.service.Employee().GetEmployees(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) DeleteEmployee(ctx context.Context, req *grpc.DeleteEmployeeRequest) (*grpc.DeleteEmployeeResponse, error) {
	res, err := s.service.Employee().DeleteEmployee(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) GetEmployeeByIdOrEmail(ctx context.Context, req *grpc.GetEmployeeByIdOrEmailRequest) (*grpc.GetEmployeeByIdOrEmailResponse, error) {
	res, err := s.service.Employee().GetEmployeeByIdOrEmail(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

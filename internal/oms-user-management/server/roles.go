package server

import (
	"context"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
)

func (s *OmsUserManagementServer) CreateRole(ctx context.Context, req *grpc.CreateRoleRequest) (*grpc.CreateRoleResponse, error) {
	res, err := s.service.Role().CreateRole(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) UpdateRole(ctx context.Context, req *grpc.UpdateRoleRequest) (*grpc.UpdateRoleResponse, error) {
	res, err := s.service.Role().UpdateRole(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetRoles(ctx context.Context, req *grpc.GetRolesRequest) (*grpc.GetRolesResponse, error) {
	res, err := s.service.Role().GetRoles(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) DeleteRole(ctx context.Context, req *grpc.DeleteRoleRequest) (*grpc.DeleteRoleResponse, error) {
	res, err := s.service.Role().DeleteRole(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetRoleById(ctx context.Context, req *grpc.GetRoleByIdRequest) (*grpc.GetRoleByIdResponse, error) {
	res, err := s.service.Role().GetRoleById(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

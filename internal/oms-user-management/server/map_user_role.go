package server

import (
	"context"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
)

func (s *OmsUserManagementServer) CreateMapUserRole(ctx context.Context, req *grpc.CreateMapUserRoleRequest) (*grpc.CreateMapUserRoleResponse, error) {
	res, err := s.service.MapUserRole().CreateMapUserRole(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) UpdateMapUserRole(ctx context.Context, req *grpc.UpdateMapUserRoleRequest) (*grpc.UpdateMapUserRoleResponse, error) {
	res, err := s.service.MapUserRole().UpdateMapUserRole(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetMapUserRoles(ctx context.Context, req *grpc.GetMapUserRolesRequest) (*grpc.GetMapUserRolesResponse, error) {
	res, err := s.service.MapUserRole().GetMapUserRoles(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) DeleteMapUserRole(ctx context.Context, req *grpc.DeleteMapUserRoleRequest) (*grpc.DeleteMapUserRoleResponse, error) {
	res, err := s.service.MapUserRole().DeleteMapUserRole(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetMapUserRoleById(ctx context.Context, req *grpc.GetMapUserRoleByIdRequest) (*grpc.GetMapUserRoleByIdResponse, error) {
	res, err := s.service.MapUserRole().GetMapUserRoleById(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

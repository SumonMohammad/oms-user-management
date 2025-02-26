package server

import (
	"context"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
)

func (s *OmsUserManagementServer) CreateMapRolePermission(ctx context.Context, req *grpc.CreateMapRolePermissionRequest) (*grpc.CreateMapRolePermissionResponse, error) {
	res, err := s.service.MapRolePermission().CreateMapRolePermission(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) UpdateMapRolePermission(ctx context.Context, req *grpc.UpdateMapRolePermissionRequest) (*grpc.UpdateMapRolePermissionResponse, error) {
	res, err := s.service.MapRolePermission().UpdateMapRolePermission(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetMapRolePermissions(ctx context.Context, req *grpc.GetMapRolePermissionsRequest) (*grpc.GetMapRolePermissionsResponse, error) {
	res, err := s.service.MapRolePermission().GetMapRolePermissions(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) DeleteMapRolePermission(ctx context.Context, req *grpc.DeleteMapRolePermissionRequest) (*grpc.DeleteMapRolePermissionResponse, error) {
	res, err := s.service.MapRolePermission().DeleteMapRolePermission(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetMapRolePermissionById(ctx context.Context, req *grpc.GetMapRolePermissionByIdRequest) (*grpc.GetMapRolePermissionByIdResponse, error) {
	res, err := s.service.MapRolePermission().GetMapRolePermissionById(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

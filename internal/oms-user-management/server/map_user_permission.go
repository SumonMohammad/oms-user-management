package server

import (
	"context"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
)

func (s *OmsUserManagementServer) CreateMapGroupPermission(ctx context.Context, req *grpc.CreateMapGroupPermissionRequest) (*grpc.CreateMapGroupPermissionResponse, error) {
	res, err := s.service.MapGroupPermission().CreateMapGroupPermission(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) UpdateMapGroupPermission(ctx context.Context, req *grpc.UpdateMapGroupPermissionRequest) (*grpc.UpdateMapGroupPermissionResponse, error) {
	res, err := s.service.MapGroupPermission().UpdateMapGroupPermission(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetMapGroupPermissions(ctx context.Context, req *grpc.GetMapGroupPermissionsRequest) (*grpc.GetMapGroupPermissionsResponse, error) {
	res, err := s.service.MapGroupPermission().GetMapGroupPermissions(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) DeleteMapGroupPermission(ctx context.Context, req *grpc.DeleteMapGroupPermissionRequest) (*grpc.DeleteMapGroupPermissionResponse, error) {
	res, err := s.service.MapGroupPermission().DeleteMapGroupPermission(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetMapGroupPermissionById(ctx context.Context, req *grpc.GetMapGroupPermissionByIdRequest) (*grpc.GetMapGroupPermissionByIdResponse, error) {
	res, err := s.service.MapGroupPermission().GetMapGroupPermissionById(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

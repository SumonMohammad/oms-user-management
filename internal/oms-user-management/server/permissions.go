package server

import (
	"context"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
)

func (s *OmsUserManagementServer) CreatePermission(ctx context.Context, req *grpc.CreatePermissionRequest) (*grpc.CreatePermissionResponse, error) {
	res, err := s.service.Permission().CreatePermission(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) UpdatePermission(ctx context.Context, req *grpc.UpdatePermissionRequest) (*grpc.UpdatePermissionResponse, error) {
	res, err := s.service.Permission().UpdatePermission(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetPermissions(ctx context.Context, req *grpc.GetPermissionsRequest) (*grpc.GetPermissionsResponse, error) {
	res, err := s.service.Permission().GetPermissions(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) DeletePermission(ctx context.Context, req *grpc.DeletePermissionRequest) (*grpc.DeletePermissionResponse, error) {
	res, err := s.service.Permission().DeletePermission(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetPermissionById(ctx context.Context, req *grpc.GetPermissionByIdRequest) (*grpc.GetPermissionByIdResponse, error) {
	res, err := s.service.Permission().GetPermissionById(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

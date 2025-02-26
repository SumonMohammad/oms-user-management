package server

import (
	"context"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
)

func (s *OmsUserManagementServer) CreateGroupPermission(ctx context.Context, req *grpc.CreateGroupPermissionRequest) (*grpc.CreateGroupPermissionResponse, error) {
	res, err := s.service.GroupPermission().CreateGroupPermission(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) UpdateGroupPermission(ctx context.Context, req *grpc.UpdateGroupPermissionRequest) (*grpc.UpdateGroupPermissionResponse, error) {
	res, err := s.service.GroupPermission().UpdateGroupPermission(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetGroupPermissions(ctx context.Context, req *grpc.GetGroupPermissionsRequest) (*grpc.GetGroupPermissionsResponse, error) {
	res, err := s.service.GroupPermission().GetGroupPermissions(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) DeleteGroupPermission(ctx context.Context, req *grpc.DeleteGroupPermissionRequest) (*grpc.DeleteGroupPermissionResponse, error) {
	res, err := s.service.GroupPermission().DeleteGroupPermission(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetGroupPermissionById(ctx context.Context, req *grpc.GetGroupPermissionByIdRequest) (*grpc.GetGroupPermissionByIdResponse, error) {
	res, err := s.service.GroupPermission().GetGroupPermissionById(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

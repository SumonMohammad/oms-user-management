package server

import (
	"context"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
)

func (s *OmsUserManagementServer) CreateMapUserPermission(ctx context.Context, req *grpc.CreateMapUserPermissionRequest) (*grpc.CreateMapUserPermissionResponse, error) {
	res, err := s.service.MapUserPermission().CreateMapUserPermission(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) UpdateMapUserPermission(ctx context.Context, req *grpc.UpdateMapUserPermissionRequest) (*grpc.UpdateMapUserPermissionResponse, error) {
	res, err := s.service.MapUserPermission().UpdateMapUserPermission(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) DeleteMapUserPermission(ctx context.Context, req *grpc.DeleteMapUserPermissionRequest) (*grpc.DeleteMapUserPermissionResponse, error) {
	res, err := s.service.MapUserPermission().DeleteMapUserPermission(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetUserPermissionsByUserId(ctx context.Context, req *grpc.GetUserPermissionsByUserIdRequest) (*grpc.GetUserPermissionsByUserIdResponse, error) {
	res, err := s.service.MapUserPermission().GetUserPermissionsByUserId(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

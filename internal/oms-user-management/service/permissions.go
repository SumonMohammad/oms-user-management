package service

import (
	"context"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Permissions interface {
	CreatePermission(ctx context.Context, req *grpc.CreatePermissionRequest) (*grpc.CreatePermissionResponse, error)
	UpdatePermission(ctx context.Context, req *grpc.UpdatePermissionRequest) (*grpc.UpdatePermissionResponse, error)
	GetPermissions(ctx context.Context, req *grpc.GetPermissionsRequest) (*grpc.GetPermissionsResponse, error)
	DeletePermission(ctx context.Context, req *grpc.DeletePermissionRequest) (*grpc.DeletePermissionResponse, error)
	GetPermissionById(ctx context.Context, req *grpc.GetPermissionByIdRequest) (*grpc.GetPermissionByIdResponse, error)
}

type Permission struct {
	service *OmsUserManagementService
}

type PermissionReceiver struct {
	*OmsUserManagementService
}

func (ms *OmsUserManagementService) Permission() Permissions {
	return &PermissionReceiver{
		ms,
	}
}

func (s *PermissionReceiver) CreatePermission(ctx context.Context, req *grpc.CreatePermissionRequest) (*grpc.CreatePermissionResponse, error) {

	permission := &model.Permission{
		Name:        req.Name,
		Description: req.Description,
		IsEnabled:   req.IsEnabled,
	}

	// Call the database method to create the trader
	err := s.db.Permission().CreatePermission(ctx, permission)
	if err != nil {
		msg := "failed to create permission"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.CreatePermissionResponse{
		Code: 0,
	}, nil
}

func (s *PermissionReceiver) UpdatePermission(ctx context.Context, req *grpc.UpdatePermissionRequest) (*grpc.UpdatePermissionResponse, error) {

	permission := &model.Permission{
		Name:        req.Name,
		Description: req.Description,
		IsEnabled:   req.IsEnabled,
	}
	// Call the database method to create the trader
	err := s.db.Permission().UpdatePermission(ctx, permission)
	if err != nil {
		msg := "failed to update permission"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.UpdatePermissionResponse{
		Code: 0,
	}, nil
}
func (s *PermissionReceiver) GetPermissions(ctx context.Context, req *grpc.GetPermissionsRequest) (*grpc.GetPermissionsResponse, error) {

	res, count, err := s.db.Permission().GetPermissions(ctx, req)
	if err != nil {
		msg := "failed to get Permission"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}
	permissions := []*grpc.GetPermissionsResponse_PermissionList{}

	for _, item := range res {

		permission := &grpc.GetPermissionsResponse_PermissionList{
			Id:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			IsEnabled:   item.IsEnabled,
		}
		permissions = append(permissions, permission)
	}
	return &grpc.GetPermissionsResponse{
		Permissions: permissions,
		PaginationResponse: &grpc.PaginationInfoResponse{
			TotalRecordCount: int32(count),
		},
	}, nil
}

func (s *PermissionReceiver) DeletePermission(ctx context.Context, req *grpc.DeletePermissionRequest) (*grpc.DeletePermissionResponse, error) {
	permission := &model.Permission{
		ID: req.Id,
	}

	// Call the database method to create the trader
	err := s.db.Permission().DeletePermission(ctx, permission)
	if err != nil {
		msg := "failed to delete permission"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.DeletePermissionResponse{
		Code: 0,
	}, nil
}

func (s *PermissionReceiver) GetPermissionById(ctx context.Context, req *grpc.GetPermissionByIdRequest) (*grpc.GetPermissionByIdResponse, error) {

	permission, err := s.db.Permission().GetPermissionById(ctx, req)
	if err != nil {
		s.log.Error(ctx, "Failed to fetch permission", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to fetch permission")
	}

	// Map the raw data to the gRPC response format
	response := &grpc.GetPermissionByIdResponse{
		Id:          permission.ID,
		Name:        permission.Name,
		Description: permission.Description,
		IsEnabled:   permission.IsEnabled,
	}

	return response, nil
}

// Map the result to the response format

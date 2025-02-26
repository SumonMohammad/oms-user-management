package service

import (
	"context"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MapRolePermissions interface {
	CreateMapRolePermission(ctx context.Context, req *grpc.CreateMapRolePermissionRequest) (*grpc.CreateMapRolePermissionResponse, error)
	UpdateMapRolePermission(ctx context.Context, req *grpc.UpdateMapRolePermissionRequest) (*grpc.UpdateMapRolePermissionResponse, error)
	GetMapRolePermissions(ctx context.Context, req *grpc.GetMapRolePermissionsRequest) (*grpc.GetMapRolePermissionsResponse, error)
	DeleteMapRolePermission(ctx context.Context, req *grpc.DeleteMapRolePermissionRequest) (*grpc.DeleteMapRolePermissionResponse, error)
	GetMapRolePermissionById(ctx context.Context, req *grpc.GetMapRolePermissionByIdRequest) (*grpc.GetMapRolePermissionByIdResponse, error)
}

type MapRolePermission struct {
	service *OmsUserManagementService
}

type MapRolePermissionReceiver struct {
	*OmsUserManagementService
}

func (ms *OmsUserManagementService) MapRolePermission() MapRolePermissions {
	return &MapRolePermissionReceiver{
		ms,
	}
}

func (s *MapRolePermissionReceiver) CreateMapRolePermission(ctx context.Context, req *grpc.CreateMapRolePermissionRequest) (*grpc.CreateMapRolePermissionResponse, error) {

	rolePermission := &model.MapRolePermission{
		ReferenceType: req.UserType.String(),
		ReferenceID:   req.ReferenceId,
		PermissionID:  req.PermissionId,
		IsEnabled:     req.IsEnabled,
	}

	// Call the database method to create the trader
	err := s.db.MapRolePermission().CreateMapRolePermission(ctx, rolePermission)
	if err != nil {
		msg := "failed to create map role Permission"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.CreateMapRolePermissionResponse{
		Code: 0,
	}, nil
}

func (s *MapRolePermissionReceiver) UpdateMapRolePermission(ctx context.Context, req *grpc.UpdateMapRolePermissionRequest) (*grpc.UpdateMapRolePermissionResponse, error) {

	rolePermission := &model.MapRolePermission{
		ReferenceType: req.ReferenceType.String(),
		ReferenceID:   req.ReferenceId,
		PermissionID:  req.PermissionId,
		IsEnabled:     req.IsEnabled,
		ID:            req.Id,
	}
	// Call the database method to create the trader
	err := s.db.MapRolePermission().UpdateMapRolePermission(ctx, rolePermission)
	if err != nil {
		msg := "failed to update map role Permission"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.UpdateMapRolePermissionResponse{
		Code: 0,
	}, nil
}
func (s *MapRolePermissionReceiver) GetMapRolePermissions(ctx context.Context, req *grpc.GetMapRolePermissionsRequest) (*grpc.GetMapRolePermissionsResponse, error) {

	res, count, err := s.db.MapRolePermission().GetMapRolePermissions(ctx, req)
	if err != nil {
		msg := "failed to get map role Permissions"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}
	rolePermissions := []*grpc.GetMapRolePermissionsResponseMappedRoleList{}

	for _, item := range res {

		rolePermission := &grpc.GetMapRolePermissionsResponseMappedRoleList{

			ReferenceId:  item.ReferenceID,
			PermissionId: item.PermissionID,
			IsEnabled:    item.IsEnabled,
		}
		rolePermissions = append(rolePermissions, rolePermission)
	}
	return &grpc.GetMapRolePermissionsResponse{
		MappedPermissionRoles: rolePermissions,
		PaginationResponse: &grpc.PaginationInfoResponse{
			TotalRecordCount: int32(count),
		},
	}, nil
}

func (s *MapRolePermissionReceiver) DeleteMapRolePermission(ctx context.Context, req *grpc.DeleteMapRolePermissionRequest) (*grpc.DeleteMapRolePermissionResponse, error) {
	rolePermission := &model.MapRolePermission{
		ID: req.Id,
	}

	// Call the database method to create the trader
	err := s.db.MapRolePermission().DeleteMapRolePermission(ctx, rolePermission)
	if err != nil {
		msg := "failed to delete map role Permission"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.DeleteMapRolePermissionResponse{
		Code: 0,
	}, nil
}

func (s *MapRolePermissionReceiver) GetMapRolePermissionById(ctx context.Context, req *grpc.GetMapRolePermissionByIdRequest) (*grpc.GetMapRolePermissionByIdResponse, error) {

	rolePermission, err := s.db.MapRolePermission().GetMapRolePermissionById(ctx, req)
	if err != nil {
		s.log.Error(ctx, "Failed to fetch map role Permission", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to fetch role Permission")
	}
	referenceType := grpc.UserType_TRADER
	if rolePermission.ReferenceType == "TRADER" {
		referenceType = grpc.UserType_TRADER
	} else if rolePermission.ReferenceType == "INVESTOR" {
		referenceType = grpc.UserType_INVESTOR
	} else if rolePermission.ReferenceType == "BROKER_ADMIN" {
		referenceType = grpc.UserType_BROKER_ADMIN
	} else if rolePermission.ReferenceType == "USER_TYPE_UNSPECIFIED" {
		referenceType = grpc.UserType_USER_TYPE_UNSPECIFIED
	} else {
		return nil, status.Error(codes.InvalidArgument, "Invalid role Permission")
	}

	// Map the raw data to the gRPC response format
	response := &grpc.GetMapRolePermissionByIdResponse{
		UserType:     referenceType,
		ReferenceId:  rolePermission.ReferenceID,
		PermissionId: rolePermission.PermissionID,
		IsEnabled:    rolePermission.IsEnabled,
		Id:           rolePermission.ID,
	}

	return response, nil
}

// Map the result to the response format

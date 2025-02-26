package service

import (
	"context"
	"encoding/json"
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
		RoleID:       req.RoleId,
		PermissionID: req.PermissionId,
		IsEnabled:    req.IsEnabled,
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
		RoleID:       req.RoleId,
		PermissionID: req.PermissionId,
		IsEnabled:    req.IsEnabled,
		ID:           req.Id,
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
			RoleId:       item.RoleID,
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
	// Call service to get permissions
	res, err := s.db.MapRolePermission().GetMapRolePermissionById(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, status.Errorf(codes.Internal, "failed to fetch permissions: %v", err)
	}

	// Transform response into expected format
	var rolesAndPermissions []*grpc.RolesAndPermissions
	for _, perm := range res {
		rolesAndPermissions = append(rolesAndPermissions, &grpc.RolesAndPermissions{
			RoleId:       perm.RoleID,
			PermissionId: perm.PermissionID,
			Name:         perm.Name,
			IsEnabled:    perm.IsEnabled,
		})
	}

	// Create response object
	response := &grpc.GetMapRolePermissionByIdResponse{
		RolesAndPermissions: rolesAndPermissions,
		Code:                0,
	}
	// Log JSON representation (optional)
	jsonData, err := json.Marshal(response)
	if err == nil {
		s.log.Info(ctx, "Response JSON: "+string(jsonData))
	} else {
		s.log.Error(ctx, "Failed to serialize response to JSON: "+err.Error())
	}

	return response, nil

}

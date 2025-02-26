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

type MapUserPermissions interface {
	CreateMapUserPermission(ctx context.Context, req *grpc.CreateMapUserPermissionRequest) (*grpc.CreateMapUserPermissionResponse, error)
	UpdateMapUserPermission(ctx context.Context, req *grpc.UpdateMapUserPermissionRequest) (*grpc.UpdateMapUserPermissionResponse, error)

	DeleteMapUserPermission(ctx context.Context, req *grpc.DeleteMapUserPermissionRequest) (*grpc.DeleteMapUserPermissionResponse, error)
	GetUserPermissionsByUserId(ctx context.Context, req *grpc.GetUserPermissionsByUserIdRequest) (*grpc.GetUserPermissionsByUserIdResponse, error)
}

type MapUserPermission struct {
	service *OmsUserManagementService
}

type MapUserPermissionReceiver struct {
	*OmsUserManagementService
}

func (ms *OmsUserManagementService) MapUserPermission() MapUserPermissions {
	return &MapUserPermissionReceiver{
		ms,
	}
}

func (s *MapUserPermissionReceiver) CreateMapUserPermission(ctx context.Context, req *grpc.CreateMapUserPermissionRequest) (*grpc.CreateMapUserPermissionResponse, error) {

	mapPermission := &model.MapUserPermission{
		UserID:       req.UserId,
		PermissionID: req.PermissionId,
		IsEnabled:    req.IsEnabled,
		IsRevoked:    req.IsRevoked,
	}

	// Call the database method to create the trader
	err := s.db.MapUserPermission().CreateMapUserPermission(ctx, mapPermission)
	if err != nil {
		msg := "failed to create Map User Permission"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.CreateMapUserPermissionResponse{
		Code: 0,
	}, nil
}

func (s *MapUserPermissionReceiver) UpdateMapUserPermission(ctx context.Context, req *grpc.UpdateMapUserPermissionRequest) (*grpc.UpdateMapUserPermissionResponse, error) {

	mapPermission := &model.MapUserPermission{
		UserID:       req.UserId,
		PermissionID: req.PermissionId,
		IsEnabled:    req.IsEnabled,
		IsRevoked:    req.IsRevoked,
	}
	// Call the database method to create the trader
	err := s.db.MapUserPermission().UpdateMapUserPermission(ctx, mapPermission)
	if err != nil {
		msg := "failed to update map user permission"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.UpdateMapUserPermissionResponse{
		Code: 0,
	}, nil
}

func (s *MapUserPermissionReceiver) DeleteMapUserPermission(ctx context.Context, req *grpc.DeleteMapUserPermissionRequest) (*grpc.DeleteMapUserPermissionResponse, error) {
	mapPermission := &model.MapUserPermission{
		UserID:       req.UserId,
		PermissionID: req.PermissionId,
	}

	// Call the database method to create the trader
	err := s.db.MapUserPermission().DeleteMapUserPermission(ctx, mapPermission)
	if err != nil {
		msg := "failed to delete map Permission"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.DeleteMapUserPermissionResponse{
		Code: 0,
	}, nil
}

func (s *MapUserPermissionReceiver) GetUserPermissionsByUserId(ctx context.Context, req *grpc.GetUserPermissionsByUserIdRequest) (*grpc.GetUserPermissionsByUserIdResponse, error) {
	// Call service to get permissions
	res, err := s.db.MapUserPermission().GetUserPermissionsByUserId(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, status.Errorf(codes.Internal, "failed to fetch permissions: %v", err)
	}

	// Transform response into expected format
	var rolesAndPermissions []*grpc.GetUserPermissionsByUserIdResponse_UserPermissions
	for _, perm := range res {
		rolesAndPermissions = append(rolesAndPermissions, &grpc.GetUserPermissionsByUserIdResponse_UserPermissions{
			PermissionId: perm.PermissionID,
			Name:         perm.Name,
		})
	}

	// Create response object
	response := &grpc.GetUserPermissionsByUserIdResponse{
		Code:            0,
		UserPermissions: rolesAndPermissions,
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

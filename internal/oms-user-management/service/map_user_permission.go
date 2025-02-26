package service

import (
	"context"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MapGroupPermissions interface {
	CreateMapGroupPermission(ctx context.Context, req *grpc.CreateMapGroupPermissionRequest) (*grpc.CreateMapGroupPermissionResponse, error)
	UpdateMapGroupPermission(ctx context.Context, req *grpc.UpdateMapGroupPermissionRequest) (*grpc.UpdateMapGroupPermissionResponse, error)
	GetMapGroupPermissions(ctx context.Context, req *grpc.GetMapGroupPermissionsRequest) (*grpc.GetMapGroupPermissionsResponse, error)
	DeleteMapGroupPermission(ctx context.Context, req *grpc.DeleteMapGroupPermissionRequest) (*grpc.DeleteMapGroupPermissionResponse, error)
	GetMapGroupPermissionById(ctx context.Context, req *grpc.GetMapGroupPermissionByIdRequest) (*grpc.GetMapGroupPermissionByIdResponse, error)
}

type MapGroupPermission struct {
	service *OmsUserManagementService
}

type MapGroupPermissionReceiver struct {
	*OmsUserManagementService
}

func (ms *OmsUserManagementService) MapGroupPermission() MapGroupPermissions {
	return &MapGroupPermissionReceiver{
		ms,
	}
}

func (s *MapGroupPermissionReceiver) CreateMapGroupPermission(ctx context.Context, req *grpc.CreateMapGroupPermissionRequest) (*grpc.CreateMapGroupPermissionResponse, error) {

	mapPermission := &model.MapGroupPermission{
		GroupID:      req.GroupId,
		PermissionID: req.PermissionId,
		IsEnabled:    req.IsEnabled,
		Status:       req.Status.String(),
	}

	// Call the database method to create the trader
	err := s.db.MapGroupPermission().CreateMapGroupPermission(ctx, mapPermission)
	if err != nil {
		msg := "failed to create group Permission"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.CreateMapGroupPermissionResponse{
		Code: 0,
	}, nil
}

func (s *MapGroupPermissionReceiver) UpdateMapGroupPermission(ctx context.Context, req *grpc.UpdateMapGroupPermissionRequest) (*grpc.UpdateMapGroupPermissionResponse, error) {

	mapPermission := &model.MapGroupPermission{
		GroupID:      req.GroupId,
		PermissionID: req.PermissionId,
		IsEnabled:    req.IsEnabled,
		Status:       req.Status.String(),
	}
	// Call the database method to create the trader
	err := s.db.MapGroupPermission().UpdateMapGroupPermission(ctx, mapPermission)
	if err != nil {
		msg := "failed to update group Permission"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.UpdateMapGroupPermissionResponse{
		Code: 0,
	}, nil
}
func (s *MapGroupPermissionReceiver) GetMapGroupPermissions(ctx context.Context, req *grpc.GetMapGroupPermissionsRequest) (*grpc.GetMapGroupPermissionsResponse, error) {

	res, count, err := s.db.MapGroupPermission().GetMapGroupPermissions(ctx, req)
	if err != nil {
		msg := "failed to get map group Permissions"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}
	mapPermissions := []*grpc.GetMapGroupPermissionsResponseMappedGroupList{}

	for _, item := range res {

		mapPermission := &grpc.GetMapGroupPermissionsResponseMappedGroupList{
			Id:           item.ID,
			PermissionId: item.PermissionID,
			GroupId:      item.GroupID,
			IsEnabled:    item.IsEnabled,
		}
		mapPermissions = append(mapPermissions, mapPermission)
	}
	return &grpc.GetMapGroupPermissionsResponse{
		MappedPermissionGroups: mapPermissions,
		PaginationResponse: &grpc.PaginationInfoResponse{
			TotalRecordCount: int32(count),
		},
	}, nil
}

func (s *MapGroupPermissionReceiver) DeleteMapGroupPermission(ctx context.Context, req *grpc.DeleteMapGroupPermissionRequest) (*grpc.DeleteMapGroupPermissionResponse, error) {
	mapPermission := &model.MapGroupPermission{
		ID: req.Id,
	}

	// Call the database method to create the trader
	err := s.db.MapGroupPermission().DeleteMapGroupPermission(ctx, mapPermission)
	if err != nil {
		msg := "failed to delete map Permission"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.DeleteMapGroupPermissionResponse{
		Code: 0,
	}, nil
}

func (s *MapGroupPermissionReceiver) GetMapGroupPermissionById(ctx context.Context, req *grpc.GetMapGroupPermissionByIdRequest) (*grpc.GetMapGroupPermissionByIdResponse, error) {

	mapPermission, err := s.db.MapGroupPermission().GetMapGroupPermissionById(ctx, req)
	if err != nil {
		s.log.Error(ctx, "Failed to fetch map group Permission", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to fetch group Permission")
	}
	statusOfPermission := grpc.Status_PENDING
	if mapPermission.Status == "ACTIVE" {
		statusOfPermission = grpc.Status_ACTIVE
	}

	// Map the raw data to the gRPC response format
	response := &grpc.GetMapGroupPermissionByIdResponse{
		Id:           mapPermission.ID,
		GroupId:      mapPermission.GroupID,
		PermissionId: mapPermission.PermissionID,
		Status:       statusOfPermission,
		IsEnabled:    mapPermission.IsEnabled,
	}

	return response, nil
}

// Map the result to the response format

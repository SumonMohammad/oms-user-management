package service

import (
	"context"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GroupPermissions interface {
	CreateGroupPermission(ctx context.Context, req *grpc.CreateGroupPermissionRequest) (*grpc.CreateGroupPermissionResponse, error)
	UpdateGroupPermission(ctx context.Context, req *grpc.UpdateGroupPermissionRequest) (*grpc.UpdateGroupPermissionResponse, error)
	GetGroupPermissions(ctx context.Context, req *grpc.GetGroupPermissionsRequest) (*grpc.GetGroupPermissionsResponse, error)
	DeleteGroupPermission(ctx context.Context, req *grpc.DeleteGroupPermissionRequest) (*grpc.DeleteGroupPermissionResponse, error)
	GetGroupPermissionById(ctx context.Context, req *grpc.GetGroupPermissionByIdRequest) (*grpc.GetGroupPermissionByIdResponse, error)
}

type GroupPermission struct {
	service *OmsUserManagementService
}

type GroupPermissionReceiver struct {
	*OmsUserManagementService
}

func (ms *OmsUserManagementService) GroupPermission() GroupPermissions {
	return &GroupPermissionReceiver{
		ms,
	}
}

func (s *GroupPermissionReceiver) CreateGroupPermission(ctx context.Context, req *grpc.CreateGroupPermissionRequest) (*grpc.CreateGroupPermissionResponse, error) {

	groupPermission := &model.GroupPermission{
		Name:        req.Name,
		Description: req.Description,
		IsDeleted:   req.IsDeleted,
		IsEnabled:   req.IsEnabled,
		Status:      req.Status.String(),
	}

	// Call the database method to create the trader
	err := s.db.GroupPermission().CreateGroupPermission(ctx, groupPermission)
	if err != nil {
		msg := "failed to create group Permission"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.CreateGroupPermissionResponse{
		Code: 0,
	}, nil
}

func (s *GroupPermissionReceiver) UpdateGroupPermission(ctx context.Context, req *grpc.UpdateGroupPermissionRequest) (*grpc.UpdateGroupPermissionResponse, error) {

	groupPermission := &model.GroupPermission{
		Name:        req.Name,
		Description: req.Description,
		IsDeleted:   req.IsDeleted,
		IsEnabled:   req.IsEnabled,
		Status:      req.Status.String(),
	}
	// Call the database method to create the trader
	err := s.db.GroupPermission().UpdateGroupPermission(ctx, groupPermission)
	if err != nil {
		msg := "failed to update group Permission"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.UpdateGroupPermissionResponse{
		Code: 0,
	}, nil
}
func (s *GroupPermissionReceiver) GetGroupPermissions(ctx context.Context, req *grpc.GetGroupPermissionsRequest) (*grpc.GetGroupPermissionsResponse, error) {

	res, count, err := s.db.GroupPermission().GetGroupPermissions(ctx, req)
	if err != nil {
		msg := "failed to get group Permissions"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}
	groupPermissions := []*grpc.GetGroupPermissionsResponsePermissionGroupList{}

	for _, item := range res {

		groupPermission := &grpc.GetGroupPermissionsResponsePermissionGroupList{
			Id:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			IsDeleted:   item.IsDeleted,
			IsEnabled:   item.IsEnabled,
		}
		groupPermissions = append(groupPermissions, groupPermission)
	}
	return &grpc.GetGroupPermissionsResponse{
		PermissionGroups: groupPermissions,
		PaginationResponse: &grpc.PaginationInfoResponse{
			TotalRecordCount: int32(count),
		},
	}, nil
}

func (s *GroupPermissionReceiver) DeleteGroupPermission(ctx context.Context, req *grpc.DeleteGroupPermissionRequest) (*grpc.DeleteGroupPermissionResponse, error) {
	groupPermission := &model.GroupPermission{
		ID: req.Id,
	}

	// Call the database method to create the trader
	err := s.db.GroupPermission().DeleteGroupPermission(ctx, groupPermission)
	if err != nil {
		msg := "failed to delete Group Permission"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.DeleteGroupPermissionResponse{
		Code: 0,
	}, nil
}

func (s *GroupPermissionReceiver) GetGroupPermissionById(ctx context.Context, req *grpc.GetGroupPermissionByIdRequest) (*grpc.GetGroupPermissionByIdResponse, error) {

	groupPermission, err := s.db.GroupPermission().GetGroupPermissionById(ctx, req)
	if err != nil {
		s.log.Error(ctx, "Failed to fetch Employee", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to fetch Employee")
	}
	statusOfPermission := grpc.Status_PENDING
	if groupPermission.Status == "ACTIVE" {
		statusOfPermission = grpc.Status_ACTIVE
	}

	// Map the raw data to the gRPC response format
	response := &grpc.GetGroupPermissionByIdResponse{
		Id:          groupPermission.ID,
		Name:        groupPermission.Name,
		Description: groupPermission.Description,
		Status:      statusOfPermission,
		IsDeleted:   groupPermission.IsDeleted,
		IsEnabled:   groupPermission.IsEnabled,
	}

	return response, nil
}

// Map the result to the response format

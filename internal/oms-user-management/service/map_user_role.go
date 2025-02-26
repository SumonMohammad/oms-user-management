package service

import (
	"context"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MapUserRoles interface {
	CreateMapUserRole(ctx context.Context, req *grpc.CreateMapUserRoleRequest) (*grpc.CreateMapUserRoleResponse, error)
	UpdateMapUserRole(ctx context.Context, req *grpc.UpdateMapUserRoleRequest) (*grpc.UpdateMapUserRoleResponse, error)
	GetMapUserRoles(ctx context.Context, req *grpc.GetMapUserRolesRequest) (*grpc.GetMapUserRolesResponse, error)
	DeleteMapUserRole(ctx context.Context, req *grpc.DeleteMapUserRoleRequest) (*grpc.DeleteMapUserRoleResponse, error)
	GetMapUserRoleById(ctx context.Context, req *grpc.GetMapUserRoleByIdRequest) (*grpc.GetMapUserRoleByIdResponse, error)
}

type MapUserRole struct {
	service *OmsUserManagementService
}

type MapUserRoleReceiver struct {
	*OmsUserManagementService
}

func (ms *OmsUserManagementService) MapUserRole() MapUserRoles {
	return &MapUserRoleReceiver{
		ms,
	}
}

func (s *MapUserRoleReceiver) CreateMapUserRole(ctx context.Context, req *grpc.CreateMapUserRoleRequest) (*grpc.CreateMapUserRoleResponse, error) {

	mapUserRole := &model.MapUserRole{
		UserId:    req.UserId,
		RoleId:    req.RoleId,
		IsEnabled: req.IsEnabled,
	}

	// Call the database method to create the trader
	err := s.db.MapUserRole().CreateMapUserRole(ctx, mapUserRole)
	if err != nil {
		msg := "failed to create mapUserRole"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.CreateMapUserRoleResponse{
		Code: 0,
	}, nil
}

func (s *MapUserRoleReceiver) UpdateMapUserRole(ctx context.Context, req *grpc.UpdateMapUserRoleRequest) (*grpc.UpdateMapUserRoleResponse, error) {

	mapUserRole := &model.MapUserRole{
		ID:        req.Id,
		UserId:    req.UserId,
		RoleId:    req.RoleId,
		IsEnabled: req.IsEnabled,
	}
	// Call the database method to create the trader
	err := s.db.MapUserRole().UpdateMapUserRole(ctx, mapUserRole)
	if err != nil {
		msg := "failed to update mapUserRole"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.UpdateMapUserRoleResponse{
		Code: 0,
	}, nil
}
func (s *MapUserRoleReceiver) GetMapUserRoles(ctx context.Context, req *grpc.GetMapUserRolesRequest) (*grpc.GetMapUserRolesResponse, error) {

	res, count, err := s.db.MapUserRole().GetMapUserRoles(ctx, req)
	if err != nil {
		msg := "failed to get MapUserRole"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}
	mapUserRoles := []*grpc.GetMapUserRolesResponseMappedUserRoleList{}

	for _, item := range res {

		mapUserRole := &grpc.GetMapUserRolesResponseMappedUserRoleList{
			Id:        item.ID,
			UserId:    item.UserId,
			RoleId:    item.RoleId,
			IsEnabled: item.IsEnabled,
		}
		mapUserRoles = append(mapUserRoles, mapUserRole)
	}
	return &grpc.GetMapUserRolesResponse{
		MappedUserRoleLists: mapUserRoles,
		PaginationResponse: &grpc.PaginationInfoResponse{
			TotalRecordCount: int32(count),
		},
	}, nil
}

func (s *MapUserRoleReceiver) DeleteMapUserRole(ctx context.Context, req *grpc.DeleteMapUserRoleRequest) (*grpc.DeleteMapUserRoleResponse, error) {
	mapUserRole := &model.MapUserRole{
		ID: req.Id,
	}

	// Call the database method to create the trader
	err := s.db.MapUserRole().DeleteMapUserRole(ctx, mapUserRole)
	if err != nil {
		msg := "failed to delete mapUserRole"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.DeleteMapUserRoleResponse{
		Code: 0,
	}, nil
}

func (s *MapUserRoleReceiver) GetMapUserRoleById(ctx context.Context, req *grpc.GetMapUserRoleByIdRequest) (*grpc.GetMapUserRoleByIdResponse, error) {

	mapUserRole, err := s.db.MapUserRole().GetMapUserRoleById(ctx, req)
	if err != nil {
		s.log.Error(ctx, "Failed to fetch Employee", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to fetch Employee")
	}

	// Map the raw data to the gRPC response format
	response := &grpc.GetMapUserRoleByIdResponse{
		Id:        mapUserRole.ID,
		UserId:    mapUserRole.UserId,
		RoleId:    mapUserRole.RoleId,
		IsEnabled: mapUserRole.IsEnabled,
	}

	return response, nil
}

// Map the result to the response format

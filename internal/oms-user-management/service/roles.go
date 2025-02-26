package service

import (
	"context"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Roles interface {
	CreateRole(ctx context.Context, req *grpc.CreateRoleRequest) (*grpc.CreateRoleResponse, error)
	UpdateRole(ctx context.Context, req *grpc.UpdateRoleRequest) (*grpc.UpdateRoleResponse, error)
	GetRoles(ctx context.Context, req *grpc.GetRolesRequest) (*grpc.GetRolesResponse, error)
	DeleteRole(ctx context.Context, req *grpc.DeleteRoleRequest) (*grpc.DeleteRoleResponse, error)
	GetRoleById(ctx context.Context, req *grpc.GetRoleByIdRequest) (*grpc.GetRoleByIdResponse, error)
}

type Role struct {
	service *OmsUserManagementService
}

type RoleReceiver struct {
	*OmsUserManagementService
}

func (ms *OmsUserManagementService) Role() Roles {
	return &RoleReceiver{
		ms,
	}
}

func (s *RoleReceiver) CreateRole(ctx context.Context, req *grpc.CreateRoleRequest) (*grpc.CreateRoleResponse, error) {

	role := &model.Role{
		RoleName:    req.RoleName,
		Description: req.Description,
		IsEnabled:   req.IsEnabled,
	}

	// Call the database method to create the trader
	err := s.db.Role().CreateRole(ctx, role)
	if err != nil {
		msg := "failed to create Role"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.CreateRoleResponse{
		Code: 0,
	}, nil
}

func (s *RoleReceiver) UpdateRole(ctx context.Context, req *grpc.UpdateRoleRequest) (*grpc.UpdateRoleResponse, error) {

	role := &model.Role{
		RoleName:    req.RoleName,
		Description: req.Description,
		IsEnabled:   req.IsEnabled,
	}
	// Call the database method to create the trader
	err := s.db.Role().UpdateRole(ctx, role)
	if err != nil {
		msg := "failed to update Role"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.UpdateRoleResponse{
		Code: 0,
	}, nil
}
func (s *RoleReceiver) GetRoles(ctx context.Context, req *grpc.GetRolesRequest) (*grpc.GetRolesResponse, error) {

	res, count, err := s.db.Role().GetRoles(ctx, req)
	if err != nil {
		msg := "failed to get Role"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}
	roles := []*grpc.GetRolesResponse_RoleList{}

	for _, item := range res {

		role := &grpc.GetRolesResponse_RoleList{
			Id:          item.ID,
			RoleName:    item.RoleName,
			Description: item.Description,
			IsEnabled:   item.IsEnabled,
		}
		roles = append(roles, role)
	}
	return &grpc.GetRolesResponse{
		Roles: roles,
		PaginationResponse: &grpc.PaginationInfoResponse{
			TotalRecordCount: int32(count),
		},
	}, nil
}

func (s *RoleReceiver) DeleteRole(ctx context.Context, req *grpc.DeleteRoleRequest) (*grpc.DeleteRoleResponse, error) {
	role := &model.Role{
		ID: req.Id,
	}

	// Call the database method to create the trader
	err := s.db.Role().DeleteRole(ctx, role)
	if err != nil {
		msg := "failed to delete Role"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.DeleteRoleResponse{
		Code: 0,
	}, nil
}

func (s *RoleReceiver) GetRoleById(ctx context.Context, req *grpc.GetRoleByIdRequest) (*grpc.GetRoleByIdResponse, error) {

	role, err := s.db.Role().GetRoleById(ctx, req)
	if err != nil {
		s.log.Error(ctx, "Failed to fetch Role", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to fetch Role")
	}

	// Map the raw data to the gRPC response format
	response := &grpc.GetRoleByIdResponse{
		Id:          role.ID,
		RoleName:    role.RoleName,
		Description: role.Description,
		IsEnabled:   role.IsEnabled,
	}

	return response, nil
}

// Map the result to the response format

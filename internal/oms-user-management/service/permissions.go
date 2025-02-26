package service

import (
	"context"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Privileges interface {
	CreatePrivilege(ctx context.Context, req *grpc.CreatePrivilegeRequest) (*grpc.CreatePrivilegeResponse, error)
	UpdatePrivilege(ctx context.Context, req *grpc.UpdatePrivilegeRequest) (*grpc.UpdatePrivilegeResponse, error)
	GetPrivileges(ctx context.Context, req *grpc.GetPrivilegesRequest) (*grpc.GetPrivilegesResponse, error)
	DeletePrivilege(ctx context.Context, req *grpc.DeletePrivilegeRequest) (*grpc.DeletePrivilegeResponse, error)
	GetPrivilegeById(ctx context.Context, req *grpc.GetPrivilegeByIdRequest) (*grpc.GetPrivilegeByIdResponse, error)
}

type Privilege struct {
	service *OmsUserManagementService
}

type PrivilegeReceiver struct {
	*OmsUserManagementService
}

func (ms *OmsUserManagementService) Privilege() Privileges {
	return &PrivilegeReceiver{
		ms,
	}
}

func (s *PrivilegeReceiver) CreatePrivilege(ctx context.Context, req *grpc.CreatePrivilegeRequest) (*grpc.CreatePrivilegeResponse, error) {

	privilege := &model.Privilege{
		Name:               req.Name,
		PrivilegeType:      req.PrivilegeType.String(),
		PrivilegeShortName: req.PrivilegeShortName,
		Description:        req.Description,
		IsDeleted:          req.IsDeleted,
		IsEnabled:          req.IsEnabled,
		Status:             req.Status.String(),
	}

	// Call the database method to create the trader
	err := s.db.Privilege().CreatePrivilege(ctx, privilege)
	if err != nil {
		msg := "failed to create employee"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.CreatePrivilegeResponse{
		Code: 0,
	}, nil
}

func (s *PrivilegeReceiver) UpdatePrivilege(ctx context.Context, req *grpc.UpdatePrivilegeRequest) (*grpc.UpdatePrivilegeResponse, error) {

	privilege := &model.Privilege{
		Name:               req.Name,
		PrivilegeType:      req.PrivilegeType.String(),
		PrivilegeShortName: req.PrivilegeShortName,
		Description:        req.Description,
		IsDeleted:          req.IsDeleted,
		IsEnabled:          req.IsEnabled,
		Status:             req.Status.String(),
	}
	// Call the database method to create the trader
	err := s.db.Privilege().UpdatePrivilege(ctx, privilege)
	if err != nil {
		msg := "failed to update oms user"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.UpdatePrivilegeResponse{
		Code: 0,
	}, nil
}
func (s *PrivilegeReceiver) GetPrivileges(ctx context.Context, req *grpc.GetPrivilegesRequest) (*grpc.GetPrivilegesResponse, error) {

	res, count, err := s.db.Privilege().GetPrivileges(ctx, req)
	if err != nil {
		msg := "failed to get privilege"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}
	privileges := []*grpc.GetPrivilegesResponsePrivilegeList{}

	for _, item := range res {

		privilege := &grpc.GetPrivilegesResponsePrivilegeList{
			Id:        item.ID,
			IsDeleted: item.IsDeleted,
			IsEnabled: item.IsEnabled,
		}
		privileges = append(privileges, privilege)
	}
	return &grpc.GetPrivilegesResponse{
		Privileges: privileges,
		PaginationResponse: &grpc.PaginationInfoResponse{
			TotalRecordCount: int32(count),
		},
	}, nil
}

func (s *PrivilegeReceiver) DeletePrivilege(ctx context.Context, req *grpc.DeletePrivilegeRequest) (*grpc.DeletePrivilegeResponse, error) {
	privilege := &model.Privilege{
		ID: req.Id,
	}

	// Call the database method to create the trader
	err := s.db.Privilege().DeletePrivilege(ctx, privilege)
	if err != nil {
		msg := "failed to delete employee"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.DeletePrivilegeResponse{
		Code: 0,
	}, nil
}

func (s *PrivilegeReceiver) GetPrivilegeById(ctx context.Context, req *grpc.GetPrivilegeByIdRequest) (*grpc.GetPrivilegeByIdResponse, error) {

	privilege, err := s.db.Privilege().GetPrivilegeById(ctx, req)
	if err != nil {
		s.log.Error(ctx, "Failed to fetch Employee", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to fetch Employee")
	}
	statusOfPrivilege := grpc.Status_PENDING
	if privilege.Status == "ACTIVE" {
		statusOfPrivilege = grpc.Status_ACTIVE
	}

	privilegeType := grpc.PrivilegeType_UNSPECIFIED_PRIVILEGE
	// Map the raw data to the gRPC response format
	response := &grpc.GetPrivilegeByIdResponse{
		Id:                 privilege.ID,
		Name:               privilege.Name,
		PrivilegeShortName: privilege.PrivilegeShortName,
		Description:        privilege.Description,
		PrivilegeType:      privilegeType,
		Status:             statusOfPrivilege,
		IsDeleted:          privilege.IsDeleted,
		IsEnabled:          privilege.IsEnabled,
	}

	return response, nil
}

// Map the result to the response format

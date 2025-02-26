package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	//"fmt"
	model "gitlab.techetronventures.com/core/grpctest/internal/grpctest/models"
	//pg "gitlab.techetronventures.com/core/grpctest/internal/grpctest/models/pg"
	"gitlab.techetronventures.com/core/grpctest/pkg/grpc"
	"go.uber.org/zap"
)

type BrokerAdmins interface {
	CreateBrokerAdmin(ctx context.Context, req *grpc.CreateBrokerAdminRequest) (*grpc.CreateBrokerAdminResponse, error)
	UpdateBrokerAdmin(ctx context.Context, req *grpc.UpdateBrokerAdminRequest) (*grpc.UpdateBrokerAdminResponse, error)
	GetBrokerAdmins(ctx context.Context, req *grpc.GetBrokerAdminsRequest) (*grpc.GetBrokerAdminsResponse, error)
	DeleteBrokerAdmin(ctx context.Context, req *grpc.DeleteBrokerAdminRequest) (*grpc.DeleteBrokerAdminResponse, error)
	GetBrokerAdminByIdOrUserName(ctx context.Context, req *grpc.GetBrokerAdminByIdOrUserNameRequest) (*grpc.GetBrokerAdminByIdOrUserNameResponse, error)
}

type BrokerAdmin struct {
	service *GrpctestService
}

type BrokerAdminReceiver struct {
	*GrpctestService
}

func (ms *GrpctestService) BrokerAdmin() BrokerAdmins {
	return &BrokerAdminReceiver{
		ms,
	}
}

func (s *BrokerAdminReceiver) CreateBrokerAdmin(ctx context.Context, req *grpc.CreateBrokerAdminRequest) (*grpc.CreateBrokerAdminResponse, error) {

	brokerAdmin := &model.BrokerAdmin{
		EmployeeId:     req.EmployeeId,
		UserName:       req.UserName,
		PasswordHash:   req.PasswordHash,
		ManagerId:      req.ManagerId,
		OfficeId:       req.OfficeId,
		CanLogin:       req.CanLogin,
		CanTrade:       req.CanTrade,
		ReadOnly:       req.ReadOnly,
		Type:           req.Type,
		IsIsolatedUser: req.IsIsolatedUser,
		IsDeleted:      req.IsDeleted,
		Status:         req.Status,
	}

	// Call the database method to create the trader
	err := s.db.BrokerAdmin().CreateBrokerAdmin(ctx, brokerAdmin)
	if err != nil {
		msg := "failed to create broker admin"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.CreateBrokerAdminResponse{
		Code: 0,
	}, nil
}

func (s *BrokerAdminReceiver) UpdateBrokerAdmin(ctx context.Context, req *grpc.UpdateBrokerAdminRequest) (*grpc.UpdateBrokerAdminResponse, error) {

	brokerAdmin := &model.BrokerAdmin{
		ID:             req.Id,
		EmployeeId:     req.EmployeeId,
		UserName:       req.UserName,
		PasswordHash:   req.PasswordHash,
		ManagerId:      req.ManagerId,
		OfficeId:       req.OfficeId,
		CanLogin:       req.CanLogin,
		CanTrade:       req.CanTrade,
		ReadOnly:       req.ReadOnly,
		Type:           req.Type,
		IsIsolatedUser: req.IsIsolatedUser,
		IsDeleted:      req.IsDeleted,
		Status:         req.Status,
	}

	// Call the database method to create the trader
	err := s.db.BrokerAdmin().UpdateBrokerAdmin(ctx, brokerAdmin)
	if err != nil {
		msg := "failed to update broker admin"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.UpdateBrokerAdminResponse{
		Code: 0,
	}, nil
}
func (s *BrokerAdminReceiver) GetBrokerAdmins(ctx context.Context, req *grpc.GetBrokerAdminsRequest) (*grpc.GetBrokerAdminsResponse, error) {

	res, count, err := s.db.BrokerAdmin().GetBrokerAdmins(ctx, req)
	if err != nil {
		msg := "failed to get broker admins"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}
	brokerAdmins := []*grpc.GetBrokerAdminsResponseBrokerAdminList{}

	for _, item := range res {

		brokerAdmin := &grpc.GetBrokerAdminsResponseBrokerAdminList{
			Id:             item.ID,
			EmployeeId:     item.EmployeeId,
			UserName:       item.UserName,
			PasswordHash:   item.PasswordHash,
			ManagerId:      item.ManagerId,
			OfficeId:       item.OfficeId,
			CanLogin:       item.CanLogin,
			CanTrade:       item.CanTrade,
			ReadOnly:       item.ReadOnly,
			Type:           item.Type,
			IsIsolatedUser: item.IsIsolatedUser,
			IsDeleted:      item.IsDeleted,
			Status:         item.Status,
		}
		brokerAdmins = append(brokerAdmins, brokerAdmin)
	}
	return &grpc.GetBrokerAdminsResponse{
		BrokerAdmins: brokerAdmins,
		PaginationResponse: &grpc.PaginationInfoResponse{
			TotalRecordCount: int32(count),
		},
	}, nil
}

func (s *BrokerAdminReceiver) DeleteBrokerAdmin(ctx context.Context, req *grpc.DeleteBrokerAdminRequest) (*grpc.DeleteBrokerAdminResponse, error) {
	brokerAdmin := &model.BrokerAdmin{
		ID: req.Id,
	}

	// Call the database method to create the trader
	err := s.db.BrokerAdmin().DeleteBrokerAdmin(ctx, brokerAdmin)
	if err != nil {
		msg := "failed to delete broker admin"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.DeleteBrokerAdminResponse{
		Code: 0,
	}, nil
}

func (s *BrokerAdminReceiver) GetBrokerAdminByIdOrUserName(ctx context.Context, req *grpc.GetBrokerAdminByIdOrUserNameRequest) (*grpc.GetBrokerAdminByIdOrUserNameResponse, error) {
	// Call the database layer to fetch the trader by ID or email
	brokerAdmin, err := s.db.BrokerAdmin().GetBrokerAdminByIdOrUserName(ctx, req)
	if err != nil {
		// Handle error if no trader is found or any other error occurs
		s.log.Error(ctx, "Failed to fetch Broker Admin by ID or User Name", zap.Error(err))
		if err.Error() == "Broker Admin not found" {
			return nil, status.Error(codes.NotFound, "Broker Admin not found")
		}
		return nil, status.Error(codes.Internal, "Failed to fetch Broker Admin")
	}

	// Map the result from the database layer to the GRPC response
	response := &grpc.GetBrokerAdminByIdOrUserNameResponse{
		Id:             brokerAdmin.Id,
		EmployeeId:     brokerAdmin.EmployeeId,
		UserName:       brokerAdmin.UserName,
		PasswordHash:   brokerAdmin.PasswordHash,
		ManagerId:      brokerAdmin.ManagerId,
		OfficeId:       brokerAdmin.OfficeId,
		CanLogin:       brokerAdmin.CanLogin,
		CanTrade:       brokerAdmin.CanTrade,
		ReadOnly:       brokerAdmin.ReadOnly,
		Type:           brokerAdmin.Type,
		IsIsolatedUser: brokerAdmin.IsIsolatedUser,
		IsDeleted:      brokerAdmin.IsDeleted,
		Status:         brokerAdmin.Status,
	}

	return response, nil
}

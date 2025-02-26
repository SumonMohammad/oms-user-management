package service

import (
	"context"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Branches interface {
	CreateBranch(ctx context.Context, req *grpc.CreateBranchRequest) (*grpc.CreateBranchResponse, error)
	UpdateBranch(ctx context.Context, req *grpc.UpdateBranchRequest) (*grpc.UpdateBranchResponse, error)
	GetBranches(ctx context.Context, req *grpc.GetBranchesRequest) (*grpc.GetBranchesResponse, error)
	DeleteBranch(ctx context.Context, req *grpc.DeleteBranchRequest) (*grpc.DeleteBranchResponse, error)
	GetBranchById(ctx context.Context, req *grpc.GetBranchByIdRequest) (*grpc.GetBranchByIdResponse, error)
}

type Branch struct {
	service *OmsUserManagementService
}

type BranchReceiver struct {
	*OmsUserManagementService
}

func (ms *OmsUserManagementService) Branch() Branches {
	return &BranchReceiver{
		ms,
	}
}

func (s *BranchReceiver) CreateBranch(ctx context.Context, req *grpc.CreateBranchRequest) (*grpc.CreateBranchResponse, error) {

	branch := &model.Branch{
		BranchName:      req.BranchName,
		BranchType:      req.BranchType,
		BrokerHouseId:   req.BrokerHouseId,
		ShortName:       req.ShortName,
		Description:     req.Description,
		Address:         req.Address,
		PhoneNumber:     req.PhoneNumber,
		TelephoneNumber: req.TelephoneNumber,
		EmailAddress:    req.EmailAddress,
		ValidCurrency:   req.ValidCurrency,
		Status:          req.Status,
		IsEnabled:       req.IsEnabled,
		CountryCode:     req.CountryCode,
		IsActive:        req.IsActive,
	}

	// Call the database method to create the trader
	err := s.db.Branch().CreateBranch(ctx, branch)
	if err != nil {
		msg := "failed to create Branch"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.CreateBranchResponse{
		Code: 0,
	}, nil
}

func (s *BranchReceiver) UpdateBranch(ctx context.Context, req *grpc.UpdateBranchRequest) (*grpc.UpdateBranchResponse, error) {

	branch := &model.Branch{
		ID:              req.Id,
		BranchName:      req.BranchName,
		BranchType:      req.BranchType,
		BrokerHouseId:   req.BrokerHouseId,
		ShortName:       req.ShortName,
		Description:     req.Description,
		Address:         req.Address,
		PhoneNumber:     req.PhoneNumber,
		TelephoneNumber: req.TelephoneNumber,
		EmailAddress:    req.EmailAddress,
		ValidCurrency:   req.ValidCurrency,
		Status:          req.Status,
		IsEnabled:       req.IsEnabled,
		CountryCode:     req.CountryCode,
		IsActive:        req.IsActive,
	}

	// Call the database method to create the trader
	err := s.db.Branch().UpdateBranch(ctx, branch)
	if err != nil {
		msg := "failed to update Branch"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.UpdateBranchResponse{
		Code: 0,
	}, nil
}
func (s *BranchReceiver) GetBranches(ctx context.Context, req *grpc.GetBranchesRequest) (*grpc.GetBranchesResponse, error) {

	res, count, err := s.db.Branch().GetBranches(ctx, req)
	if err != nil {
		msg := "failed to get Branches"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}
	branches := []*grpc.GetBranchesResponseBranchList{}

	for _, item := range res {

		branch := &grpc.GetBranchesResponseBranchList{
			Id:              item.ID,
			BranchName:      item.BranchName,
			BranchType:      item.BranchType,
			BrokerHouseId:   item.BrokerHouseId,
			ShortName:       item.ShortName,
			Description:     item.Description,
			Address:         item.Address,
			PhoneNumber:     item.PhoneNumber,
			TelephoneNumber: item.TelephoneNumber,
			EmailAddress:    item.EmailAddress,
			ValidCurrency:   item.ValidCurrency,
			Status:          item.Status,
			IsEnabled:       item.IsEnabled,
			CountryCode:     item.CountryCode,
			IsActive:        item.IsActive,
		}
		branches = append(branches, branch)
	}
	return &grpc.GetBranchesResponse{
		Branches: branches,
		PaginationResponse: &grpc.PaginationInfoResponse{
			TotalRecordCount: int32(count),
		},
	}, nil
}

func (s *BranchReceiver) DeleteBranch(ctx context.Context, req *grpc.DeleteBranchRequest) (*grpc.DeleteBranchResponse, error) {
	branch := &model.Branch{
		ID: req.Id,
	}

	// Call the database method to create the trader
	err := s.db.Branch().DeleteBranch(ctx, branch)
	if err != nil {
		msg := "failed to delete branch"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.DeleteBranchResponse{
		Code: 0,
	}, nil
}

func (s *BranchReceiver) GetBranchById(ctx context.Context, req *grpc.GetBranchByIdRequest) (*grpc.GetBranchByIdResponse, error) {
	// Call the database layer to fetch the trader by ID or email
	branch, err := s.db.Branch().GetBranchById(ctx, req)
	if err != nil {
		// Handle error if no trader is found or any other error occurs
		s.log.Error(ctx, "Failed to fetch Branch by ID", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to fetch Branch")
	}

	// Map the result from the database layer to the GRPC response
	response := &grpc.GetBranchByIdResponse{
		Id:              branch.Id,
		BranchName:      branch.BranchName,
		BranchType:      branch.BranchType,
		BrokerHouseId:   branch.BrokerHouseId,
		ShortName:       branch.ShortName,
		Description:     branch.Description,
		Address:         branch.Address,
		PhoneNumber:     branch.PhoneNumber,
		TelephoneNumber: branch.TelephoneNumber,
		EmailAddress:    branch.EmailAddress,
		ValidCurrency:   branch.ValidCurrency,
		Status:          branch.Status,
		IsEnabled:       branch.IsEnabled,
		CountryCode:     branch.CountryCode,
		IsActive:        branch.IsActive,
	}

	return response, nil
}

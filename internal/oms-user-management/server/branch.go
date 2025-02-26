package server

import (
	"context"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
)

func (s *OmsUserManagementServer) CreateBranch(ctx context.Context, req *grpc.CreateBranchRequest) (*grpc.CreateBranchResponse, error) {
	res, err := s.service.Branch().CreateBranch(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) UpdateBranch(ctx context.Context, req *grpc.UpdateBranchRequest) (*grpc.UpdateBranchResponse, error) {
	res, err := s.service.Branch().UpdateBranch(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetBranches(ctx context.Context, req *grpc.GetBranchesRequest) (*grpc.GetBranchesResponse, error) {
	res, err := s.service.Branch().GetBranches(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) DeleteBranch(ctx context.Context, req *grpc.DeleteBranchRequest) (*grpc.DeleteBranchResponse, error) {
	res, err := s.service.Branch().DeleteBranch(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetBranchById(ctx context.Context, req *grpc.GetBranchByIdRequest) (*grpc.GetBranchByIdResponse, error) {
	res, err := s.service.Branch().GetBranchById(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

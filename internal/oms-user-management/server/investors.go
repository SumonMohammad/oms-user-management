package server

import (
	"context"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *OmsUserManagementServer) CreateInvestor(ctx context.Context, req *grpc.CreateInvestorRequest) (*grpc.CreateInvestorResponse, error) {
	res, err := s.service.Investor().CreateInvestor(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) UpdateInvestor(ctx context.Context, req *grpc.UpdateInvestorRequest) (*grpc.UpdateInvestorResponse, error) {
	res, err := s.service.Investor().UpdateInvestor(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}
	return res, nil
}

func (s *OmsUserManagementServer) GetInvestorById(ctx context.Context, req *grpc.GetInvestorByIdRequest) (*grpc.GetInvestorByIdResponse, error) {
	res, err := s.service.Investor().GetInvestorById(ctx, req.UserId)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *OmsUserManagementServer) GetInvestors(ctx context.Context, req *grpc.GetInvestorsRequest) (*grpc.GetInvestorsResponse, error) {
	response, err := s.service.Investor().GetInvestors(ctx, req.Page, req.Limit)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, status.Errorf(codes.Internal, "Failed to get OMS users with Investors: %v", err)
	}
	return response, nil
}

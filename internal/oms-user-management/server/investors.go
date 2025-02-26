package server

import (
	"context"
	"gitlab.techetronventures.com/core/grpctest/pkg/grpc"
)

func (s *GrpctestServer) CreateInvestor(ctx context.Context, req *grpc.CreateInvestorRequest) (*grpc.CreateInvestorResponse, error) {
	res, err := s.service.Investor().CreateInvestor(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) UpdateInvestor(ctx context.Context, req *grpc.UpdateInvestorRequest) (*grpc.UpdateInvestorResponse, error) {
	res, err := s.service.Investor().UpdateInvestor(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

func (s *GrpctestServer) GetInvestors(ctx context.Context, req *grpc.GetInvestorsRequest) (*grpc.GetInvestorsResponse, error) {
	res, err := s.service.Investor().GetInvestors(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}
func (s *GrpctestServer) GetInvestorByIdOrEmail(ctx context.Context, req *grpc.GetInvestorByIdOrEmailRequest) (*grpc.GetInvestorByIdOrEmailResponse, error) {
	res, err := s.service.Investor().GetInvestorByIdOrEmail(ctx, req)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, err
	}

	return res, nil
}

package pg

import (
	"context"
	"github.com/uptrace/bun"
	"gitlab.techetronventures.com/core/backend/pkg/log"
	model "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
	//"time"
)

func (db *DB) Branch() model.Branches {
	return &Branch{
		IDB: db.DB,
		log: db.log,
	}
}

func (db *Tx) Branch() model.Branches {
	return &Branch{
		IDB: db.Tx,
		log: db.log,
	}
}

type Branch struct {
	bun.IDB
	log *log.Logger
}

func (s *Branch) CreateBranch(ctx context.Context, branch *model.Branch) error {
	_, err := s.NewInsert().Model(branch).
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *Branch) UpdateBranch(ctx context.Context, branch *model.Branch) error {
	branch.UpdatedAt = time.Now()
	_, err := s.NewUpdate().Model(branch).
		Where("id = ?", branch.ID).
		ExcludeColumn("created_at").
		Column("updated_at").
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *Branch) GetBranches(ctx context.Context, req *grpc.GetBranchesRequest) ([]*model.Branch, int64, error) {
	branches := []*model.Branch{}
	offset := (req.PaginationRequest.PageToken - 1) * req.PaginationRequest.PageSize
	query := s.NewSelect().
		Model((*model.Branch)(nil)).
		Where("deleted_at IS NULL")
	count, err := query.Limit(int(req.PaginationRequest.PageSize)).
		Offset(int(offset)).
		Order("created_at DESC").
		ScanAndCount(ctx, &branches)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, 0, err
	}

	return branches, int64(count), nil
}

func (s *Branch) GetBranchById(ctx context.Context, req *grpc.GetBranchByIdRequest) (*grpc.GetBranchByIdResponse, error) {
	branch := &model.Branch{}
	query := s.NewSelect().Model(branch).Where("deleted_at IS NULL")
	query = query.Where("id = ?", req.Id)
	err := query.Scan(ctx)
	if err != nil {
		msg := "Failed to fetch Broker House"
		return nil, status.Error(codes.Internal, msg)
	}

	// Map the result to the response format
	response := &grpc.GetBranchByIdResponse{
		Id:              branch.ID,
		BranchName:      branch.BranchName,
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
		BranchType:      branch.BranchType,
		BrokerHouseId:   branch.BrokerHouseId,
	}

	return response, nil
}

func (s *Branch) DeleteBranch(ctx context.Context, branch *model.Branch) error {
	branch.DeletedAt = time.Now()
	_, err := s.NewUpdate().
		Model(branch).
		Column("deleted_at").
		Where("id = ?", branch.ID).
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

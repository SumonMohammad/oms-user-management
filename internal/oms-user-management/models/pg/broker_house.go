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

func (db *DB) BrokerHouse() model.BrokerHouses {
	return &BrokerHouse{
		IDB: db.DB,
		log: db.log,
	}
}

func (db *Tx) BrokerHouse() model.BrokerHouses {
	return &BrokerHouse{
		IDB: db.Tx,
		log: db.log,
	}
}

type BrokerHouse struct {
	bun.IDB
	log *log.Logger
}

func (s *BrokerHouse) CreateBrokerHouse(ctx context.Context, brokerHouse *model.BrokerHouse) error {
	_, err := s.NewInsert().Model(brokerHouse).
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *BrokerHouse) UpdateBrokerHouse(ctx context.Context, brokerHouse *model.BrokerHouse) error {
	brokerHouse.UpdatedAt = time.Now()
	_, err := s.NewUpdate().Model(brokerHouse).
		Where("id = ?", brokerHouse.ID).
		ExcludeColumn("created_at").
		Column("updated_at").
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

func (s *BrokerHouse) GetBrokerHouses(ctx context.Context, req *grpc.GetBrokerHousesRequest) ([]*model.BrokerHouse, int64, error) {
	brokerHouses := []*model.BrokerHouse{}
	offset := (req.PaginationRequest.PageToken - 1) * req.PaginationRequest.PageSize
	query := s.NewSelect().
		Model((*model.BrokerHouse)(nil)).
		Where("deleted_at IS NULL")
	count, err := query.Limit(int(req.PaginationRequest.PageSize)).
		Offset(int(offset)).
		Order("created_at DESC").
		ScanAndCount(ctx, &brokerHouses)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return nil, 0, err
	}

	return brokerHouses, int64(count), nil
}

func (s *BrokerHouse) GetBrokerHouseById(ctx context.Context, req *grpc.GetBrokerHouseByIdRequest) (*grpc.GetBrokerHouseByIdResponse, error) {
	brokerHouse := &model.BrokerHouse{}
	query := s.NewSelect().Model(brokerHouse).Where("deleted_at IS NULL")
	query = query.Where("id = ?", req.Id)
	err := query.Scan(ctx)
	if err != nil {
		msg := "Failed to fetch Broker House"
		return nil, status.Error(codes.Internal, msg)
	}

	// Map the result to the response format
	response := &grpc.GetBrokerHouseByIdResponse{
		Id:              brokerHouse.ID,
		BrokerHouseName: brokerHouse.BrokerHouseName,
		ShortName:       brokerHouse.ShortName,
		Description:     brokerHouse.Description,
		Address:         brokerHouse.Address,
		PhoneNumber:     brokerHouse.PhoneNumber,
		TelephoneNumber: brokerHouse.TelephoneNumber,
		EmailAddress:    brokerHouse.EmailAddress,
		ValidCurrency:   brokerHouse.ValidCurrency,
		Status:          brokerHouse.Status,
		IsEnabled:       brokerHouse.IsEnabled,
		CountryCode:     brokerHouse.CountryCode,
	}

	return response, nil
}

func (s *BrokerHouse) DeleteBrokerHouse(ctx context.Context, brokerHouse *model.BrokerHouse) error {
	brokerHouse.DeletedAt = time.Now()
	_, err := s.NewUpdate().
		Model(brokerHouse).
		Column("deleted_at").
		Where("id = ?", brokerHouse.ID).
		Exec(ctx)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	return nil
}

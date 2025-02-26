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

type Employees interface {
	CreateEmployee(ctx context.Context, req *grpc.CreateEmployeeRequest) (*grpc.CreateEmployeeResponse, error)
	UpdateEmployee(ctx context.Context, req *grpc.UpdateEmployeeRequest) (*grpc.UpdateEmployeeResponse, error)
	GetEmployees(ctx context.Context, req *grpc.GetEmployeesRequest) (*grpc.GetEmployeesResponse, error)
	DeleteEmployee(ctx context.Context, req *grpc.DeleteEmployeeRequest) (*grpc.DeleteEmployeeResponse, error)
	GetEmployeeByIdOrEmail(ctx context.Context, req *grpc.GetEmployeeByIdOrEmailRequest) (*grpc.GetEmployeeByIdOrEmailResponse, error)
}

type Employee struct {
	service *GrpctestService
}

type EmployeeReceiver struct {
	*GrpctestService
}

func (ms *GrpctestService) Employee() Employees {
	return &EmployeeReceiver{
		ms,
	}
}

func (s *EmployeeReceiver) CreateEmployee(ctx context.Context, req *grpc.CreateEmployeeRequest) (*grpc.CreateEmployeeResponse, error) {

	employee := &model.Employee{
		FullName:     req.FullName,
		Designation:  req.Designation,
		EmailAddress: req.EmailAddress,
		PhoneNumber:  req.PhoneNumber,
		NidNumber:    req.NidNumber,
		IsDeleted:    req.IsDeleted,
		IsEnabled:    req.IsEnabled,
		Status:       req.Status,
	}

	// Call the database method to create the trader
	err := s.db.Employee().CreateEmployee(ctx, employee)
	if err != nil {
		msg := "failed to create employee"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.CreateEmployeeResponse{
		Code: 0,
	}, nil
}

func (s *EmployeeReceiver) UpdateEmployee(ctx context.Context, req *grpc.UpdateEmployeeRequest) (*grpc.UpdateEmployeeResponse, error) {

	employee := &model.Employee{
		ID:           req.Id,
		FullName:     req.FullName,
		Designation:  req.Designation,
		EmailAddress: req.EmailAddress,
		PhoneNumber:  req.PhoneNumber,
		NidNumber:    req.NidNumber,
		IsDeleted:    req.IsDeleted,
		IsEnabled:    req.IsEnabled,
		Status:       req.Status,
	}

	// Call the database method to create the trader
	err := s.db.Employee().UpdateEmployee(ctx, employee)
	if err != nil {
		msg := "failed to update oms user"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.UpdateEmployeeResponse{
		Code: 0,
	}, nil
}
func (s *EmployeeReceiver) GetEmployees(ctx context.Context, req *grpc.GetEmployeesRequest) (*grpc.GetEmployeesResponse, error) {

	res, count, err := s.db.Employee().GetEmployees(ctx, req)
	if err != nil {
		msg := "failed to get employee"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}
	employees := []*grpc.GetEmployeesResponseEmployeeList{}

	for _, item := range res {

		employee := &grpc.GetEmployeesResponseEmployeeList{
			Id:           item.ID,
			FullName:     item.FullName,
			Designation:  item.Designation,
			EmailAddress: item.EmailAddress,
			PhoneNumber:  item.PhoneNumber,
			NidNumber:    item.NidNumber,
			IsDeleted:    item.IsDeleted,
			IsEnabled:    item.IsEnabled,
			Status:       item.Status,
		}
		employees = append(employees, employee)
	}
	return &grpc.GetEmployeesResponse{
		Employees: employees,
		PaginationResponse: &grpc.PaginationInfoResponse{
			TotalRecordCount: int32(count),
		},
	}, nil
}

func (s *EmployeeReceiver) DeleteEmployee(ctx context.Context, req *grpc.DeleteEmployeeRequest) (*grpc.DeleteEmployeeResponse, error) {
	employee := &model.Employee{
		ID: req.Id,
	}

	// Call the database method to create the trader
	err := s.db.Employee().DeleteEmployee(ctx, employee)
	if err != nil {
		msg := "failed to delete employee"
		s.log.Error(ctx, msg, zap.Error(err))
		return nil, err
	}

	// Return a successful response
	return &grpc.DeleteEmployeeResponse{
		Code: 0,
	}, nil
}

func (s *EmployeeReceiver) GetEmployeeByIdOrEmail(ctx context.Context, req *grpc.GetEmployeeByIdOrEmailRequest) (*grpc.GetEmployeeByIdOrEmailResponse, error) {
	// Call the database layer to fetch the trader by ID or email
	employee, err := s.db.Employee().GetEmployeeByIdOrEmail(ctx, req)
	if err != nil {
		// Handle error if no trader is found or any other error occurs
		s.log.Error(ctx, "Failed to fetch Employee by ID or email", zap.Error(err))
		if err.Error() == "Employee not found" {
			return nil, status.Error(codes.NotFound, "Employee not found")
		}
		return nil, status.Error(codes.Internal, "Failed to fetch Employee")
	}

	// Map the result from the database layer to the GRPC response
	response := &grpc.GetEmployeeByIdOrEmailResponse{
		Id:           employee.Id,
		FullName:     employee.FullName,
		Designation:  employee.Designation,
		EmailAddress: employee.EmailAddress,
		PhoneNumber:  employee.PhoneNumber,
		NidNumber:    employee.NidNumber,
		IsEnabled:    employee.IsEnabled,
		Status:       employee.Status,
		IsDeleted:    employee.IsDeleted,
	}

	return response, nil
}

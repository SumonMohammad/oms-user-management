package client

import (
	"context"
	"github.com/gogo/protobuf/types"
	pb "gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"google.golang.org/grpc"
)

type Clients interface {
	HealthCheck(ctx context.Context) (*pb.HealthCheckResponse, error)
}

type OmsUserManagement struct {
	client pb.OmsUserManagementClient
}

func NewOmsUserManagementClient(target string) (Clients, error) {

	m := &OmsUserManagement{}
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		return nil, err
	}
	// defer conn.Close()
	m.client = pb.NewOmsUserManagementClient(conn)

	return m, err
}

func (m *OmsUserManagement) HealthCheck(ctx context.Context) (*pb.HealthCheckResponse, error) {
	return m.client.HealthCheck(ctx, &types.Empty{})
}

// Map User Permission API's
func (m *OmsUserManagement) GetUserPermissionsByUserId(ctx context.Context, req *pb.GetUserPermissionsByUserIdRequest) (*pb.GetUserPermissionsByUserIdResponse, error) {
	return m.client.GetUserPermissionsByUserId(ctx, req)
}
func (m *OmsUserManagement) CreateMapUserPermission(ctx context.Context, req *pb.CreateMapUserPermissionRequest) (*pb.CreateMapUserPermissionResponse, error) {
	return m.client.CreateMapUserPermission(ctx, req)
}
func (m *OmsUserManagement) UpdateMapUserPermission(ctx context.Context, req *pb.UpdateMapUserPermissionRequest) (*pb.UpdateMapUserPermissionResponse, error) {
	return m.client.UpdateMapUserPermission(ctx, req)
}
func (m *OmsUserManagement) DeleteMapUserPermission(ctx context.Context, req *pb.DeleteMapUserPermissionRequest) (*pb.DeleteMapUserPermissionResponse, error) {
	return m.client.DeleteMapUserPermission(ctx, req)
}

// Roles API's
func (m *OmsUserManagement) CreateRole(ctx context.Context, req *pb.CreateRoleRequest) (*pb.CreateRoleResponse, error) {
	return m.client.CreateRole(ctx, req)
}

func (m *OmsUserManagement) UpdateRole(ctx context.Context, req *pb.UpdateRoleRequest) (*pb.UpdateRoleResponse, error) {
	return m.client.UpdateRole(ctx, req)
}

func (m *OmsUserManagement) DeleteRole(ctx context.Context, req *pb.DeleteRoleRequest) (*pb.DeleteRoleResponse, error) {
	return m.client.DeleteRole(ctx, req)
}

func (m *OmsUserManagement) GetRoles(ctx context.Context, req *pb.GetRolesRequest) (*pb.GetRolesResponse, error) {
	return m.client.GetRoles(ctx, req)
}

// Permissions API's

func (m *OmsUserManagement) CreatePermission(ctx context.Context, req *pb.CreatePermissionRequest) (*pb.CreatePermissionResponse, error) {
	return m.client.CreatePermission(ctx, req)
}

func (m *OmsUserManagement) UpdatePermission(ctx context.Context, req *pb.UpdatePermissionRequest) (*pb.UpdatePermissionResponse, error) {
	return m.client.UpdatePermission(ctx, req)
}
func (m *OmsUserManagement) DeletePermission(ctx context.Context, req *pb.DeletePermissionRequest) (*pb.DeletePermissionResponse, error) {
	return m.client.DeletePermission(ctx, req)
}
func (m *OmsUserManagement) GetPermissions(ctx context.Context, req *pb.GetPermissionsRequest) (*pb.GetPermissionsResponse, error) {
	return m.client.GetPermissions(ctx, req)
}

//Traders APIs

func (m *OmsUserManagement) CreateTrader(ctx context.Context, req *pb.CreateTraderRequest) (*pb.CreateTraderResponse, error) {
	return m.client.CreateTrader(ctx, req)
}

func (m *OmsUserManagement) UpdateTrader(ctx context.Context, req *pb.UpdateTraderRequest) (*pb.UpdateTraderResponse, error) {
	return m.client.UpdateTrader(ctx, req)
}
func (m *OmsUserManagement) GetTraderById(ctx context.Context, req *pb.GetTraderByIdRequest) (*pb.GetTraderByIdResponse, error) {
	return m.client.GetTraderById(ctx, req)
}

// Broker Admins API
func (m *OmsUserManagement) CreateBrokerAdmin(ctx context.Context, req *pb.CreateBrokerAdminRequest) (*pb.CreateBrokerAdminResponse, error) {
	return m.client.CreateBrokerAdmin(ctx, req)
}

func (m *OmsUserManagement) UpdateBrokerAdmin(ctx context.Context, req *pb.UpdateBrokerAdminRequest) (*pb.UpdateBrokerAdminResponse, error) {
	return m.client.UpdateBrokerAdmin(ctx, req)
}

func (m *OmsUserManagement) GetBrokerAdminById(ctx context.Context, req *pb.GetBrokerAdminByIdRequest) (*pb.GetBrokerAdminByIdResponse, error) {
	return m.client.GetBrokerAdminById(ctx, req)
}

// Employee APIs
func (m *OmsUserManagement) CreateEmployee(ctx context.Context, req *pb.CreateEmployeeRequest) (*pb.CreateEmployeeResponse, error) {
	return m.client.CreateEmployee(ctx, req)
}

func (m *OmsUserManagement) UpdateEmployee(ctx context.Context, req *pb.UpdateEmployeeRequest) (*pb.UpdateEmployeeResponse, error) {
	return m.client.UpdateEmployee(ctx, req)
}

func (m *OmsUserManagement) GetEmployeeById(ctx context.Context, req *pb.GetEmployeeByIdRequest) (*pb.GetEmployeeByIdResponse, error) {
	return m.client.GetEmployeeById(ctx, req)
}

// Branches API
func (m *OmsUserManagement) CreateBranch(ctx context.Context, req *pb.CreateBranchRequest) (*pb.CreateBranchResponse, error) {
	return m.client.CreateBranch(ctx, req)
}

func (m *OmsUserManagement) UpdateBranch(ctx context.Context, req *pb.UpdateBranchRequest) (*pb.UpdateBranchResponse, error) {
	return m.client.UpdateBranch(ctx, req)
}

func (m *OmsUserManagement) GetBranchById(ctx context.Context, req *pb.GetBranchByIdRequest) (*pb.GetBranchByIdResponse, error) {
	return m.client.GetBranchById(ctx, req)
}
func (m *OmsUserManagement) GetBranches(ctx context.Context, req *pb.GetBranchesRequest) (*pb.GetBranchesResponse, error) {
	return m.client.GetBranches(ctx, req)
}

func (m *OmsUserManagement) DeleteBranch(ctx context.Context, req *pb.DeleteBranchRequest) (*pb.DeleteBranchResponse, error) {
	return m.client.DeleteBranch(ctx, req)
}

//Broker Houses

func (m *OmsUserManagement) CreateBrokerHouse(ctx context.Context, req *pb.CreateBrokerHouseRequest) (*pb.CreateBrokerHouseResponse, error) {
	return m.client.CreateBrokerHouse(ctx, req)
}

func (m *OmsUserManagement) UpdateBrokerHouse(ctx context.Context, req *pb.UpdateBrokerHouseRequest) (*pb.UpdateBrokerHouseResponse, error) {
	return m.client.UpdateBrokerHouse(ctx, req)
}

func (m *OmsUserManagement) GetBrokerHouseById(ctx context.Context, req *pb.GetBrokerHouseByIdRequest) (*pb.GetBrokerHouseByIdResponse, error) {
	return m.client.GetBrokerHouseById(ctx, req)
}
func (m *OmsUserManagement) GetBrokerHouses(ctx context.Context, req *pb.GetBrokerHousesRequest) (*pb.GetBrokerHousesResponse, error) {
	return m.client.GetBrokerHouses(ctx, req)
}

func (m *OmsUserManagement) DeleteBrokerHouse(ctx context.Context, req *pb.DeleteBrokerHouseRequest) (*pb.DeleteBrokerHouseResponse, error) {
	return m.client.DeleteBrokerHouse(ctx, req)
}

// Investors APIS
func (m *OmsUserManagement) CreateInvestor(ctx context.Context, req *pb.CreateInvestorRequest) (*pb.CreateInvestorResponse, error) {
	return m.client.CreateInvestor(ctx, req)
}

func (m *OmsUserManagement) UpdateInvestor(ctx context.Context, req *pb.UpdateInvestorRequest) (*pb.UpdateInvestorResponse, error) {
	return m.client.UpdateInvestor(ctx, req)
}
func (m *OmsUserManagement) GetInvestorById(ctx context.Context, req *pb.GetInvestorByIdRequest) (*pb.GetInvestorByIdResponse, error) {
	return m.client.GetInvestorById(ctx, req)
}

func (m *OmsUserManagement) GetInvestors(ctx context.Context, req *pb.GetInvestorsRequest) (*pb.GetInvestorsResponse, error) {
	return m.client.GetInvestors(ctx, req)
}

//Map User Roles

func (m *OmsUserManagement) CreateMapUserRole(ctx context.Context, req *pb.CreateMapUserRoleRequest) (*pb.CreateMapUserRoleResponse, error) {
	return m.client.CreateMapUserRole(ctx, req)
}

func (m *OmsUserManagement) UpdateMapUserRole(ctx context.Context, req *pb.UpdateMapUserRoleRequest) (*pb.UpdateMapUserRoleResponse, error) {
	return m.client.UpdateMapUserRole(ctx, req)
}
func (m *OmsUserManagement) GetMapUserRoles(ctx context.Context, req *pb.GetMapUserRolesRequest) (*pb.GetMapUserRolesResponse, error) {
	return m.client.GetMapUserRoles(ctx, req)
}

func (m *OmsUserManagement) GetMapUserRoleById(ctx context.Context, req *pb.GetMapUserRoleByIdRequest) (*pb.GetMapUserRoleByIdResponse, error) {
	return m.client.GetMapUserRoleById(ctx, req)
}

func (m *OmsUserManagement) DeleteMapUserRole(ctx context.Context, req *pb.DeleteMapUserRoleRequest) (*pb.DeleteMapUserRoleResponse, error) {
	return m.client.DeleteMapUserRole(ctx, req)
}

// Map Roles & Permissions APIs
func (m *OmsUserManagement) CreateMapRolePermission(ctx context.Context, req *pb.CreateMapRolePermissionRequest) (*pb.CreateMapRolePermissionResponse, error) {
	return m.client.CreateMapRolePermission(ctx, req)
}

func (m *OmsUserManagement) UpdateMapRolePermission(ctx context.Context, req *pb.UpdateMapRolePermissionRequest) (*pb.UpdateMapRolePermissionResponse, error) {
	return m.client.UpdateMapRolePermission(ctx, req)
}
func (m *OmsUserManagement) GetMapRolePermissions(ctx context.Context, req *pb.GetMapRolePermissionsRequest) (*pb.GetMapRolePermissionsResponse, error) {
	return m.client.GetMapRolePermissions(ctx, req)
}

func (m *OmsUserManagement) GetMapRolePermissionById(ctx context.Context, req *pb.GetMapRolePermissionByIdRequest) (*pb.GetMapRolePermissionByIdResponse, error) {
	return m.client.GetMapRolePermissionById(ctx, req)
}

func (m *OmsUserManagement) DeleteMapRolePermission(ctx context.Context, req *pb.DeleteMapRolePermissionRequest) (*pb.DeleteMapRolePermissionResponse, error) {
	return m.client.DeleteMapRolePermission(ctx, req)
}

//Mapping TWS with Trader

func (m *OmsUserManagement) CreateTraderTws(ctx context.Context, req *pb.CreateTraderTwsMapRequest) (*pb.CreateTraderTwsMapResponse, error) {
	return m.client.CreateTraderTws(ctx, req)
}

func (m *OmsUserManagement) GetTradersTws(ctx context.Context, req *pb.GetTradersTwsMapRequest) (*pb.GetTradersTwsMapResponse, error) {
	return m.client.GetTradersTws(ctx, req)
}
func (m *OmsUserManagement) UpdateTraderTws(ctx context.Context, req *pb.UpdateTraderTwsMapRequest) (*pb.UpdateTraderTwsMapResponse, error) {
	return m.client.UpdateTraderTws(ctx, req)
}

func (m *OmsUserManagement) DeleteTraderTws(ctx context.Context, req *pb.DeleteTraderTwsMapRequest) (*pb.DeleteTraderTwsMapResponse, error) {
	return m.client.DeleteTraderTws(ctx, req)
}

// Trader Team API's

func (m *OmsUserManagement) CreateTeam(ctx context.Context, req *pb.CreateTraderTeamRequest) (*pb.CreateTraderTeamResponse, error) {
	return m.client.CreateTeam(ctx, req)
}
func (m *OmsUserManagement) UpdateTeam(ctx context.Context, req *pb.UpdateTraderTeamRequest) (*pb.UpdateTraderTeamResponse, error) {
	return m.client.UpdateTeam(ctx, req)
}
func (m *OmsUserManagement) GetTeams(ctx context.Context, req *pb.GetTraderTeamRequest) (*pb.GetTraderTeamResponse, error) {
	return m.client.GetTeams(ctx, req)
}
func (m *OmsUserManagement) DeleteTeam(ctx context.Context, req *pb.DeleteTraderTeamRequest) (*pb.DeleteTraderTeamResponse, error) {
	return m.client.DeleteTeam(ctx, req)
}

//Mapping trader with team Creation

func (m *OmsUserManagement) CreateMapTeam(ctx context.Context, req *pb.CreateTraderMapRequest) (*pb.CreateTraderMapResponse, error) {
	return m.client.CreateMapTeam(ctx, req)
}
func (m *OmsUserManagement) UpdateMapTeam(ctx context.Context, req *pb.UpdateTraderMapRequest) (*pb.UpdateTraderMapResponse, error) {
	return m.client.UpdateMapTeam(ctx, req)
}
func (m *OmsUserManagement) GetMappedTeams(ctx context.Context, req *pb.GetTraderMapRequest) (*pb.GetTraderMapResponse, error) {
	return m.client.GetMappedTeams(ctx, req)
}
func (m *OmsUserManagement) DeleteMappedTeam(ctx context.Context, req *pb.DeleteTraderMapRequest) (*pb.DeleteTraderMapResponse, error) {
	return m.client.DeleteMappedTeam(ctx, req)
}

// TWS API's
func (m *OmsUserManagement) CreateTws(ctx context.Context, req *pb.CreateTwsRequest) (*pb.CreateTwsResponse, error) {
	return m.client.CreateTws(ctx, req)
}
func (m *OmsUserManagement) UpdateTws(ctx context.Context, req *pb.UpdateTwsRequest) (*pb.UpdateTwsResponse, error) {
	return m.client.UpdateTws(ctx, req)
}
func (m *OmsUserManagement) GetTws(ctx context.Context, req *pb.GetTwsRequest) (*pb.GetTwsResponse, error) {
	return m.client.GetTws(ctx, req)
}
func (m *OmsUserManagement) DeleteTws(ctx context.Context, req *pb.DeleteTwsRequest) (*pb.DeleteTwsResponse, error) {
	return m.client.DeleteTws(ctx, req)
}

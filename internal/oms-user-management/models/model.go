package model

import (
	"context"
)

type Repository interface {
	Health() HealthRepository
	Trader() Traders
	Investor() Investors
	User() Users
	TraderTws() TradersTws
	TraderTeam() TraderTeams
	MapTraderTeam() MapTraderTeams
	Role() Roles
	TWS() Tws
	Employee() Employees
	BrokerAdmin() BrokerAdmins
	BrokerHouse() BrokerHouses
	Branch() Branches
	AuditLog() AuditLogs
	Permission() Permissions
	MapUserRole() MapUserRoles
	MapRolePermission() MapRolePermissions
	MapUserPermission() MapUserPermissions
}

// TxFunc represents function to run in an SQL-transaction.
type TxFunc func(ctx context.Context, tx Repository) (err error)

type DB interface {
	// Repository access without transactions
	Repository
	// InTx runs given function in transaction
	InTx(context.Context, TxFunc) error
}

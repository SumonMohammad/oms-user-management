package service

import (
	"gitlab.techetronventures.com/core/backend/pkg/log"
	"gitlab.techetronventures.com/core/backend/pkg/pubsub"
	"gitlab.techetronventures.com/core/backend/pkg/rabbitmq"
	"gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/config"
	"gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models"
)

type Service interface {
	Health() HealthService
	Trader() Traders
	Investor() Investors
	TraderTws() TradersTws
	TraderTeam() TraderTeams
	MapTraderTeam() MapTraderTeams
	TWS() Tws
	Employee() Employees
	BrokerAdmin() BrokerAdmins
	BrokerHouse() BrokerHouses
	Branch() Branches
	AuditLog() AuditLogs
	Permission() Permissions
	Role() Roles
	MapUserPermission() MapUserPermissions
	MapUserRole() MapUserRoles
	MapRolePermission() MapRolePermissions
}

type OmsUserManagementService struct {
	db           model.DB
	config       *config.Config
	log          *log.Logger
	queue        rabbitmq.Queue
	PubSubConfig *pubsub.Config
}

type Init struct {
	Db        model.DB
	Cnf       *config.Config
	Log       *log.Logger
	Qu        rabbitmq.Queue
	PubConfig *pubsub.Config
}

func New(init *Init) (*OmsUserManagementService, error) {
	return &OmsUserManagementService{
		db:           init.Db,
		config:       init.Cnf,
		log:          init.Log.Named("service"),
		queue:        init.Qu,
		PubSubConfig: init.PubConfig,
	}, nil
}

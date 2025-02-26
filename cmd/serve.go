package main

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"gitlab.techetronventures.com/core/backend/pkg/log"
	"gitlab.techetronventures.com/core/backend/pkg/migration"
	"gitlab.techetronventures.com/core/backend/pkg/server"
	"gitlab.techetronventures.com/core/backend/pkg/sigint"
	oum "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/server"
	"gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/service"
	"gitlab.techetronventures.com/core/oms-user-management/pkg/grpc"
	"gitlab.techetronventures.com/core/oms-user-management/sql"
	"go.uber.org/zap"
	"net/http"
)

func run(ctx context.Context, init *service.Init) {
	conf := init.Cnf
	logger := log.New().Named("oms-user-management")

	// migrations
	migrateDirection, migrateOnly := conf.MigrationDirectionFlag()
	migrateDB, err := migration.SQLFromUrl(conf.DB.URL)
	if err != nil {
		panic(err)
	}

	migrations := sql.GetMigrations()
	err = migration.MigrateFromFS(migrateDB, migrateDirection, "oms_user_management", migrations)
	if err != nil {
		panic(err)
	}
	_ = migrateDB.Close()

	if migrateOnly {
		logger.Info(ctx, "Migration complete, exiting")
		return
	}

	// service
	var s service.Service
	s, err = service.New(init)
	if err != nil {
		panic(err)
	}

	// server initialization
	var srv *server.Server
	reg := prometheus.NewRegistry()
	if srv, err = server.New(logger, conf.GRPC, reg); err != nil {
		panic(err)
	}

	logger.Info(ctx, "server port will be: ", zap.String("network", conf.GRPC.Bind))
	meServer := oum.New(logger, s)
	grpc.RegisterOmsUserManagementServer(srv.GRPC(), meServer)

	// Create a HTTP server for prometheus.
	httpServer := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		EnableOpenMetrics: true,
	}), Addr: fmt.Sprintf("0.0.0.0:%d", 9092)}

	// Start your http server for prometheus.
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			logger.Error(ctx, err.Error())
		}
	}()

	srv.Run()
	defer func(srv *server.Server) {
		_ = srv.Close()
	}(srv)

	sigint.Wait()
	logger.Info(ctx, "stopping server!!")
}

func serve(ctx context.Context, init *service.Init) *cobra.Command {
	return &cobra.Command{
		Use: "serve",
		Run: func(cmd *cobra.Command, args []string) {
			run(ctx, init)
		},
	}
}

package main

import (
  "context"
  "github.com/spf13/cobra"
  "gitlab.techetronventures.com/core/backend/pkg/log"
  "gitlab.techetronventures.com/core/backend/pkg/rabbitmq"
  "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/config"
  "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/models/pg"
  "gitlab.techetronventures.com/core/oms-user-management/internal/oms-user-management/service"
  "os"
)

func main() {
  var (
    ctx    = context.Background()
    conf   = config.New().Load()
    err    error
    logger = log.New()
  )

  logger = logger.Named("oms-user-management")

  // database
  postgres, err := pg.New(conf.DB, logger)
  if err != nil {
    panic(err)
  }
  defer func(postgres *pg.DB) {
    _ = postgres.DB.Close()
  }(postgres)

  // Message Queue
  qu, err := rabbitmq.New(conf.Queue, logger)
  if err != nil {
    panic(err)
  }

  // service
  omsUserManagementInit := &service.Init{
    Db:        postgres,
    Cnf:       conf,
    Log:       logger,
    Qu:        qu,
    PubConfig: conf.PubSub,
  }

  var rootCmd = &cobra.Command{}
  rootCmd.AddCommand(serve(ctx, omsUserManagementInit))

  if err = rootCmd.Execute(); err != nil {
    logger.Error(ctx, err.Error())
    os.Exit(1)
  }
}

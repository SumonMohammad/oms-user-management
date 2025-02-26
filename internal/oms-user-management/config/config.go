package config

import (
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"gitlab.techetronventures.com/core/backend/pkg/bundb"
	"gitlab.techetronventures.com/core/backend/pkg/migration"
	"gitlab.techetronventures.com/core/backend/pkg/pubsub"
	"gitlab.techetronventures.com/core/backend/pkg/rabbitmq"
	"gitlab.techetronventures.com/core/backend/pkg/server"
	"os"
)

// Config of entire application.
type Config struct {
	Url                 string               `json:"url" yaml:"url" toml:"url" mapstructure:"url"`
	GRPC                *server.Config       `json:"grpc" yaml:"grpc" toml:"grpc" mapstructure:"grpc"`
	DB                  *bundb.Config        `json:"db" yaml:"db" toml:"db" mapstructure:"db"` // nolint
	MigrateDirection    migration.Direction  `json:"migrate"`
	PubSub              *pubsub.Config       `json:"pubsub" yaml:"pubsub" toml:"pubsub" mapstructure:"pubsub"`
	Queue               *rabbitmq.Config     `json:"rabbit_mq" yaml:"rabbit_mq" toml:"rabbit_mq" mapstructure:"rabbit_mq"`
}

// New default configurations.
func New() (conf *Config) {
	conf = new(Config)
	conf.GRPC = server.NewConfig()
	conf.PubSub = pubsub.NewConfig()
	return
}

// Load the Config from configuration files. This method panics on error.
func (c *Config) Load() *Config {

	consulPath := os.Getenv("OMS_USER_MANAGEMENT_CONSUL_PATH")
	consulURL := os.Getenv("OMS_USER_MANAGEMENT_CONSUL_URL")

	viper.AddRemoteProvider("consul", consulURL, consulPath)
	viper.SetConfigType("yaml") // Need to explicitly set this to json

	err := viper.ReadRemoteConfig()
	if err != nil {
		panic(err)
	}
	viper.Unmarshal(c)

	migrate := c.MigrateDirection
	c.MigrateDirection = migration.Direction(migrate)
	if err = c.MigrateDirection.Check(); err != nil {
		panic(err)
	}

	return c
}

// MigrationDirectionFlag returns migration direction and migrateOnly flag
func (c *Config) MigrationDirectionFlag() (
	migrateDirection migration.Direction, migrateOnly bool) {
	if c.MigrateDirection == "" {
		return migration.DirectionUp, false
	}

	return c.MigrateDirection, true
}


package config

import (
	"strings"

	"github.com/dorm-parcel-manager/dpm/common/client"
	"github.com/dorm-parcel-manager/dpm/common/db"
	"github.com/dorm-parcel-manager/dpm/common/rabbitmq"
	"github.com/dorm-parcel-manager/dpm/common/server"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

type Config struct {
	Server   *server.Config
	Client   *client.Config
	DB       *db.Config
	Rabbitmq *rabbitmq.Config
}

var ConfigSet = wire.NewSet(
	ProvideConfig,
	wire.FieldsOf(new(*Config), "Server", "Client", "DB"),
)

func ProvideConfig() *Config {
	config := Config{}

	viper.SetDefault("server.port", 4002)

	viper.SetDefault("client.userserviceurl", "localhost:4001")

	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "5434")
	viper.SetDefault("db.user", "dpm")
	viper.SetDefault("db.password", "dpm")
	viper.SetDefault("db.dbname", "dpm")

	viper.SetDefault("rabbitmq.host", "localhost")
	viper.SetDefault("rabbitmq.port", "5672")
	viper.SetDefault("rabbitmq.user", "dpm")
	viper.SetDefault("rabbitmq.password", "dpm")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	viper.Unmarshal(&config)

	return &config
}

package config

import (
	"strings"

	"github.com/dorm-parcel-manager/dpm/common/mongo"
	"github.com/dorm-parcel-manager/dpm/common/rabbitmq"

	"github.com/dorm-parcel-manager/dpm/common/client"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

type Config struct {
	Server   *serverConfig
	Client   *client.Config
	DB       *mongo.Config
	Rabbitmq *rabbitmq.Config
}

type serverConfig struct {
	Port int
}

var ConfigSet = wire.NewSet(
	ProvideConfig,
	wire.FieldsOf(new(*Config), "Server", "Client", "DB"),
)

func ProvideConfig() *Config {
	config := Config{}

	viper.SetDefault("server.port", 4003)

	viper.SetDefault("client.notificationserviceurl", "localhost:4003")

	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "27017")
	viper.SetDefault("db.user", "mongo")
	viper.SetDefault("db.password", "mongo")
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

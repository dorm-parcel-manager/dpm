package config

import (
	"strings"

	"github.com/dorm-parcel-manager/dpm/common/client"
	"github.com/dorm-parcel-manager/dpm/common/db"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

type Config struct {
	Server *serverConfig
	Client *client.Config
	DB     *db.Config
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
	viper.SetDefault("db.port", "5434")
	viper.SetDefault("db.user", "dpm")
	viper.SetDefault("db.password", "dpm")
	viper.SetDefault("db.dbname", "dpm")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	viper.Unmarshal(&config)

	return &config
}

package config

import (
	"strings"

	"github.com/dorm-parcel-manager/dpm/pkg/db"
	"github.com/dorm-parcel-manager/dpm/pkg/server"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

type Config struct {
	Server *server.Config
	DB     *db.Config
}

var ConfigSet = wire.NewSet(
	ProvideConfig,
	wire.FieldsOf(new(*Config), "Server", "DB"),
)

func ProvideConfig() *Config {
	config := Config{}

	viper.SetDefault("server.port", 4002)
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

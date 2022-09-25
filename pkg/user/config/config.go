package config

import (
	"strings"

	"github.com/dorm-parcel-manager/dpm/pkg/server"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

type Config struct {
	Server *server.Config
}

var ConfigSet = wire.NewSet(
	ProvideConfig,
	wire.FieldsOf(new(*Config), "Server"),
)

func ProvideConfig() *Config {
	config := Config{}

	viper.SetDefault("server.port", 4000)

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	viper.Unmarshal(&config)

	return &config
}

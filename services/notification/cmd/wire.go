//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/dorm-parcel-manager/dpm/common/mongo"
	"github.com/dorm-parcel-manager/dpm/common/rabbitmq"
	"github.com/dorm-parcel-manager/dpm/services/notification/config"
	"github.com/dorm-parcel-manager/dpm/services/notification/server"
	"github.com/dorm-parcel-manager/dpm/services/notification/service"
	"github.com/google/wire"
)

func InitializeServer() (*server.Server, func(), error) {
	wire.Build(
		config.ConfigSet,
		mongo.NewDb,
		rabbitmq.NewChannel,
		service.NewNotificationService,
		server.NewServer,
	)
	return &server.Server{}, nil, nil
}

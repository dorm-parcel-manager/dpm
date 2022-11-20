// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"github.com/dorm-parcel-manager/dpm/common/mongo"
	"github.com/dorm-parcel-manager/dpm/common/rabbitmq"
	"github.com/dorm-parcel-manager/dpm/services/notification/config"
	"github.com/dorm-parcel-manager/dpm/services/notification/server"
	"github.com/dorm-parcel-manager/dpm/services/notification/service"
)

// Injectors from wire.go:

func InitializeServer() (*server.Server, func(), error) {
	configConfig := config.ProvideConfig()
	mongoConfig := configConfig.DB
	database, err := mongo.NewDb(mongoConfig)
	if err != nil {
		return nil, nil, err
	}
	rabbitmqConfig := configConfig.Rabbitmq
	channel, err := rabbitmq.NewChannel(rabbitmqConfig)
	if err != nil {
		return nil, nil, err
	}
	vapidKeyPair := configConfig.VapidKeyPair
	notificationService := service.NewNotificationService(database, channel, vapidKeyPair)
	serverConfig := configConfig.Server
	serverServer := server.NewServer(notificationService, serverConfig)
	return serverServer, func() {
	}, nil
}

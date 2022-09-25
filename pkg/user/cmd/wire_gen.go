// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"github.com/dorm-parcel-manager/dpm/pkg/api"
	"github.com/dorm-parcel-manager/dpm/pkg/db"
	"github.com/dorm-parcel-manager/dpm/pkg/server"
	"github.com/dorm-parcel-manager/dpm/pkg/user/config"
	"github.com/dorm-parcel-manager/dpm/pkg/user/service"
	"google.golang.org/grpc"
)

// Injectors from wire.go:

func InitializeServer() (*server.Server, error) {
	configConfig := config.ProvideConfig()
	serverConfig := configConfig.Server
	dbConfig := configConfig.DB
	gormDB, err := db.NewDb(dbConfig)
	if err != nil {
		return nil, err
	}
	userServiceServer := service.NewUserServiceServer(gormDB)
	grpcServer := ProvideGrpcServer(userServiceServer)
	serverServer := server.NewServer(serverConfig, grpcServer)
	return serverServer, nil
}

// wire.go:

func ProvideGrpcServer(userService api.UserServiceServer) *grpc.Server {
	s := grpc.NewServer()
	api.RegisterUserServiceServer(s, userService)
	return s
}

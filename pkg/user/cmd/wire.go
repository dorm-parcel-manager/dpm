//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/dorm-parcel-manager/dpm/pkg/api"
	"github.com/dorm-parcel-manager/dpm/pkg/db"
	"github.com/dorm-parcel-manager/dpm/pkg/server"
	"github.com/dorm-parcel-manager/dpm/pkg/user/config"
	"github.com/dorm-parcel-manager/dpm/pkg/user/service"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

func InitializeServer() (*server.Server, error) {
	wire.Build(
		config.ConfigSet,
		server.NewServer,
		db.NewDb,
		ProvideGrpcServer,
		service.NewUserServiceServer,
	)
	return &server.Server{}, nil
}

func ProvideGrpcServer(userService api.UserServiceServer) *grpc.Server {
	s := grpc.NewServer()
	api.RegisterUserServiceServer(s, userService)
	return s
}

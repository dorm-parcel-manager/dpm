//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/dorm-parcel-manager/dpm/common/db"
	"github.com/dorm-parcel-manager/dpm/common/pb"
	"github.com/dorm-parcel-manager/dpm/common/server"
	"github.com/dorm-parcel-manager/dpm/services/user/config"
	"github.com/dorm-parcel-manager/dpm/services/user/service"
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

func ProvideGrpcServer(userService pb.UserServiceServer) *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, userService)
	return s
}

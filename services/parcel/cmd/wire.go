//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/dorm-parcel-manager/dpm/common/client"
	"github.com/dorm-parcel-manager/dpm/common/pb"
	"github.com/dorm-parcel-manager/dpm/common/rabbitmq"
	"github.com/dorm-parcel-manager/dpm/common/server"
	"github.com/dorm-parcel-manager/dpm/services/parcel/config"
	"github.com/dorm-parcel-manager/dpm/services/parcel/service"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

func InitializeServer() (*server.Server, func(), error) {
	wire.Build(
		config.ConfigSet,
		server.NewServer,
		client.Providers,
		rabbitmq.NewChannel,
		ProvideGrpcServer,
		service.NewParcelServiceServer,
	)
	return &server.Server{}, nil, nil
}

func ProvideGrpcServer(parcelService pb.ParcelServiceServer) *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterParcelServiceServer(s, parcelService)
	return s
}

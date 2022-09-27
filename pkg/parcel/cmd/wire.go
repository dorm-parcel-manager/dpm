//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/dorm-parcel-manager/dpm/pkg/parcel/config"
	"github.com/dorm-parcel-manager/dpm/pkg/parcel/service"
	"github.com/dorm-parcel-manager/dpm/pkg/pb"
	"github.com/dorm-parcel-manager/dpm/pkg/server"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

func InitializeServer() (*server.Server, error) {
	wire.Build(
		config.ConfigSet,
		server.NewServer,
		ProvideGrpcServer,
		service.NewParcelServiceServer,
	)
	return &server.Server{}, nil
}

func ProvideGrpcServer(parcelService pb.ParcelServiceServer) *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterParcelServiceServer(s, parcelService)
	return s
}

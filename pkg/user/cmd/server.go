package cmd

import (
	"github.com/dorm-parcel-manager/dpm/pkg/api"
	"github.com/dorm-parcel-manager/dpm/pkg/server"
	"github.com/dorm-parcel-manager/dpm/pkg/user/service"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func RunServer() error {
	s := grpc.NewServer()
	api.RegisterUserServiceServer(s, service.NewUserServiceServer())

	server := server.NewServer(&server.Config{Port: 4000}, s)
	err := server.Start()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

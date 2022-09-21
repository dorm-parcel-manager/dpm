package cmd

import (
	"fmt"
	"github.com/dorm-parcel-manager/dpm/pkg/api"
	"github.com/dorm-parcel-manager/dpm/pkg/user/service"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

func RunServer() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 4000))
	if err != nil {
		return errors.WithStack(err)
	}

	s := grpc.NewServer()
	api.RegisterUserServiceServer(s, service.NewUserServiceServer())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		log.Printf("stopping...")
		s.Stop()
	}()

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

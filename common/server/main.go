package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Config struct {
	Port int
}

type Server struct {
	config     *Config
	grpcServer *grpc.Server
}

func NewServer(config *Config, grpcServer *grpc.Server) *Server {
	return &Server{
		config:     config,
		grpcServer: grpcServer,
	}
}

func (s *Server) Start() error {
	reflection.Register(s.grpcServer)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.Port))
	if err != nil {
		return errors.WithStack(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		log.Printf("stopping...")
		s.grpcServer.Stop()
	}()

	log.Printf("server listening at %v", lis.Addr())
	if err := s.grpcServer.Serve(lis); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

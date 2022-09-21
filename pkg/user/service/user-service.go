package service

import (
	"context"
	"github.com/dorm-parcel-manager/dpm/pkg/api"
)

type userServiceServer struct {
	api.UnimplementedUserServiceServer
}

func NewUserServiceServer() api.UserServiceServer {
	return &userServiceServer{}
}

func (s *userServiceServer) Hello(ctx context.Context, in *api.Empty) (*api.Empty, error) {
	return &api.Empty{}, nil
}

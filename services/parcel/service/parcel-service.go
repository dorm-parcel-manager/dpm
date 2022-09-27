package service

import (
	"context"
	"fmt"

	"github.com/dorm-parcel-manager/dpm/common/pb"
	"github.com/pkg/errors"
)

type parcelServiceServer struct {
	pb.UnimplementedParcelServiceServer

	userService pb.UserServiceClient
}

func NewParcelServiceServer(userService pb.UserServiceClient) (pb.ParcelServiceServer, error) {
	return &parcelServiceServer{
		userService: userService,
	}, nil
}

func (s *parcelServiceServer) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	response, err := s.userService.Hello(ctx, in)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	message := fmt.Sprintf("User service: %v", response.Message)
	return &pb.HelloResponse{
		Message: message,
	}, nil
}

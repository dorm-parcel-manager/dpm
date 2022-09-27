package service

import (
	"github.com/dorm-parcel-manager/dpm/common/pb"
)

type parcelServiceServer struct {
	pb.UnimplementedParcelServiceServer
}

func NewParcelServiceServer() (pb.ParcelServiceServer, error) {
	return &parcelServiceServer{}, nil
}

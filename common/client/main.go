package client

import (
	"github.com/dorm-parcel-manager/dpm/common/pb"
	"github.com/dorm-parcel-manager/dpm/common/utils"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	UserServiceUrl   string
	ParcelServiceUrl string
}

var Providers = wire.NewSet(
	ProvideUserServiceClient,
	ProvideParcelServiceClient,
)

var opts = []grpc.DialOption{
	grpc.WithTransportCredentials(insecure.NewCredentials()),
}

func ProvideUserServiceClient(config *Config) (pb.UserServiceClient, func(), error) {
	conn, err := grpc.Dial(config.UserServiceUrl, opts...)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	client := pb.NewUserServiceClient(conn)
	return client, utils.ErrorToFatal(conn.Close), nil
}

func ProvideParcelServiceClient(config *Config) (pb.ParcelServiceClient, func(), error) {
	conn, err := grpc.Dial(config.ParcelServiceUrl, opts...)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	client := pb.NewParcelServiceClient(conn)
	return client, utils.ErrorToFatal(conn.Close), nil
}

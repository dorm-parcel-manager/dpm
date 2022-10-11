package client

import (
	"github.com/dorm-parcel-manager/dpm/common/pb"
	sd "github.com/dorm-parcel-manager/dpm/common/service-discovery"
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

func ProvideUserServiceClient(sdClint *sd.ServiceDiscoveryClient) (pb.UserServiceClient, func(), error) {

	url := sdClint.ServiceDiscovery(string(sd.ServiceName_USER_SERVICE))

	conn, err := grpc.Dial(url, opts...)

	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	client := pb.NewUserServiceClient(conn)
	return client, utils.ErrorToFatal(conn.Close), nil
}

func ProvideParcelServiceClient(sdClint *sd.ServiceDiscoveryClient) (pb.ParcelServiceClient, func(), error) {
	
	url := sdClint.ServiceDiscovery(string(sd.ServiceName_USER_SERVICE))
	
	conn, err := grpc.Dial(url, opts...)

	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	client := pb.NewParcelServiceClient(conn)
	return client, utils.ErrorToFatal(conn.Close), nil
}

package appcontext

import (
	"fmt"
	"github.com/dorm-parcel-manager/dpm/common/pb"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type appContextImpl struct {
	context *pb.Context
}

func NewAppContext(context *pb.Context) AppContext {
	return &appContextImpl{
		context: context,
	}
}

func (a *appContextImpl) RequireLoggedIn() error {
	if a.context.UserId == 0 {
		return status.Error(codes.Unauthenticated, "unauthenticated")
	}
	return nil
}

func (a *appContextImpl) requireUserType(userType pb.UserType) error {
	err := a.RequireLoggedIn()
	if err != nil {
		return errors.WithStack(err)
	}
	if a.context.UserType != userType {
		return status.Error(codes.PermissionDenied, fmt.Sprintf("user type must be %v", userType))
	}
	return nil
}

func (a *appContextImpl) RequireStudent() error {
	return a.requireUserType(pb.UserType_TYPE_STUDENT)
}

func (a *appContextImpl) RequireStaff() error {
	return a.requireUserType(pb.UserType_TYPE_STAFF)
}

func (a *appContextImpl) RequireAdmin() error {
	return a.requireUserType(pb.UserType_TYPE_ADMIN)
}

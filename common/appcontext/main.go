package appcontext

import "github.com/dorm-parcel-manager/dpm/common/pb"

type AppContext interface {
	RequireLoggedIn() error

	RequireStudent() error
	RequireStaff() error
	RequireAdmin() error

	GetUserID() uint
	GetUserType() pb.UserType
}

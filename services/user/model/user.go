package model

import (
	"github.com/dorm-parcel-manager/dpm/common/pb"
	"time"
)

type User struct {
	ID        uint
	OauthID   string
	Email     string
	FirstName string
	LastName  string
	Type      pb.UserType

	CreatedAt time.Time
	UpdatedAt time.Time
}

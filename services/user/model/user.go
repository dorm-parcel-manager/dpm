package model

import (
	"time"

	"github.com/dorm-parcel-manager/dpm/common/pb"
)

type User struct {
	ID        uint
	OauthID   string
	Email     string
	FirstName string
	LastName  string
	Picture   string
	Type      pb.UserType

	CreatedAt time.Time
	UpdatedAt time.Time
}

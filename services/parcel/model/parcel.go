package model

import (
	"time"

	"github.com/dorm-parcel-manager/dpm/common/pb"
)

type Parcel struct {
	ID               uint `gorm:"primaryKey"`
	OwnerID          uint
	ArrivalDate      time.Time
	Name             string
	TransportCompany string
	TrackingNumber   string
	Sender           string
	Description      string
	Status           pb.ParcelStatus
}

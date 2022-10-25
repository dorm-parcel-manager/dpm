package model

import (
	"database/sql"

	"github.com/dorm-parcel-manager/dpm/common/pb"
)

type Parcel struct {
	ID               uint `gorm:"primaryKey"`
	OwnerID          uint
	ArrivalDate      sql.NullTime
	PickedUpDate     sql.NullTime
	Name             string
	TransportCompany string
	TrackingNumber   string
	Sender           string
	Description      string
	Status           pb.ParcelStatus
	CreatedAt        sql.NullTime
	UpdatedAt        sql.NullTime
}

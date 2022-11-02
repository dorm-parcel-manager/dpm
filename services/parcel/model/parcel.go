package model

import (
	"database/sql"

	"github.com/dorm-parcel-manager/dpm/common/pb"
)

type Parcel struct {
	ID               uint 	`gorm:"primaryKey"`
	OwnerID          uint	`gorm:"not null"`
	ArrivalDate      sql.NullTime
	PickedUpDate     sql.NullTime
	Name             string	`gorm:"not null"`
	TransportCompany string	`gorm:"not null"`
	TrackingNumber   string	`gorm:"unique;not null"`
	Sender           string	`gorm:"not null"`
	Description      string
	Status           pb.ParcelStatus
	CreatedAt        sql.NullTime `gorm:"not null"`
	UpdatedAt        sql.NullTime
}

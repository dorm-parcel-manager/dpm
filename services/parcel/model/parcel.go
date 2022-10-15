package model

import (
	"time"

	"github.com/dorm-parcel-manager/dpm/common/pb"
)

type Parcel struct {
	ID                uint `gorm:"primaryKey"`
	Owner_ID          uint
	Arrival_Date      time.Time
	Transport_Company string
	Tracking_Number   string
	Sender            string
	Description       string
	Status            pb.ParcelStatus
}

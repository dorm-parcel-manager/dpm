package model

import (
	"time"

	"github.com/dorm-parcel-manager/dpm/common/pb"
)

type Parcel struct {
	ID        			uint
	Owner_ID			uint
	Arrival_Date		time.Time
	Transport_Company  	string
	Tracking_Number   	string
	Sender      		string
	Status 				pb.ParcelStatus
}

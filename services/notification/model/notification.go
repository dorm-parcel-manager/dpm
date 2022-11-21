package model

import (
	webpush "github.com/SherClockHolmes/webpush-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Title    string             `bson:"title"`
	Message  string             `bson:"message"`
	Link     string             `bson:"link"`
	UserID   uint               `bson:"userId"`
	Read     bool               `bson:"read"`
	UnixTime int64              `bson:"unixTime"`
}

type NotificationSubscription struct {
	UserID       uint                 `bson:"userId"`
	Subscription webpush.Subscription `bson:"subscription"`
}

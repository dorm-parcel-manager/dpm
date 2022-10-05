package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Notification struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Title    string             `bson:"title"`
	Message  string             `bson:"message"`
	Link     string             `bson:"link"`
	UserID   string             `bson:"userId"`
	Read     bool               `bson:"read"`
	UnixTime int64              `bson:"unixTime"`
}

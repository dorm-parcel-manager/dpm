package model

import (
	"time"
)

type User struct {
	ID        uint
	OauthID   string
	Email     string
	FirstName string
	LastName  string
	
	CreatedAt time.Time
	UpdatedAt time.Time
}

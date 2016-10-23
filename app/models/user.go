package models

import "time"

type Base struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct { // example user fields
	Username string
	Password []byte `json:"-"`
	Base
}

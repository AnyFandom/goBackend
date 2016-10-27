package models

type User struct { // example user fields
	Username string
	Password []byte `json:"-"`
	Base
}

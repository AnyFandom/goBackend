package models

type User struct { // example user fields
	Username string `json:"username"`
	Password []byte `json:"-"`
	Base
}

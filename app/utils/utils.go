package utils

import "crypto/sha256"

func HashPassword(password string) []byte {
	var h = sha256.New()
	h.Write([]byte(password))
	var encr = h.Sum(nil)
	return encr
}

type Location struct {
	Location string
}

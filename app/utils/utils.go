package utils

import (
	"crypto/sha256"
	"fmt"
	"reflect"

	jwt "github.com/dgrijalva/jwt-go"
)

func HashPassword(password string) []byte {
	var h = sha256.New()
	h.Write([]byte(password))
	var encr = h.Sum(nil)
	return encr
}

type Location struct {
	Location string
}

func CreateToken(id uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"id": id})

	tokenString, err := token.SignedString([]byte("test-key"))
	if err != nil {
		panic(err)
	}
	return tokenString
}

func ParseToken(tokenString string) jwt.MapClaims {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("test-key"), nil
	})

	if err != nil {
		panic(err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Print(claims)
	if !ok && !token.Valid {
		panic(err)
	}
	return claims
}

type IncludeItem struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func contains(s []IncludeItem, e IncludeItem) bool {
	for _, a := range s {
		if reflect.DeepEqual(a, e) {
			return true
		}
	}
	return false
}

func AppendInclude(slice []IncludeItem, item IncludeItem) []IncludeItem {
	if !contains(slice, item) {
		slice = append(slice, item)
	}
	return slice
}

func ExtendInclude(slice []IncludeItem, items []IncludeItem) []IncludeItem {
	for _, v := range items {
		slice = AppendInclude(slice, v)
	}
	return slice
}

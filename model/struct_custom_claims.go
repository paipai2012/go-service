package model

import "github.com/dgrijalva/jwt-go"

type Payload struct {
	UserId string `json:"userId"`
}

type CustomClaims struct {
	*jwt.StandardClaims
	*Payload
}

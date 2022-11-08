package model

import "github.com/dgrijalva/jwt-go"

type Payload struct {
	UserId string `json:"userId"`
}

type AgentPayload struct {
	AppId string `json:"appId"`
}
type CustomClaims struct {
	*jwt.StandardClaims
	*Payload
}

type AgentCustomClaims struct {
	*jwt.StandardClaims
	*AgentPayload
}

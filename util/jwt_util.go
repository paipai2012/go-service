package util

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtUtil struct{}

var (
	mySigningKey = []byte("SigningKey")
)

func CreateJWT() (string, error) {
	// claims := jwt.StandardClaims{
	// 	NotBefore: int64(time.Now().Unix()),
	// 	ExpiresAt: int64(time.Now().Unix() + 1000),
	// 	Issuer:    "Bitch",
	// }
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	// ss, err := token.SignedString(mySigningKey)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// log.Print(ss)

	claims := &jwt.StandardClaims{}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

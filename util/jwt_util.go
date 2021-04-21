package util

import (
	"fmt"
	"log"
	"moose-go/model"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtUtil struct{}

var verifyKey = []byte("moose-go")

func GeneratorJwt(payload *model.Payload) (string, error) {
	claims := &model.CustomClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60 * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		payload,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(verifyKey)
}

func ParseJwt(tokenStr string) *jwt.Token {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return verifyKey, nil
	})

	if token.Valid {
		log.Println("You look nice today")
		return token
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			log.Printf("%v", ve)
			return nil
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			log.Printf("%v", ve)
			return nil
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func ParseBearerToken(token string) string {
	tokens := strings.Split(token, " ")
	if len(tokens) != 2 || !strings.EqualFold("Bearer", tokens[0]) {
		return ""
	}
	return tokens[1]
}

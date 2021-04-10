package util

import (
	"fmt"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 普通测试
func TestCreateJWT(t *testing.T) {
	// claims := &jwt.StandardClaims{
	// 	ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
	// 	IssuedAt:  time.Now().Unix(),
	// 	Subject:   "test jwt",
	// }

	// claims := make(jwt.MapClaims)
	// user := &model.UserInfo{
	// 	UserId:   "123132131",
	// 	UserName: "测试",
	// }
	// // claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
	// claims["exp"] = time.Now().Add(time.Millisecond * 1).Unix()
	// claims["user"] = user

	claims := &CustomClaims{
		&jwt.StandardClaims{
			// set the expire time
			// see http://tools.ietf.org/html/draft-ietf-oauth-json-web-token-20#section-4.1.4
			ExpiresAt: time.Now().Add(time.Millisecond * 1).Unix(),
		},
		nil,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtStr, err := token.SignedString(verifyKey)
	if err != nil {
		t.Error()
	}
	t.Log(jwtStr)
}

func TestParseJWT(t *testing.T) {
	jwtStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTczMjg3MDYsInVzZXJJZCI6IjEyMyIsInVzZXJOYW1lIjoidGVzdCIsInBob25lIjoiIiwiZ2VuZGVyIjoiIiwiZW1haWwiOiIiLCJhZGRyZXNzIjoiIiwiZGVzY3JpcHRpb24iOiIiLCJhdmF0YXIiOiIiLCJjcmVhdGVUaW1lIjoiIiwidXBkYXRlVGltZSI6IiJ9.pn3oYOg9YsGBwxGinxPWgNpp5G5M1sPM6U0FGL66spA"
	token, err := jwt.Parse(jwtStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return verifyKey, nil
	})

	if token.Valid {
		fmt.Println("You look nice today")
		fmt.Println(token.Claims)
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}

}

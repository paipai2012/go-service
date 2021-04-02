package middleware

import (
	"log"
	"moose-go/api"
	"moose-go/util"
	"strings"

	"github.com/gin-gonic/gin"
)

var anonymous = []string{
	"/api/v1/account/register",
	"/api/v1/account/login",
	"/api/v1/qrcode/get",
	"/api/v1/qrcode/ask",
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if util.In(path, anonymous) {
			c.Next()
			return
		}

		token := c.GetHeader("Authorization")
		log.Println("auth check token... ", token)

		if token == "" {
			panic(api.JwtValidationErr)
		}

		// Bearer xxxx
		tokens := strings.Split(token, " ")
		if len(tokens) != 2 || !strings.EqualFold("Bearer", tokens[0]) {
			panic(api.JwtValidationErr)
		}
		util.ParseJwt(tokens[1])
		c.Next()
	}
}

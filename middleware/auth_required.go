package middleware

import (
	"log"
	"moose-go/api"
	"moose-go/util"

	"github.com/gin-gonic/gin"
)

var anonymous = []string{
	"/api/v1/account/register",
	"/api/v1/account/login",
	"/api/v1/qrcode/get",
	"/api/v1/qrcode/ask",
	"/api/v1/sms/send",
	"/api/v1/user/info",
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if util.In(path, anonymous) {
			c.Next()
			return
		}

		bearerToken := c.GetHeader("Authorization")
		log.Println("auth check token... ", bearerToken)

		if bearerToken == "" {
			panic(api.JwtValidationErr)
		}

		// Bearer xxxx
		token := util.ParseBearerToken(bearerToken)
		util.ParseJwt(token)
		c.Next()
	}
}

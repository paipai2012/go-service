package middleware

import (
	"log"
	"moose-go/api"
	"moose-go/common"
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

const IS_TEST_ENV = true

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if util.In(path, anonymous) {
			c.Next()
			return
		}

		if !IS_TEST_ENV {
			bearerToken := c.GetHeader("Authorization")
			log.Println("auth check token... ", bearerToken)

			if bearerToken == "" {
				common.JSON(c, api.JsonError(api.JwtValidationErr))
				return
			}

			// Bearer xxxx
			token := util.ParseBearerToken(bearerToken)
			if token == "" {
				common.JSON(c, api.JsonError(api.JwtValidationErr))
				return
			}

			jwtToken := util.ParseJwt(token)
			if jwtToken == nil {
				common.JSON(c, api.JsonError(api.JwtValidationErr))
				return
			}
		}

		c.Next()
	}
}

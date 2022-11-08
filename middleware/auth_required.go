package middleware

import (
	"fmt"
	"sale-service/api"
	"sale-service/common"
	"sale-service/util"

	"github.com/gin-gonic/gin"
)

var anonymous = []string{
	"/api/account/agent_login",
	"/api/luck/get",
	"/api/luck/addDraw",
}

var app_anonymous = []string{
	"/api/account/agent_login",
}

const IS_TEST_ENV = false

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		if util.In(path, app_anonymous) {
			c.Next()
			return
		}
		bearerAgentToken := c.GetHeader("AgentAuthorization")
		if bearerAgentToken == "" {
			common.JSON(c, api.JsonError(api.JwtAgentValidationErr))
			return
		}

		// Bearer xxxx
		agentToken := util.ParseBearerToken(bearerAgentToken)
		if agentToken == "" {
			common.JSON(c, api.JsonError(api.JwtAgentValidationErr))
			return
		}

		jwtAgentToken := util.ParseJwt(agentToken)
		fmt.Print(jwtAgentToken)
		if jwtAgentToken == nil {
			common.JSON(c, api.JsonError(api.JwtAgentValidationErr))
			return
		}

		// if util.In(path, anonymous) {
		// 	c.Next()
		// 	return
		// }

		// // if !IS_TEST_ENV {
		// bearerToken := c.GetHeader("Authorization")

		// if bearerToken == "" {
		// 	common.JSON(c, api.JsonError(api.JwtValidationErr))
		// 	return
		// }

		// // Bearer xxxx
		// token := util.ParseBearerToken(bearerToken)
		// if token == "" {
		// 	common.JSON(c, api.JsonError(api.JwtValidationErr))
		// 	return
		// }

		// jwtToken := util.ParseJwt(token)
		// if jwtToken == nil {
		// 	common.JSON(c, api.JsonError(api.JwtValidationErr))
		// 	return
		// }

		c.Next()
	}
}

package controller

import (
	"sale-service/common"
	"sale-service/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (uc *UserController) RegisterRouter(app *gin.Engine) {
	group := app.Group("/api/v1/user")
	group.GET("/info", uc.Info)
}

func (uc *UserController) Info(c *gin.Context) {
	bearerToken := c.GetHeader("Authorization")
	userService := service.UserService{}
	common.JSON(c, userService.GetUserWithToken(bearerToken))
}

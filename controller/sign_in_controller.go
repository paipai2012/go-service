package controller

import (
	"moose-go/common"
	"moose-go/service"

	"github.com/gin-gonic/gin"
)

type SignInController struct{}

func (ac *SignInController) RegisterRouter(app *gin.Engine) {
	group := app.Group("/api/v1/signin")
	group.POST("/doSignIn", ac.doSignIn)
	group.POST("/getSignInList", ac.getSignInList)
}

func (sic *SignInController) doSignIn(c *gin.Context) {
	signInService := service.SignInService{}
	common.JSON(c, signInService.SignIn(c))
}

func (sic *SignInController) getSignInList(c *gin.Context) {
	signInService := service.SignInService{}
	common.JSON(c, signInService.GetCurrentList(c))
}

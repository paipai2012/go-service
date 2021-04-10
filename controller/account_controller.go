package controller

import (
	"log"
	"moose-go/common"
	"moose-go/model"
	"moose-go/service"

	"github.com/gin-gonic/gin"
)

type AccountController struct{}

func (ac *AccountController) RegisterRouter(app *gin.Engine) {
	group := app.Group("/api/v1/account")
	group.POST("/register", ac.Register)
	group.POST("/login", ac.Login)
}

func (ac *AccountController) Register(c *gin.Context) {
	var registerInfo model.RegisterInfo
	if err := c.BindJSON(&registerInfo); err != nil {
		log.Println(err.Error())
		common.FailedParam(c)
		return
	}
	accountService := service.AccountService{}
	common.JSON(c, accountService.Register(&registerInfo))
}

func (ac *AccountController) Login(c *gin.Context) {
	var loginInfo model.LoginInfo
	if err := c.BindJSON(&loginInfo); err != nil {
		log.Println(err.Error())
		common.FailedParam(c)
		return
	}
	accountService := service.AccountService{}
	common.JSON(c, accountService.Login(&loginInfo))
}

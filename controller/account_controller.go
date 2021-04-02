package controller

import (
	"log"
	"moose-go/api"
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
		panic(api.ErrParam.WithErrMsg(err.Error()))
	}
	log.Printf("receiver param %s", registerInfo)
	accountService := service.AccountService{}
	accountService.Register(&registerInfo)
	common.Success(c, 1)
}

func (ac *AccountController) Login(c *gin.Context) {
	var loginInfo model.LoginInfo
	if err := c.BindJSON(&loginInfo); err != nil {
		log.Println(err.Error())
		panic(api.ErrParam.WithErrMsg(err.Error()))
	}
	accountService := service.AccountService{}
	common.Success(c, accountService.Login(&loginInfo))
}

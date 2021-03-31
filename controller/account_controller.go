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
		common.Failed(c, api.ParseParamCode, err.Error())
		return
	}
	log.Printf("receiver param %s", registerInfo)
	accountService := service.AccountService{}
	accountService.AddUser(&registerInfo)
	common.Success(c, 1)
}

func (ac *AccountController) Login(c *gin.Context) {

	// c.ShouldBindJSON()
	common.Success(c, 1)
}

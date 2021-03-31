package controller

import (
	"moose-go/api"
	"moose-go/common"
	"moose-go/model"
	"moose-go/service"

	"github.com/gin-gonic/gin"
)

type AccountController struct{}

func init() {
}

func (ac *AccountController) RegisterRouter(app *gin.Engine) {
	group := app.Group("/api/v1/account")
	group.POST("/regist", ac.Regist)
	group.POST("/login", ac.Login)
}

func (ac *AccountController) Regist(c *gin.Context) {
	var registParam model.RegistParam
	if err := c.BindJSON(&registParam); err != nil {
		common.Failed(c, api.ParseParamCode, err.Error())
		return
	}

	accountService := service.AccountService{}
	accountService.Regist(&registParam)

	common.Success(c, 1)
}

func (ac *AccountController) Login(c *gin.Context) {

	// c.ShouldBindJSON()
	common.Success(c, 1)
}

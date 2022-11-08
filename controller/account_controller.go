package controller

import (
	"log"
	"sale-service/common"
	"sale-service/model"
	"sale-service/service"

	"github.com/gin-gonic/gin"
)

type AccountController struct{}

func (ac *AccountController) RegisterRouter(app *gin.Engine) {
	group := app.Group("/api/account")
	group.POST("/register", ac.Register)
	group.POST("/login", ac.Login)
	group.POST("/agent_login", ac.AgentLogin)
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

func (ac *AccountController) AgentLogin(c *gin.Context) {
	var loginInfo model.AgentLoginInfo
	if err := c.BindJSON(&loginInfo); err != nil {
		log.Println(err.Error())
		common.FailedParam(c)
		return
	}
	accountService := service.AccountService{}
	common.JSON(c, accountService.AgentLogin(&loginInfo))
}

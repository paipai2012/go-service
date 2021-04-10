package controller

import (
	"log"
	"moose-go/common"
	"moose-go/model"
	"moose-go/service"

	"github.com/gin-gonic/gin"
)

type SmsController struct{}

func (sc *SmsController) RegisterRouter(app *gin.Engine) {
	group := app.Group("/api/v1/sms/")
	group.POST("/send", sc.send)
}

func (sc *SmsController) send(c *gin.Context) {
	var sms model.Sms
	if err := c.BindJSON(&sms); err != nil {
		log.Println(err.Error())
		common.FailedParam(c)
		return
	}

	smsService := service.SenderMessageService{}
	common.JSON(c, smsService.Send(&sms))
}

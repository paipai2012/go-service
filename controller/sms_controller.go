package controller

import (
	"log"
	"moose-go/api"
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
		panic(api.ErrParam.WithErrMsg(err.Error()))
	}
	smsService := service.SenderMessageService{}
	smsService.Send(&sms)
	common.Success(c, 1)
}

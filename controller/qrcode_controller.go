package controller

import (
	"context"
	"fmt"
	"moose-go/common"
	"moose-go/constant"
	"moose-go/engine"
	"moose-go/model"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type QRCodeController struct {
}

var ctx = context.Background()

var redisHelper *engine.RedisHelper

func (qrc *QRCodeController) RegisterRouter(app *gin.Engine) {
	redisHelper = engine.GetRedisHelper()
	group := app.Group("/api/v1/qrcode")
	group.GET("/get", qrc.GetQRCode)
}

func (qrc *QRCodeController) GetQRCode(c *gin.Context) {
	mTicket := uuid.NewV4().String()
	qrCodeUrl := fmt.Sprintf("http://192.168.1.100:8090/api/v1/qrcode?m_ticket&%s", mTicket)
	// result, err := _redisHelper.SetNX(ctx, constant.MOOSE_TICKET, mTicket, 10*time.Minute).Result()
	ticketKey := fmt.Sprintf(constant.MOOSE_TICKET, mTicket)
	result, err := redisHelper.HMSet(ctx, ticketKey, mTicket, 0).Result()
	if err != nil {
		common.Failed(c, err.Error())
		return
	}
	redisHelper.Expire(ctx, ticketKey, 10*time.Minute)
	fmt.Println(result)
	qrCodeInfo := &model.QRCodeInfo{
		CodeUrl: qrCodeUrl,
	}
	common.Success(c, qrCodeInfo, "获取二维码成功")
}

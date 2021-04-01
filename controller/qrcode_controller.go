package controller

import (
	"moose-go/common"
	"moose-go/service"

	"github.com/gin-gonic/gin"
)

type QRCodeController struct {
}

func (qrc *QRCodeController) RegisterRouter(app *gin.Engine) {
	group := app.Group("/api/v1/qrcode/")
	group.GET("/get", qrc.GetQRCode)
	group.GET("/ask", qrc.AskQRCode)
	group.POST("/sanlogin", qrc.ScanLogin)
}

func (qrc *QRCodeController) GetQRCode(c *gin.Context) {
	qrCodeService := service.QRCodeService{}
	qrCodeInfo := qrCodeService.GenerateQRCode(c)
	common.Success(c, qrCodeInfo)
}

func (qrc *QRCodeController) AskQRCode(c *gin.Context) {
	qrCodeService := service.QRCodeService{}
	status := qrCodeService.AskQRCode(c)
	common.Success(c, status)
}

func (qrc *QRCodeController) ScanLogin(c *gin.Context) {
	qrCodeService := service.QRCodeService{}
	qrCodeService.ScanLogin(c)
	common.Success(c, 1)
}

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
	common.JSON(c, qrCodeService.GenerateQRCode(c))
}

func (qrc *QRCodeController) AskQRCode(c *gin.Context) {
	qrCodeService := service.QRCodeService{}
	common.JSON(c, qrCodeService.AskQRCode(c))
}

func (qrc *QRCodeController) ScanLogin(c *gin.Context) {
	qrCodeService := service.QRCodeService{}
	common.JSON(c, qrCodeService.ScanLogin(c))
}

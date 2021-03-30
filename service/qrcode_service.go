package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"moose-go/constant"
	"moose-go/engine"
	"moose-go/model"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type QRCodeService struct {
}

var ctx = context.Background()

func (qrs *QRCodeService) GenerateQRCode(c *gin.Context) *model.QRCodeInfo {
	mTicket := uuid.NewV4().String()
	qrCodeUrl := fmt.Sprintf("http://192.168.1.100:8090/qrcode?m_ticket=%s", mTicket)

	// result, err := _redisHelper.SetNX(ctx, constant.MOOSE_TICKET, mTicket, 3*time.Minute).Result()
	redisHelper := engine.GetRedisHelper()
	ticketKey := fmt.Sprintf(constant.MOOSE_SCAN_TICKET, mTicket)
	authInfo := &model.AuthInfo{Token: "", ScanStatus: 0}
	_, err := redisHelper.HMSet(ctx, ticketKey, mTicket, authInfo).Result()
	if err != nil {
		log.Println(err)
		panic("get qrcode url fail")
	}
	// 设置过期时间，三分钟
	redisHelper.Expire(ctx, ticketKey, 3*time.Minute)

	c.SetCookie("m_ticket", mTicket, 3*60, "/", "http://localhost:8090", true, true)

	qrCodeInfo := &model.QRCodeInfo{
		CodeUrl: qrCodeUrl,
	}
	return qrCodeInfo
}

func (qrs *QRCodeService) AskQRCode(c *gin.Context) *model.AuthInfo {
	mTicket, err := c.Cookie("m_ticket")
	log.Printf("the m_ticket %s", mTicket)

	if err != nil {
		log.Println(err)
		panic("retry get qrcode url")
	}

	redisHelper := engine.GetRedisHelper()
	ticketKey := fmt.Sprintf(constant.MOOSE_SCAN_TICKET, mTicket)
	result, err := redisHelper.HMGet(ctx, ticketKey, mTicket).Result()

	if err != nil {
		log.Println(err)
		panic("retry get qrcode url")
	}

	if len(result) <= 0 {
		panic("retry get qrcode url")
	}

	authStr := result[0]
	if authStr == nil {
		panic("retry get qrcode url")
	}

	var authInfo model.AuthInfo
	json.Unmarshal([]byte(result[0].(string)), &authInfo)
	log.Printf("key %s result %s %s", ticketKey, result, result[0].(string))
	return &authInfo
}

// 判断 m_ticket 是否存在
// 需要登录
// 校验 token
func (qrs *QRCodeService) ScanLogin(c *gin.Context) []interface{} {
	mTicket := c.Query("m_ticket")
	if mTicket == "" {
		panic("retry get qrcode url")
	}
	// 校验 m_ticket

	redisHelper := engine.GetRedisHelper()
	ticketKey := fmt.Sprintf(constant.MOOSE_SCAN_TICKET, mTicket)
	result, err := redisHelper.HMGet(ctx, ticketKey, mTicket).Result()

	if err != nil {
		log.Println(err)
		panic("retry get qrcode url")
	}
	var authInfo model.AuthInfo
	json.Unmarshal([]byte(result[0].(string)), &authInfo)
	log.Printf("key %s result %s %s", ticketKey, result, result[0].(string))
	return result
}

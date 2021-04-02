package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"moose-go/api"
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
	qrCodeUrl := fmt.Sprintf("http://192.168.1.100:7000/api/v1/qrcode/sanlogin?m_ticket=%s", mTicket)

	// result, err := _redisHelper.SetNX(ctx, constant.MOOSE_TICKET, mTicket, 3*time.Minute).Result()
	redisHelper := engine.GetRedisHelper()
	ticketKey := fmt.Sprintf(constant.MOOSE_SCAN_TICKET, mTicket)
	authInfo := &model.AuthInfo{Token: "", ScanStatus: 0}
	_, err := redisHelper.HMSet(ctx, ticketKey, mTicket, authInfo).Result()
	if err != nil {
		log.Println(err)
		panic(api.QRCodeGetFailErr)
	}
	// 设置过期时间，三分钟
	redisHelper.Expire(ctx, ticketKey, 3*time.Minute)

	c.SetCookie("m_ticket", mTicket, 3*60, "/", "localhost", true, true)

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
		panic(api.QRCodeRetryErr)
	}

	redisHelper := engine.GetRedisHelper()
	ticketKey := fmt.Sprintf(constant.MOOSE_SCAN_TICKET, mTicket)
	result, err := redisHelper.HMGet(ctx, ticketKey, mTicket).Result()

	if err != nil {
		panic(api.QRCodeRetryErr)
	}

	// 检查 ticket
	checkTicket(result)

	var authInfo model.AuthInfo
	json.Unmarshal([]byte(result[0].(string)), &authInfo)
	log.Printf("key %s result %s %s", ticketKey, result, result[0].(string))
	return &authInfo
}

// 需要登录，从 app 调用
//   - app 必须是登录状态
//   - app 扫描之后，携带 token 校验是否合法
// 判断 m_ticket 是否存在
// 校验 token 是否合法
func (qrs *QRCodeService) ScanLogin(c *gin.Context) {
	mTicket := c.Query("m_ticket")
	// 校验 m_ticket
	if mTicket == "" {
		panic(api.QRCodeRetryErr)
	}

	redisHelper := engine.GetRedisHelper()
	ticketKey := fmt.Sprintf(constant.MOOSE_SCAN_TICKET, mTicket)
	result, err := redisHelper.HMGet(ctx, ticketKey, mTicket).Result()

	if err != nil {
		log.Println(err)
		panic(api.QRCodeRetryErr)
	}

	// 检查 ticket
	checkTicket(result)

	var authInfo model.AuthInfo
	err = json.Unmarshal([]byte(result[0].(string)), &authInfo)
	if err != nil {
		log.Println(err)
		panic(api.QRCodeRetryErr)
	}

	authInfo.Status = 1

	mToken := c.Query("token")
	// if token valid success
	if mToken != "" {
		authInfo.Token = mToken
	}

	_, err = redisHelper.HMSet(ctx, ticketKey, mTicket, &authInfo).Result()
	if err != nil {
		log.Println(err)
		panic(api.QRCodeGetFailErr)
	}
	// 设置过期时间，三分钟
	redisHelper.Expire(ctx, ticketKey, 3*time.Minute).Result()
	log.Printf("key %s result %s", ticketKey, result)
}

func checkTicket(result []interface{}) {
	if len(result) <= 0 {
		panic(api.QRCodeGetFailErr)
	}

	authStr := result[0]
	if authStr == nil {
		panic(api.QRCodeGetFailErr)
	}
}

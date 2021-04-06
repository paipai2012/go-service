package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"moose-go/api"
	"moose-go/constant"
	"moose-go/engine"
	"moose-go/model"
	"time"
)

type SenderMessageService struct {
}

func (sms SenderMessageService) Send(s *model.Sms) {
	number := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	smsKey := fmt.Sprintf(constant.MOOSE_SMS, s.SmsType, s.Mobile)
	redisHelper := engine.GetRedisEngine()
	result, err := redisHelper.Set(context.Background(), smsKey, number, 5*time.Minute).Result()
	if err != nil {
		panic(api.SmsCodeSendErr)
	}
	log.Println("发送短信验证码成功", result, number)
}

func (sms SenderMessageService) CheckSms(s *model.Sms) string {
	smsKey := fmt.Sprintf(constant.MOOSE_SMS, s.SmsType, s.Mobile)
	redisHelper := engine.GetRedisEngine()
	smsCode, err := redisHelper.Get(context.Background(), smsKey).Result()
	if err != nil {
		panic(api.SmsCodeErr)
	}
	return smsCode
}

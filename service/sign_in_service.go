package service

import (
	"context"
	"fmt"
	"log"
	"sale-service/api"
	"sale-service/constant"
	"sale-service/engine"
	"sale-service/util"
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/go-redis/redis/v8"
)

type SignInService struct{}

const TEST_USER_ID = 1000

func buildSignInKey(userId int64) string {
	yearMoth := util.GetYearMonth()
	return fmt.Sprintf(constant.MOOSE_SIGN_IN, userId, yearMoth)
}

func (sis *SignInService) SignIn(c *gin.Context) *api.JsonResult {
	// bearerToken := c.GetHeader("Authorization")
	// userSrevice := UserService{}

	// result := userSrevice.getCurrentUserId(bearerToken)
	// if result.Code != 200 {
	// 	return result
	// }

	userId := TEST_USER_ID

	var redisHelper = engine.GetRedisEngine()

	signInKey := buildSignInKey(int64(userId))

	log.Printf("signin key %s \n", signInKey)

	offset := util.GetDay()

	for i := 1; i < 15; i++ {
		r, _ := redisHelper.SetBit(context.Background(), signInKey, offset+int64(i), 1).Result()
		log.Println(r)
	}

	return api.JsonData("success")
}

func (sis *SignInService) CheckStatus(c *gin.Context) *api.JsonResult {

	// bearerToken := c.GetHeader("Authorization")
	// userSrevice := UserService{}

	// result := userSrevice.getCurrentUserId(bearerToken)
	// if result.Code != 200 {
	// 	return result
	// }

	userId := TEST_USER_ID

	var redisHelper = engine.GetRedisEngine()
	signInKey := buildSignInKey(int64(userId))

	offset := util.GetDay()
	r, _ := redisHelper.GetBit(context.Background(), signInKey, offset).Result()
	return api.JsonData(r)
}

func (sis *SignInService) GetCurrentList(c *gin.Context) *api.JsonResult {
	userId := TEST_USER_ID

	var redisHelper = engine.GetRedisEngine()

	signInKey := buildSignInKey(int64(userId))

	// year := time.Now().Year()
	// month := time.Now().Month()
	// day := util.YearMonthDay(year, int(month))

	// signMap := make(map[string]bool, day)

	lengthOfMonth := util.YearMonthDay(time.Now().Year(), int(time.Now().Month()))
	log.Println("lengthOfMonth :: ", lengthOfMonth)

	r, err := redisHelper.BitField(context.Background(), signInKey, "GET", fmt.Sprintf("u%d", lengthOfMonth), "0").Result()

	if r != nil && len(r) > 0 {
		v := int64(0)
		if r[0] != 0 {
			v = r[0]
		}
		for i := lengthOfMonth; i > 0; i-- {
			log.Println("v>>1<<1 != v", v>>1<<1 != v, v)
			v >>= 1
		}
	}
	log.Println(r, err)

	// v := 10
	// log.Println(v>>1<<1 != v)

	return api.JsonData(r)
}

func getSignInfo(userId int64) {

}

func (sis *SignInService) GetRanking() *api.JsonResult {
	return api.JsonSuccess()
}

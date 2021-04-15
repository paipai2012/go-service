package controller

import (
	"context"
	"fmt"
	"log"
	"moose-go/common"
	"moose-go/constant"
	"moose-go/engine"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type SignInController struct{}

func (ac *SignInController) RegisterRouter(app *gin.Engine) {
	group := app.Group("/api/v1/signin")
	group.POST("/check", ac.check)
	group.POST("/list", ac.list)
}

func getCurrentDate() int64 {
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	currentDate, err := strconv.Atoi(fmt.Sprintf("%d%d%d", year, month, day))
	if err != nil {
		log.Printf("format error")
	}
	return int64(currentDate)
}

func (sic *SignInController) check(c *gin.Context) {
	var redisHelper = engine.GetRedisEngine()

	userId := 100001
	signInKey := fmt.Sprintf(constant.MOOSE_SIGN_IN, userId)
	currentDate := getCurrentDate()
	r, _ := redisHelper.SetBit(context.Background(), signInKey, currentDate, 1).Result()
	common.JSON(c, r)
}

func (sic *SignInController) list(c *gin.Context) {
	var redisHelper = engine.GetRedisEngine()
	currentDate := getCurrentDate()

	userId := 100001
	signInKey := fmt.Sprintf(constant.MOOSE_SIGN_IN, userId)
	r, _ := redisHelper.GetBit(context.Background(), signInKey, currentDate).Result()
	common.JSON(c, r)
}

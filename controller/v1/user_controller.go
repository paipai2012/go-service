package controller

import (
	"context"
	"encoding/json"
	"log"
	"moose-go/common"
	"moose-go/engine"
	"moose-go/model"
	"moose-go/service"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

var ctx = context.Background()

var redisHelper *engine.RedisHelper

func (uc *UserController) RegisterRouter(app *gin.Engine) {
	redisHelper = engine.GetRedisHelper()

	group := app.Group("/api/v1/user")
	group.POST("/add", uc.Add)
	group.GET("/get", uc.GetUser)
	group.GET("/list", uc.List)
	group.GET("/cache", uc.CacheUser)
	group.GET("/cache/get", uc.GetCacheUser)
}

func (uc *UserController) Add(c *gin.Context) {
	userName, _ := c.GetQuery("userName")
	userService := service.UserService{}
	row, err := userService.AddUser(userName)
	if err == nil && row > 0 {
		common.Success(c, 1, "添加用户成功")
		return
	}
	log.Fatal(err)
	common.Failed(c, "添加用户失败")
}

func (uc *UserController) GetUser(c *gin.Context) {
	userId := c.GetString("userId")
	userService := service.UserService{}
	common.Success(c, userService.GetUserByUserId(userId), "获取用户")
}

func (uc *UserController) CacheUser(c *gin.Context) {
	userInfo := &model.UserInfo{UserId: "56867897283718"}
	name, err := redisHelper.Set(ctx, "moose-go", userInfo, 10*time.Minute).Result()
	if err != nil {
		log.Fatal(err)
		common.Failed(c, "缓存用户失败")
		return
	}
	common.Success(c, name, "缓存用户成功")
}

func (uc *UserController) GetCacheUser(c *gin.Context) {
	name, err := redisHelper.Get(ctx, "moose-go").Result()
	if err != nil {
		log.Fatal(err)
		common.Failed(c, "获取缓存用户失败")
		return
	}

	var userInfo model.UserInfo
	json.Unmarshal([]byte(name), &userInfo)
	common.Success(c, userInfo, "获取缓存用户成功")
}

func (uc *UserController) List(c *gin.Context) {
	// userList := make([]model.UserInfo, 0)
	// for i := 0; i < 2; i++ {
	// 	user := model.UserInfo{UserName: "测试用户", Phone: "1537898764", Avatar: "https://gitee.com/shizidada/moose-resource/raw/master/emoji/custom/over.jpeg"}
	// 	userList = append(userList, user)
	// }
	userService := service.UserService{}
	common.Success(c, userService.GetAllUser(), "获取用户列表")
}

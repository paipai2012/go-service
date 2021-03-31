package controller

import (
	"moose-go/common"
	"moose-go/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (uc *UserController) RegisterRouter(app *gin.Engine) {
	group := app.Group("/api/v1/user")
	group.GET("/info", uc.Info)
	group.GET("/get", uc.GetUser)
	group.GET("/list", uc.List)
	group.GET("/cache", uc.CacheUser)
	group.GET("/cache/get", uc.GetCacheUser)
}

func (uc *UserController) Info(c *gin.Context) {
	// 江景 -->
	common.Success(c, "775113183131074580")
}

func (uc *UserController) GetUser(c *gin.Context) {
	userId := c.GetString("userId")
	userService := service.UserService{}
	common.Success(c, userService.GetUserByUserId(userId))
}

func (uc *UserController) CacheUser(c *gin.Context) {
	userService := service.UserService{}
	common.Success(c, userService.CacheUser(c))
}

func (uc *UserController) GetCacheUser(c *gin.Context) {
	userService := service.UserService{}
	common.Success(c, userService.GetCacheUser(c))
}

func (uc *UserController) List(c *gin.Context) {
	// userList := make([]model.UserInfo, 0)
	// for i := 0; i < 2; i++ {
	// 	user := model.UserInfo{UserName: "测试用户", Phone: "1537898764", Avatar: "https://gitee.com/shizidada/moose-resource/raw/master/emoji/custom/over.jpeg"}
	// 	userList = append(userList, user)
	// }
	userService := service.UserService{}
	common.Success(c, userService.GetAllUser())
}

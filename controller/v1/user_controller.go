package v1

import (
	"moose-go/common"
	"moose-go/model"
	"moose-go/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (uc *UserController) RegisterRouter(engine *gin.Engine) {
	group := engine.Group("/api/v1/user")
	// group.GET("/add", uc.Add)
	group.GET("/list", uc.List)
}

func (uc *UserController) Add(c *gin.Context) {
	userName, _ := c.GetQuery("userName")
	userService := service.UserService{}
	row, err := userService.AddUser(userName)
	if err == nil && row > 0 {
		common.Success(c, 1, "添加用户成功")
		return
	}
	common.Failed(c, "添加用户失败")
}

func (uc *UserController) List(c *gin.Context) {
	userList := make([]model.UserInfo, 0)
	for i := 0; i < 2; i++ {
		user := model.UserInfo{UserName: "测试用户", Phone: "1537898764", Avatar: "https://gitee.com/shizidada/moose-resource/raw/master/emoji/custom/over.jpeg"}
		userList = append(userList, user)
	}
	// userService := service.UserService{}
	common.Success(c, userList, "获取用户列表")
}

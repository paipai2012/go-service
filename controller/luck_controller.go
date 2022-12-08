package controller

import (
	"log"
	"sale-service/api"
	"sale-service/common"
	"sale-service/model"
	"sale-service/service"

	"github.com/gin-gonic/gin"
)

type LuckController struct{}

func (lc *LuckController) RegisterRouter(app *gin.Engine) {
	group := app.Group("/api/luck")
	group.POST("/create", lc.addLuck)
	group.GET("/get", lc.getLuck)
	group.POST("/addDraw", lc.addDraw)
}

func (lc *LuckController) addLuck(c *gin.Context) {
	var luck model.Luck
	if err := c.BindJSON(&luck); err != nil {
		log.Println("err:" + err.Error())
		common.FailedParam(c)
		return
	}
	common.JSON(c, service.LuckServiceInstance.AddLuck(&luck))
}

func (lc *LuckController) getLuck(c *gin.Context) {
	var query struct {
		Id int64 `form:"id" binding:"required"`
	}
	if err := c.ShouldBindQuery(&query); err != nil {
		common.JSON(c, api.JsonError(api.GetLuckFailErr).JsonWithMsg(err.Error()))
		return
	}
	common.JSON(c, service.LuckServiceInstance.GetLuck(query.Id))
}

func (lc *LuckController) addDraw(c *gin.Context) {
	var query struct {
		LuckId   int64  `json:"luck_id" binding:"required"`
		Username string `json:"username" binding:"required"`
	}
	if err := c.BindJSON(&query); err != nil {
		common.JSON(c, api.JsonError(api.GetLuckFailErr).JsonWithMsg(err.Error()))
		return
	}
	common.JSON(c, service.LuckServiceInstance.AddDraw(query.LuckId, query.Username))
}

func (lc *LuckController) updateUserPhone(c *gin.Context) {
	var query struct {
		Phone    string `json:"phone" binding:"required"`
		Username string `json:"username" binding:"required"`
	}
	if err := c.BindJSON(&query); err != nil {
		common.JSON(c, api.JsonError(api.GetLuckFailErr).JsonWithMsg(err.Error()))
		return
	}

	common.JSON(c, service.LuckServiceInstance.UpdateUserPhone(query.Username, query.Phone))
}

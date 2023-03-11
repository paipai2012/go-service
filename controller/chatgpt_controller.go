package controller

import (
	"context"
	"log"
	"sale-service/api"
	"sale-service/common"

	"github.com/gin-gonic/gin"
	gogpt "github.com/sashabaranov/go-gpt3"
)

type ChatgptController struct{}

func (cc *ChatgptController) RegisterRouter(app *gin.Engine) {
	group := app.Group("/chatgpt")
	// group.POST("/create", cc.addLuck)
	group.POST("/chat", cc.chat)
	// group.POST("/addDraw", cc.addDraw)
	// group.POST("/updateUserPhone", cc.updateUserPhone)
}

// func (lc *LuckController) addLuck(c *gin.Context) {
// 	var luck model.Luck
// 	if err := c.BindJSON(&luck); err != nil {
// 		log.Println("err:" + err.Error())
// 		common.FailedParam(c)
// 		return
// 	}
// 	common.JSON(c, service.LuckServiceInstance.AddLuck(&luck))
// }

func (cc *ChatgptController) chat(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization")
	var query struct {
		Prompt string `form:"prompt" binding:"required"`
	}
	if err := c.BindJSON(&query); err != nil {
		common.JSON(c, api.JsonError(api.GetLuckFailErr).JsonWithMsg(err.Error()))
		return
	}
	cg := gogpt.NewClient("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")
	req := gogpt.CompletionRequest{
		Model:       gogpt.GPT3TextDavinci003, // 选择的模型
		MaxTokens:   500,
		N:           1,
		Stop:        nil,
		Temperature: 0.5,
		Prompt:      query.Prompt, //要问的问题
	}
	ctx := context.Background()
	resp, err := cg.CreateCompletion(ctx, req)
	if err != nil {
		log.Println("err:" + err.Error())
		common.JSON(c, api.JsonError(api.GetLuckFailErr).JsonWithMsg(err.Error()))
		return
	}
	common.JSON(c, api.JsonSuccess().JsonWithData(resp))
	// common.JSON(c, api.JsonError(api.GetLuckFailErr).JsonWithMsg(resp.Choices[0].Text))
}

// func (lc *LuckController) addDraw(c *gin.Context) {
// 	var query struct {
// 		LuckId   int64  `json:"luck_id" binding:"required"`
// 		Username string `json:"username" binding:"required"`
// 	}
// 	if err := c.BindJSON(&query); err != nil {
// 		common.JSON(c, api.JsonError(api.GetLuckFailErr).JsonWithMsg(err.Error()))
// 		return
// 	}
// 	common.JSON(c, service.LuckServiceInstance.AddDraw(query.LuckId, query.Username))
// }

// func (lc *LuckController) updateUserPhone(c *gin.Context) {
// 	var query struct {
// 		Phone    string `json:"phone" binding:"required"`
// 		Username string `json:"username" binding:"required"`
// 	}
// 	if err := c.BindJSON(&query); err != nil {
// 		common.JSON(c, api.JsonError(api.GetLuckFailErr).JsonWithMsg(err.Error()))
// 		return
// 	}
// 	common.JSON(c, service.LuckServiceInstance.UpdateUserPhone(query.Username, query.Phone))
// }

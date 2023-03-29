package controller

import (
	"context"
	"sale-service/api"
	"sale-service/common"

	"github.com/gin-gonic/gin"
	"github.com/importcjj/sensitive"
	openai "github.com/sashabaranov/go-openai"
)

type chatgptController struct {
	filter *sensitive.Filter
}

var ChatgptController *chatgptController

func init() {
	ChatgptController = &chatgptController{}
	ChatgptController.filter = sensitive.New()
	//filter.LoadNetWordDict("https://raw.githubusercontent.com/importcjj/sensitive/master/dict/dict.txt")
	ChatgptController.filter.LoadWordDict("dict.txt")
}

func (cc *chatgptController) RegisterRouter(app *gin.Engine) {
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

func (cc *chatgptController) chat(c *gin.Context) {
	var query struct {
		Prompt string `form:"prompt" binding:"required"`
	}
	if err := c.BindJSON(&query); err != nil {
		common.JSON(c, api.JsonError(api.GetLuckFailErr).JsonWithMsg(err.Error()))
		return
	}
	res, word := cc.filter.Validate(query.Prompt)
	if !res {
		common.JSON(c, api.JsonError(api.ChatgptFailErr).JsonWithMsg("谨言慎行啊朋友！违禁词："+word))
		return
	}

	client := openai.NewClient("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")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: query.Prompt,
				},
			},
		},
	)

	if err != nil {
		common.JSON(c, api.JsonError(api.GetLuckFailErr).JsonWithMsg(err.Error()))
		return
	}

	common.JSON(c, api.JsonSuccess().JsonWithData(resp))

	// cg := gogpt.NewClient("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")
	// req := gogpt.CompletionRequest{
	// 	Model:       gogpt.GPT4, // 选择的模型
	// 	MaxTokens:   500,
	// 	N:           1,
	// 	Stop:        nil,
	// 	Temperature: 0.5,
	// 	Prompt:      query.Prompt, //要问的问题
	// }
	// ctx := context.Background()
	// resp, err := cg.CreateCompletion(ctx, req)
	// if err != nil {
	// 	common.JSON(c, api.JsonError(api.GetLuckFailErr).JsonWithMsg(err.Error()))
	// 	return
	// }
	// common.JSON(c, api.JsonSuccess().JsonWithData(resp))
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

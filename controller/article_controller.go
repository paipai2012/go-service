package controller

import (
	"log"
	"moose-go/common"
	"moose-go/model"
	"moose-go/service"

	"github.com/gin-gonic/gin"
)

type ArticleController struct{}

func (ac *ArticleController) RegisterRouter(app *gin.Engine) {
	group := app.Group("/api/v1/article")
	group.POST("/create", ac.addArticle)
}

func (sic *ArticleController) addArticle(c *gin.Context) {

	var articleInfo model.ArticleInfo
	if err := c.BindJSON(&articleInfo); err != nil {
		log.Println(err.Error())
		common.FailedParam(c)
		return
	}

	articleService := service.ArticleService{}
	common.JSON(c, articleService.Add(&articleInfo))
}

package service

import (
	"moose-go/api"
	"moose-go/dao"
	"moose-go/engine"
	"moose-go/model"
)

type ArticleService struct{}

func (as *ArticleService) Add(articleInfo *model.ArticleInfo) *api.JsonResult {

	articleDao := dao.ArticleDao{DbEngine: engine.GetOrmEngine()}

	articleInfo.Id = 1
	articleInfo.UserId = 385758996637745152

	result, err := articleDao.InsertArticle(articleInfo)
	if err != nil {
		return api.JsonError(api.AddArticleFailErr).JsonWithMsg(err.Error())
	}

	return api.JsonSuccess().JsonWithData(result)
}

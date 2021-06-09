package dao

import (
	"moose-go/engine"
	"moose-go/model"
)

type ArticleDao struct {
	DbEngine *engine.Orm
}

// 添加用户
func (ud *ArticleDao) InsertArticle(articleInfo *model.ArticleInfo) (int64, error) {
	return ud.DbEngine.InsertOne(articleInfo)
}

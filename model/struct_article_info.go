package model

type ArticleInfo struct {
	Id      int64  `json:"id"`
	UserId  int64  `json:"userId"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

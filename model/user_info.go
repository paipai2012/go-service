package model

type UserInfo struct {
	UserId   int64  `json:"userId"`
	UserName string `json:"userName"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
}

package model

type UserInfo struct {
	UserId      int64  `json:"userId" xorm:"user_id"`
	UserName    string `json:"userName" xorm:"username"`
	AccountId   int64  `json:"accountId" xorm:"account_id"`
	AccountName string `json:"accountName" xorm:"account_name"`
	Phone       string `json:"phone" xorm:"phone"`
	Gender      string `json:"gender" xorm:"gender"`
	Email       string `json:"email" xorm:"email"`
	Address     string `json:"address" xorm:"address"`
	Description string `json:"description" xorm:"description"`
	Avatar      string `json:"avatar" xorm:"avatar"`
	CreateTime  string `json:"createTime" xorm:"<- create_time"`
	UpdateTime  string `json:"updateTime" xorm:"<- update_time"`
}

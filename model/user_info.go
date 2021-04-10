package model

import "encoding/json"

type UserInfo struct {
	UserId      int64  `json:"userId,string" xorm:"user_id"`
	UserName    string `json:"userName" xorm:"username"`
	Phone       string `json:"phone" xorm:"phone"`
	Gender      string `json:"gender" xorm:"gender"`
	Email       string `json:"email" xorm:"email"`
	Address     string `json:"address" xorm:"address"`
	Description string `json:"description" xorm:"description"`
	Avatar      string `json:"avatar" xorm:"avatar"`
	CreateTime  string `json:"createTime" xorm:"<- create_time"`
	UpdateTime  string `json:"updateTime" xorm:"<- update_time"`
}

func (u *UserInfo) NewUser(userId int64, userName, phone, gender, email, address, description string) *UserInfo {
	return &UserInfo{
		UserId:      userId,
		UserName:    userName,
		Phone:       phone,
		Gender:      gender,
		Email:       email,
		Address:     address,
		Description: description,
	}
}

func (u *UserInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

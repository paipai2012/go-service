package model

import "time"

type User struct {
	Id         int64     `json:"id"`
	Username   string    `json:"username" binding:"required"`
	Mobile     string    `json:"mobile"`
	Password   string    `json:"password"`
	Salt       string    `json:"salt"`
	CreateTime time.Time `json:"create_time" xorm:"'create_time' created comment('创建时间') DATETIME"`
	UpdateTime time.Time `json:"update_time" xorm:"'update_time' updated comment('更新时间') DATETIME"`
}

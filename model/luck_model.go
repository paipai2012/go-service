package model

import "time"

type Luck struct {
	Id         int64     `json:"id"`
	Title      string    `json:"title" binding:"required"`
	Status     int8      `json:"status" binding:"required"`
	CreateTime time.Time `json:"create_time" xorm:"'create_time' created comment('创建时间') DATETIME"`
	UpdateTime time.Time `json:"update_time" xorm:"'update_time' updated comment('更新时间') DATETIME"`
}

type LuckItem struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title" binding:"required"`
	LuckId      int64     `json:"luck_id"`
	Img         string    `json:"img"`
	Probability int       `json:"probability" binding:"required"`
	Quantity    int64     `json:"quantity" binding:"required"`
	Status      int8      `json:"status" binding:"required"`
	CreateTime  time.Time `json:"create_time" xorm:"'create_time' created comment('创建时间') DATETIME"`
	UpdateTime  time.Time `json:"update_time" xorm:"'update_time' updated comment('更新时间') DATETIME"`
}

type LuckUser struct {
	Id         int64     `json:"id"`
	Username   string    `json:"username" binding:"required"`
	Mobile     string    `json:"mobile"`
	CreateTime time.Time `json:"create_time" xorm:"'create_time' created comment('创建时间') DATETIME"`
	UpdateTime time.Time `json:"update_time" xorm:"'update_time' updated comment('更新时间') DATETIME"`
}

type LuckRecord struct {
	Id         int64     `json:"id"`
	UserId     int64     `json:"user_id" binding:"required"`
	ItemId     int64     `json:"item_id"`
	LuckId     int64     `json:"luck_id"`
	CreateTime time.Time `json:"create_time" xorm:"'create_time' created comment('创建时间') DATETIME"`
	UpdateTime time.Time `json:"update_time" xorm:"'update_time' updated comment('更新时间') DATETIME"`
}

type LuckResponse struct {
	Id         int64       `json:"id"`
	Title      string      `json:"title" binding:"required"`
	Status     int8        `json:"status" binding:"required"`
	Items      []*LuckItem `json:"items"`
	CreateTime time.Time   `json:"create_time" xorm:"'create_time' created comment('创建时间') DATETIME"`
	UpdateTime time.Time   `json:"update_time" xorm:"'update_time' updated comment('更新时间') DATETIME"`
}

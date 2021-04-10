package model

type Password struct {
	PwdId  int64  `xorm:"pwd_id"`
	UserId int64  `xorm:"user_id"`
	Pwd    string `xorm:"pwd"`
}

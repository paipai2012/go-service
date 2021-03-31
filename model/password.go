package model

type Password struct {
	PwdId  string `xorm:"pwd_id"`
	UserId string `xorm:"user_id"`
	Pwd    string `xorm:"pwd"`
}

package model

type RegisterInfo struct {
	UserName    string `form:"userName" json:"userName" binding:"required"`
	Password    string `form:"password" json:"password" binding:"required"`
	RePassword  string `form:"rePassword" json:"rePassword" binding:"required,eqfield=Password"`
	Phone       string `form:"phone" json:"phone" binding:"required"`
	SmsCode     string `form:"smsCode" json:"smsCode" binding:"required"`
	Gender      string `form:"gender" json:"gender"`
	Avatar      string `form:"avatar" json:"avatar"`
	Email       string `form:"email" json:"email" binding:"omitempty,email"`
	Job         string `form:"job" json:"job"`
	Address     string `form:"address" json:"address"`
	Description string `form:"description" json:"description"`
}

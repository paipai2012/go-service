package model

type LoginInfo struct {
	UserName  string `json:"userName"`
	Password  string `json:"password"`
	Mobile    string `json:"mobile"`
	SmsCode   string `json:"smsCode"`
	LoginType string `json:"loginType" binding:"required,valuein=password sms"`
}

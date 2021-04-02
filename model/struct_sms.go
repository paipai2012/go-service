package model

type Sms struct {
	SmsType string `json:"smsType" binding:"required,valuein=login"`
	Mobile  string `json:"mobile" binding:"required"`
}

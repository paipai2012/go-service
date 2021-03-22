package param

type RegisterInfoParam struct {
	UserName   string `json:"userName"`
	Password   string `json:"password"`
	RePassword string `json:"rePassword"`
	Phone      string `json:"phone"`
	SmsCode    string `json:"smsCode"`
}

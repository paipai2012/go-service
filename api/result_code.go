package api

const (
	AddUserFailCode = 10001

	QRCodeRetryCode   = 20001
	QRCodeGetFailCode = 20002

	PasswordInconsistencyCode = 30001

	ParseParamCode = -10001
)

var resultCodeText = map[int]string{
	AddUserFailCode: "添加用户失败",

	QRCodeRetryCode:   "重新获取二维码",
	QRCodeGetFailCode: "获取二维码失败",

	PasswordInconsistencyCode: "密码不一致",

	ParseParamCode: "参数解析错误",
}

func StatusText(code int) (string, bool) {
	message, ok := resultCodeText[code]
	return message, ok
}

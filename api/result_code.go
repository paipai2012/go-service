package api

const (
	AddUserFail = 10001

	QRCodeRetry   = 20001
	QRCodeGetFail = 30002
)

var resultCodeText = map[int]string{
	AddUserFail: "添加用户失败",

	QRCodeRetry:   "重新获取二维码",
	QRCodeGetFail: "获取二维码失败",
}

func StatusText(code int) (string, bool) {
	message, ok := resultCodeText[code]
	return message, ok
}

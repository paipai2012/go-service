package api

var (
	OK = NewException(200, "请求成功")

	// 服务级错误码
	ErrServer    = NewException(10001, "服务异常，请联系管理员")
	ErrParam     = NewException(10002, "参数有误")
	ErrSignParam = NewException(10003, "签名参数有误")

	AddUserFailErr   = NewException(20001, "添加用户失败")
	PhoneExistErr    = NewException(20002, "手机号码已存在")
	UserNameExistErr = NewException(20003, "用户名已存在")
	QueryUserFailErr = NewException(20004, "获取用户信息失败")

	UserNameEmptyErr = NewException(20005, "用户名不能为空")
	PasswordEmptyErr = NewException(20006, "密码不能为空")

	PhoneEmptyErr   = NewException(20007, "手机号码不能为空")
	SmdCodeEmptyErr = NewException(20008, "短信验证码不能为空")

	UserNameOrPasswordErr = NewException(20009, "账号或密码错误")
	LoginTypeErr          = NewException(20010, "登录方式错误")

	QRCodeRetryErr   = NewException(30001, "重新获取二维码")
	QRCodeGetFailErr = NewException(30002, "获取二维码失败")

	PasswordErr = NewException(30001, "两次密码不一致")
)

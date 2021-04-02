package api

var (
	OK = NewException(200, "请求成功")

	// 服务级错误码
	ErrServer    = NewException(1001, "服务异常，请联系管理员")
	ErrParam     = NewException(1002, "参数有误")
	ErrSignParam = NewException(1003, "签名参数有误")

	// 业务异常
	UserAddFailErr        = NewException(20001, "添加用户失败")
	UserNameExistErr      = NewException(20002, "用户名已存在")
	UserNameOrPasswordErr = NewException(20003, "账号或密码错误")
	UserNameEmptyErr      = NewException(20004, "用户名不能为空")

	PhoneExistErr  = NewException(20101, "手机号码已存在")
	PhoneEmptyErr  = NewException(20102, "手机号码不能为空")
	PhoneNumberErr = NewException(20103, "手机号码错误")

	SmdCodeEmptyErr = NewException(20201, "短信验证码不能为空")
	SmdCodeSendErr  = NewException(20202, "短信验证码发送失败")
	SmdCodeErr      = NewException(20203, "短信验证码错误")

	PasswordEmptyErr = NewException(20302, "密码不能为空")
	PasswordErr      = NewException(20403, "两次密码不一致")
	LoginTypeErr     = NewException(20504, "登录方式错误")

	QueryUserFailErr = NewException(30001, "获取用户信息失败")
	QRCodeRetryErr   = NewException(30002, "重新获取二维码")
	QRCodeGetFailErr = NewException(30003, "获取二维码失败")

	JwtValidationErr = NewException(90001, "令牌验证失败")
	JwtExpiresErr    = NewException(90002, "无效令牌")
	JwtGeneratorErr  = NewException(90003, "生成令牌失败")
)

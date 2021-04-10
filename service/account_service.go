package service

import (
	"log"
	"moose-go/api"
	"moose-go/dao"
	"moose-go/engine"
	"moose-go/model"
	"moose-go/util"
	"strconv"
	"strings"
)

type AccountService struct{}

func (us *AccountService) Register(registerInfo *model.RegisterInfo) *api.JsonResult {
	log.Printf("receiver param %s", registerInfo)

	password := registerInfo.Password

	worker, err := util.NewWorker(16)
	if err != nil {
		log.Print(err)
		return api.JsonError(api.UserAddFailErr)
	}

	userDao := dao.UserDao{DbEngine: engine.GetOrmEngine()}
	exist, err := userDao.QueryByPhone(registerInfo.Phone)
	log.Printf("check Phone %v %v", exist, err)
	if exist || err != nil {
		return api.JsonError(api.PhoneExistErr)
	}

	exist, err = userDao.QueryByUserName(registerInfo.UserName)
	log.Printf("check UserName %v %v", exist, err)
	if exist || err != nil {
		return api.JsonError(api.UserNameExistErr)
	}

	userInfo := &model.UserInfo{
		UserName:    registerInfo.UserName,
		UserId:      worker.GetId(),
		Gender:      registerInfo.Gender,
		Phone:       registerInfo.Phone,
		Avatar:      registerInfo.Avatar,
		Email:       registerInfo.Email,
		Address:     registerInfo.Address,
		Description: registerInfo.Description,
	}

	result, err := userDao.InsertUser(userInfo)
	log.Printf("add user result %v error %v", result, err)
	if err != nil {
		return api.JsonError(api.UserAddFailErr)
	}

	passwordInfo := &model.Password{
		PwdId:  worker.GetId(),
		UserId: userInfo.UserId,
		Pwd:    util.CryptoSha1(password),
	}
	result, err = userDao.InsertPassword(passwordInfo)
	log.Printf("add password result %v error %v", result, err)
	if err != nil {
		return api.JsonError(api.UserAddFailErr)
	}
	return api.JsonSuccess()
}

func (as *AccountService) Login(loginInfo *model.LoginInfo) *api.JsonResult {
	loginType := loginInfo.LoginType
	if loginType == "password" {
		return loginWithPassword(loginInfo)
	} else if loginType == "sms" {
		return loginWithSmsCode(loginInfo)
	}
	return api.JsonError(api.LoginTypeErr)
}

// login width sms code
func loginWithSmsCode(loginInfo *model.LoginInfo) *api.JsonResult {

	mobile := loginInfo.Mobile
	if mobile == "" {
		return api.JsonError(api.PhoneEmptyErr)
	}

	smsCode := loginInfo.SmsCode
	if smsCode == "" {
		return api.JsonError(api.SmsCodeEmptyErr)
	}

	userDao := dao.UserDao{DbEngine: engine.GetOrmEngine()}
	userResult, userErr := userDao.QueryUserIdByPhone(mobile)
	if userErr != nil {
		return api.JsonError(api.PhoneNumberErr)
	}
	if len(userResult) == 0 {
		log.Println(userResult)
		return api.JsonError(api.PhoneNumberErr)
	}

	smsService := &SenderMessageService{}
	sms := &model.Sms{SmsType: "login", Mobile: mobile}
	saveSmsCode := smsService.CheckSms(sms)

	if !strings.EqualFold(smsCode, saveSmsCode) {
		return api.JsonError(api.SmsCodeErr)
	}

	userId, _ := strconv.ParseInt(string(userResult[0]["user_id"]), 10, 64)
	token, err := createToken(userId)
	if err != nil {
		return api.JsonError(api.JwtGeneratorErr)
	}
	return api.JsonData(token)
}

// login with password
func loginWithPassword(loginInfo *model.LoginInfo) *api.JsonResult {

	userName := loginInfo.UserName
	if userName == "" {
		return api.JsonError(api.UserNameEmptyErr)
	}

	password := loginInfo.Password
	if password == "" {
		return api.JsonError(api.PasswordEmptyErr)
	}

	userDao := dao.UserDao{DbEngine: engine.GetOrmEngine()}
	userResult, userErr := userDao.QueryUserIdByUserName(loginInfo.UserName)
	if userErr != nil || len(userResult) == 0 {
		log.Println(userErr)
		return api.JsonError(api.UserNameOrPasswordErr)
	}

	userId, err := strconv.ParseInt(string(userResult[0]["user_id"]), 10, 64)
	if userId == 0 || err != nil {
		log.Println("user id not exist")
		return api.JsonError(api.UserNameOrPasswordErr)
	}

	passwordDao := dao.PasswordDao{DbEngine: engine.GetOrmEngine()}
	pwdResult, pwdErr := passwordDao.QueryPasswordByUserId(userId)
	if pwdErr != nil || len(pwdResult) == 0 {
		log.Println(pwdErr)
		return api.JsonError(api.UserNameOrPasswordErr)
	}

	pwd := string(pwdResult[0]["pwd"])
	encodePwd := util.CryptoSha1(loginInfo.Password)
	if !strings.EqualFold(pwd, encodePwd) {
		// log.Println(pwd, " \n ", encodePwd)
		return api.JsonError(api.UserNameOrPasswordErr)
	}
	token, err := createToken(userId)
	if err != nil {
		return api.JsonError(api.JwtGeneratorErr)
	}
	return api.JsonData(token)
}

// cerate jwt token
func createToken(userId int64) (string, error) {
	return util.GeneratorJwt(&model.Payload{UserId: strconv.FormatInt(userId, 10)})
}

package service

import (
	"log"
	"moose-go/api"
	"moose-go/dao"
	"moose-go/engine"
	"moose-go/model"
	"moose-go/service"
	"moose-go/util"
	"strconv"
	"strings"
)

type AccountService struct{}

func (us *AccountService) Register(registerInfo *model.RegisterInfo) {
	password := registerInfo.Password

	worker, err := util.NewWorker(16)
	if err != nil {
		log.Print(err)
		panic(api.UserAddFailErr)
	}

	// check user info
	checkRegisterInfo(registerInfo)

	userInfo := &model.UserInfo{
		UserName:    registerInfo.UserName,
		UserId:      strconv.FormatInt(worker.GetId(), 10),
		Gender:      registerInfo.Gender,
		Phone:       registerInfo.Phone,
		Avatar:      registerInfo.Avatar,
		Email:       registerInfo.Email,
		Address:     registerInfo.Address,
		Description: registerInfo.Description,
	}

	userDao := dao.UserDao{DbEngine: engine.GetOrmEngine()}
	result, err := userDao.InsertUser(userInfo)
	log.Printf("add user result %v error %v", result, err)
	if err != nil {
		panic(api.UserAddFailErr.WithErrMsg(err))
	}

	passwordInfo := &model.Password{
		PwdId:  strconv.FormatInt(worker.GetId(), 10),
		UserId: userInfo.UserId,
		Pwd:    util.CryptoSha1(password),
	}
	result, err = userDao.InsertPassword(passwordInfo)
	log.Printf("add password result %v error %v", result, err)
	if err != nil {
		panic(api.UserAddFailErr)
	}
}

func (as *AccountService) Login(loginInfo *model.LoginInfo) string {

	loginType := loginInfo.LoginType
	if loginType == "password" {
		// check user info
		checkLoginInfo(loginInfo)
		return loginWithPassword(loginInfo)
	} else if loginType == "sms" {
		return loginWithSmsCode(loginInfo)
	}
	panic(api.LoginTypeErr)
}

// login width sms code
func loginWithSmsCode(loginInfo *model.LoginInfo) string {

	mobile := loginInfo.Mobile

	userDao := dao.UserDao{DbEngine: engine.GetOrmEngine()}
	exist, err := userDao.QueryByPhone(loginInfo.Mobile)
	if err != nil || !exist {
		panic(api.PhoneNumberErr)
	}

	smsCode := loginInfo.SmsCode
	loginType := loginInfo.LoginType

	smsService := service.SenderMessageService{}
	smsService.CheckSms(loginType, smsCode, mobile)

	return ""
}

// login with password
func loginWithPassword(loginInfo *model.LoginInfo) string {

	userDao := dao.UserDao{DbEngine: engine.GetOrmEngine()}
	userResult, userErr := userDao.QueryUserIdByUserName(loginInfo.UserName)
	if userErr != nil {
		log.Println(userErr)
		panic(api.UserNameOrPasswordErr.WithErrMsg(userErr.Error()))
	}
	if len(userResult) == 0 {
		log.Println(userResult)
		panic(api.UserNameOrPasswordErr)
	}

	userId := string(userResult[0]["user_id"])
	if userId == "" {
		log.Println("user id not exist")
		panic(api.UserNameOrPasswordErr)
	}
	passwordDao := dao.PasswordDao{DbEngine: engine.GetOrmEngine()}
	pwdResult, pwdErr := passwordDao.QueryPasswordByUserId(userId)
	if pwdErr != nil {
		log.Println(pwdErr)
		panic(api.UserNameOrPasswordErr.WithErrMsg(pwdErr.Error()))
	}
	if len(pwdResult) == 0 {
		log.Println(pwdResult)
		panic(api.UserNameOrPasswordErr)
	}
	pwd := string(pwdResult[0]["pwd"])
	encodePwd := util.CryptoSha1(loginInfo.Password)
	if !strings.EqualFold(pwd, encodePwd) {
		// log.Println(pwd, " \n ", encodePwd)
		panic(api.UserNameOrPasswordErr)
	}
	return createToken(&model.UserInfo{UserId: userId})
}

// cerate jwt tokn
func createToken(userInfo *model.UserInfo) string {
	token, err := util.GeneratorJwt(userInfo)
	if err != nil {
		panic(api.JwtGeneratorErr)
	}
	return token
}

// check login info
func checkLoginInfo(loginInfo *model.LoginInfo) {
	loginType := loginInfo.LoginType

	if loginType == "password" {
		userName := loginInfo.UserName
		if userName == "" {
			panic(api.UserNameEmptyErr)
		}
		password := loginInfo.Password
		if password == "" {
			panic(api.PasswordEmptyErr)
		}
	}
	if loginType == "sms_code" {
		mobile := loginInfo.Mobile
		if mobile == "" {
			panic(api.PhoneEmptyErr)
		}
		smsCode := loginInfo.SmsCode
		if smsCode == "" {
			panic(api.SmdCodeEmptyErr)
		}
	}
}

// check register info
func checkRegisterInfo(registerInfo *model.RegisterInfo) {
	userDao := dao.UserDao{DbEngine: engine.GetOrmEngine()}
	exist, err := userDao.QueryByPhone(registerInfo.Phone)
	log.Printf("check Phone %v %v", exist, err)
	if exist {
		panic(api.PhoneExistErr)
	}
	exist, err = userDao.QueryByUserName(registerInfo.UserName)
	log.Printf("check UserName %v %v", exist, err)
	if exist {
		panic(api.UserNameExistErr)
	}
}

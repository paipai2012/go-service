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

func (us *AccountService) Register(registerInfo *model.RegisterInfo) {
	password := registerInfo.Password

	worker, err := util.NewWorker(16)
	if err != nil {
		log.Print(err)
		panic(api.AddUserFailErr)
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
		panic(api.AddUserFailErr.WithErrMsg(err))
	}

	passwordInfo := &model.Password{
		PwdId:  strconv.FormatInt(worker.GetId(), 10),
		UserId: userInfo.UserId,
		Pwd:    util.CryptoSha1(password),
	}
	result, err = userDao.InsertPassword(passwordInfo)
	log.Printf("add password result %v error %v", result, err)
	if err != nil {
		panic(api.AddUserFailErr)
	}
}

func (as *AccountService) Login(loginInfo *model.LoginInfo) {

	loginType := loginInfo.LoginType
	if loginType == "password" {
		// check user info
		checkLoginInfo(loginInfo)
		loginWithPassword(loginInfo)
		return
	} else if loginType == "sms_code" {
		loginWithSmsCode(loginInfo)
		return
	}
	panic(api.LoginTypeErr)
}

// login width sms code
func loginWithSmsCode(loginInfo *model.LoginInfo) {

}

// login with password
func loginWithPassword(loginInfo *model.LoginInfo) {

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
}

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

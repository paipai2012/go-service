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

func (us *AccountService) AddUser(registerInfo *model.RegisterInfo) {
	password := registerInfo.Password
	repassword := registerInfo.RePassword

	if !strings.EqualFold(password, repassword) {
		api.NewException(api.PasswordInconsistencyCode)
	}

	worker, err := util.NewWorker(16)
	if err != nil {
		log.Print(err)
		api.NewException(api.AddUserFailCode)
	}

	userDao := dao.UserDao{DbEngine: engine.GetOrmEngine()}

	// check user info
	checkUserInfo(registerInfo)

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
	result, err := userDao.InsertUser(userInfo)
	log.Printf("add user result %d error %s", result, err)
	if err != nil {
		panic(api.NewException(api.AddUserFailCode))
	}

	passwordInfo := &model.Password{
		PwdId:  strconv.FormatInt(worker.GetId(), 10),
		UserId: userInfo.UserId,
		Pwd:    util.CryptoSha1(password),
	}
	result, err = userDao.InsertPassword(passwordInfo)
	log.Printf("add password result %d error %s", result, err)
	if err != nil {
		panic(api.NewException(api.AddUserFailCode))
	}
}

func checkUserInfo(registerInfo *model.RegisterInfo) {
	userDao := dao.UserDao{DbEngine: engine.GetOrmEngine()}
	exist, err := userDao.QueryByPhone(registerInfo.Phone)
	log.Printf("check Phone %v %v", exist, err)
	if exist {
		panic(api.NewException(api.PhoneExistCode))
	}

	exist, err = userDao.QueryByUserName(registerInfo.UserName)
	log.Printf("check UserName %v %v", exist, err)
	if exist {
		panic(api.NewException(api.UserNameExistCode))
	}
}

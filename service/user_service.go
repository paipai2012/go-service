package service

import (
	"bytes"
	"encoding/binary"
	"moose-go/dao"
	"moose-go/engine"
	"moose-go/model"
)

type UserService struct {
}

func (us *UserService) AddUser(userName string) (int64, error) {
	userInfo := model.UserInfo{
		UserName:    userName,
		UserId:      1,
		AccountId:   1,
		AccountName: "JiangJing",
		Gender:      "1",
		Phone:       "15798980298",
		Avatar:      "https://www.gitee.com/shizidada",
		Email:       "jiangjing@163,com",
		Address:     "中国",
		Description: "我是江景啊",
	}
	userDao := dao.UserDao{DbEngine: engine.GetOrmEngine()}
	return userDao.InsertUser(&userInfo)
}

func (us *UserService) GetUserByUserId(userId int64) *model.UserInfo {
	userDao := dao.UserDao{DbEngine: engine.GetOrmEngine()}
	return userDao.QueryByUserId(userId)
}

func (us *UserService) GetAllUser() []*model.UserInfo {
	userDao := dao.UserDao{DbEngine: engine.GetOrmEngine()}
	rows := userDao.QueryUserList()
	// []map[string][]byte
	list := make([]*model.UserInfo, len(rows))
	for index, value := range rows {
		// []byte
		UserId := bytesToInt(value["user_id"])
		UserName := string(value["username"])
		Phone := string(value["phone"])
		Avatar := string(value["avatar"])
		userInfo := model.UserInfo{UserId: UserId, UserName: UserName, Phone: Phone, Avatar: Avatar}
		list[index] = &userInfo
		// list = append(list, &userInfo)
	}
	return list
}

func bytesToInt(bys []byte) int64 {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int64(data)
}

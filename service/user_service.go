package service

import (
	"moose-go/dao"
	"moose-go/engine"
	"moose-go/model"
)

type UserService struct {
}

func (us *UserService) AddUser(userName string) (int64, error) {
	userDao := dao.UserDao{DbEngine: engine.DbEngine}
	return userDao.InsertUser(userName)
}

func (us *UserService) FindAllUser() []*model.UserInfo {
	userDao := dao.UserDao{DbEngine: engine.DbEngine}
	return userDao.QueryUserList()
}

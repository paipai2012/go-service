package dao

import (
	"fmt"
	"log"
	"moose-go/engine"
	"moose-go/model"
)

type UserDao struct {
	DbEngine *engine.Orm
}

// 添加用户
func (ud *UserDao) InsertUser(userInfo *model.UserInfo) (int64, error) {
	result, err := ud.DbEngine.InsertOne(userInfo)
	fmt.Println(result, err)
	if err != nil {
		return 0, nil
	}
	return result, err
}

func (ud *UserDao) QueryByUserId(userId int64) *model.UserInfo {
	userInfo := model.UserInfo{
		UserId: userId,
	}
	has, err := ud.DbEngine.Get(&userInfo)
	if err != nil {
		log.Panicln(err)
		return nil
	}
	fmt.Println(has)
	return &userInfo
}

// 查询用户列表
func (ud *UserDao) QueryUserList() []map[string][]byte {
	rows, err := ud.DbEngine.Query("select * from t_user_info limit 0, 10")
	if err != nil {
		return nil
	}
	return rows
}

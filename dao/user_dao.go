package dao

import (
	"sale-service/engine"
	"sale-service/model"
)

type UserDao struct {
	DbEngine *engine.Orm
}

// 添加用户
func (ud *UserDao) InsertUser(userInfo *model.UserInfo) (int64, error) {
	return ud.DbEngine.InsertOne(userInfo)
}

func (ud *UserDao) InsertPassword(password *model.Password) (int64, error) {
	return ud.DbEngine.InsertOne(password)
}

func (ud *UserDao) QueryUserIdByUserName(userName string) ([]map[string][]byte, error) {
	sql := "select user_id from t_user_info where username = ? "
	return ud.DbEngine.Query(sql, userName)
}

func (ud *UserDao) QueryUserIdByPhone(phone string) ([]map[string][]byte, error) {
	sql := "select user_id from t_user_info where phone = ? "
	return ud.DbEngine.Query(sql, phone)
}

func (ud *UserDao) QueryByUserId(userId int64) ([]map[string][]byte, error) {
	sql := "select user_id, username, phone, gender, avatar, email, job, address, description from t_user_info where user_id = ? "
	return ud.DbEngine.Query(sql, userId)
}

func (ud *UserDao) QueryByPhone(phone string) (bool, error) {
	userInfo := model.UserInfo{
		Phone: phone,
	}
	return ud.DbEngine.Get(&userInfo)
}

func (ud *UserDao) QueryByUserName(userName string) (bool, error) {
	userInfo := model.UserInfo{
		UserName: userName,
	}
	return ud.DbEngine.Get(&userInfo)
}

// 查询用户列表
func (ud *UserDao) QueryUserList() ([]map[string][]byte, error) {
	return ud.DbEngine.Query("select * from t_user_info limit 0, 10")
}

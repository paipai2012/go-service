package dao

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"moose-go/engine"
	"moose-go/model"
)

type UserDao struct {
	DbEngine *engine.Orm
}

// 添加用户
func (ud *UserDao) InsertUser(userName string) (int64, error) {
	userInfo := model.UserInfo{UserName: userName, UserId: 1231371387187, Phone: "15798980298", Avatar: "https://www.gitee.com/shizidada"}
	result, err := ud.DbEngine.InsertOne(userInfo)
	fmt.Println(result, err)
	if err != nil {
		return 0, nil
	}
	return result, err
}

// 查询用户列表
func (ud *UserDao) QueryUserList() []*model.UserInfo {

	rows, err := ud.DbEngine.Query("select * from t_user_info limit 0, 10")
	if err != nil {
		return nil
	}

	// []map[string][]byte
	list := make([]*model.UserInfo, len(rows))
	for index, value := range rows {
		// []byte
		UserId := _BytesToInt(value["user_id"])
		UserName := string(value["username"])
		Phone := string(value["phone"])
		Avatar := string(value["avatar"])
		userInfo := model.UserInfo{UserId: UserId, UserName: UserName, Phone: Phone, Avatar: Avatar}
		list[index] = &userInfo
		// list = append(list, &userInfo)
	}
	return list
}

func _BytesToInt(bys []byte) int64 {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int64(data)
}

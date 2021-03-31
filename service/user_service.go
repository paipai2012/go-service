package service

import (
	"encoding/json"
	"log"
	"moose-go/api"
	"moose-go/dao"
	"moose-go/engine"
	"moose-go/model"
	"time"

	"github.com/gin-gonic/gin"
)

type UserService struct {
}

func (us *UserService) GetUserByUserId(userId string) []map[string][]byte {
	userDao := dao.UserDao{DbEngine: engine.GetOrmEngine()}
	result, err := userDao.QueryByUserId(userId)
	if err != nil {
		panic(api.NewException(api.QueryUserFail))
	}
	return result
}

func (us *UserService) GetAllUser() []*model.UserInfo {
	userDao := dao.UserDao{DbEngine: engine.GetOrmEngine()}
	rows := userDao.QueryUserList()
	// []map[string][]byte
	list := make([]*model.UserInfo, len(rows))
	for index, value := range rows {
		// []byte
		UserId := string(value["user_id"])
		UserName := string(value["username"])
		Phone := string(value["phone"])
		Avatar := string(value["avatar"])
		userInfo := model.UserInfo{UserId: UserId, UserName: UserName, Phone: Phone, Avatar: Avatar}
		list[index] = &userInfo
		// list = append(list, &userInfo)
	}
	return list
}

func (uc *UserService) CacheUser(c *gin.Context) string {
	redisHelper := engine.GetRedisHelper()
	userInfo := &model.UserInfo{UserId: "56867897283718"}
	name, err := redisHelper.Set(ctx, "moose-go", userInfo, 10*time.Minute).Result()
	if err != nil {
		log.Panic(err)
	}
	return name
}

func (uc *UserService) GetCacheUser(c *gin.Context) *model.UserInfo {
	redisHelper := engine.GetRedisHelper()
	name, err := redisHelper.Get(ctx, "moose-go").Result()
	if err != nil {
		log.Panic(err)
		return nil
	}

	var userInfo model.UserInfo
	json.Unmarshal([]byte(name), &userInfo)
	return &userInfo
}

// func bytesToInt(bys []byte) int64 {
// 	bytebuff := bytes.NewBuffer(bys)
// 	var data int64
// 	binary.Read(bytebuff, binary.BigEndian, &data)
// 	return int64(data)
// }

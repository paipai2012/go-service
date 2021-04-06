package service

import (
	"encoding/json"
	"fmt"
	"log"
	"moose-go/api"
	"moose-go/dao"
	"moose-go/engine"
	"moose-go/model"
	"moose-go/util"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserService struct {
}

func (us *UserService) GetUserByUserId(userId string) []map[string][]byte {
	userDao := dao.UserDao{DbEngine: engine.GetOrmEngine()}
	result, err := userDao.QueryByUserId(userId)
	if err != nil {
		panic(api.QueryUserFailErr)
	}
	return result
}

func (us *UserService) GetAllUser() []*model.UserInfo {
	userDao := dao.UserDao{DbEngine: engine.GetOrmEngine()}
	rows, err := userDao.QueryUserList()
	if err != nil {
		panic(api.QueryUserFailErr)
	}
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
	redisHelper := engine.GetRedisEngine()
	userInfo := &model.UserInfo{UserId: "56867897283718"}
	name, err := redisHelper.Set(ctx, "moose-go", userInfo, 10*time.Minute).Result()
	if err != nil {
		log.Panic(err)
	}
	return name
}

func (uc *UserService) GetCacheUser(c *gin.Context) *model.UserInfo {
	redisHelper := engine.GetRedisEngine()
	name, err := redisHelper.Get(ctx, "moose-go").Result()
	if err != nil {
		log.Panic(err)
		return nil
	}

	var userInfo model.UserInfo
	json.Unmarshal([]byte(name), &userInfo)
	return &userInfo
}

func (uc *UserService) GetUserByToken(header string) *model.UserInfo {
	token := util.ParseBearerToken(header)
	jwtToken := util.ParseJwt(token)
	data, err := json.Marshal(jwtToken.Claims)
	if err != nil {
		panic(api.QueryUserFailErr)
	}

	var claims jwt.MapClaims
	err = json.Unmarshal(data, &claims)
	if err != nil {
		panic(api.QueryUserFailErr)
	}

	userId, ok := claims["userId"]
	if userId == "" || !ok {
		panic(api.QueryUserFailErr)
	}

	result := uc.GetUserByUserId(fmt.Sprintf("%s", userId))
	userInfo := &model.UserInfo{
		UserId:      string(result[0]["user_id"]),
		UserName:    string(result[0]["username"]),
		Phone:       string(result[0]["phone"]),
		Avatar:      string(result[0]["avatar"]),
		Gender:      string(result[0]["gender"]),
		Address:     string(result[0]["address"]),
		Email:       string(result[0]["email"]),
		Description: string(result[0]["description"]),
	}
	return userInfo
}

// func bytesToInt(bys []byte) int64 {
// 	bytebuff := bytes.NewBuffer(bys)
// 	var data int64
// 	binary.Read(bytebuff, binary.BigEndian, &data)
// 	return int64(data)
// }

package service

import (
	"encoding/json"
	"fmt"
	"log"
	"sale-service/api"
	"sale-service/dao"
	"sale-service/engine"
	"sale-service/model"
	"sale-service/util"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

type UserService struct {
}

func (us *UserService) GetAllUser() *api.JsonResult {
	userDao := dao.UserDao{DbEngine: engine.GetOrmEngine()}
	rows, err := userDao.QueryUserList()
	if err != nil {
		return api.JsonError(api.QueryUserFailErr)
	}
	// []map[string][]byte
	list := make([]*model.UserInfo, len(rows))
	for index, value := range rows {
		// []byte
		UserId, _ := strconv.ParseInt(string(value["user_id"]), 10, 64)
		userInfo := model.UserInfo{
			UserId:      UserId,
			UserName:    string(value["username"]),
			Phone:       string(value["phone"]),
			Avatar:      string(value["avatar"]),
			Gender:      string(value["gender"]),
			Address:     string(value["address"]),
			Email:       string(value["email"]),
			Description: string(value["description"]),
		}
		list[index] = &userInfo
	}
	return api.JsonData(list)
}

func (uc *UserService) GetUserWithToken(header string) *api.JsonResult {
	jsonResult := uc.getCurrentUserId(header)
	if jsonResult.Code != 200 {
		return jsonResult
	}

	userId, _ := strconv.ParseInt(fmt.Sprintf("%s", jsonResult.Data), 10, 64)
	result, err := uc.getUserWithUserId(userId)
	if err != nil {
		return api.JsonError(api.QueryUserFailErr)
	}
	userInfo := &model.UserInfo{
		UserId:      userId,
		UserName:    string(result[0]["username"]),
		Phone:       string(result[0]["phone"]),
		Avatar:      string(result[0]["avatar"]),
		Gender:      string(result[0]["gender"]),
		Address:     string(result[0]["address"]),
		Email:       string(result[0]["email"]),
		Description: string(result[0]["description"]),
	}
	return api.JsonData(userInfo)
}

func (us *UserService) getUserWithUserId(userId int64) ([]map[string][]byte, error) {
	userDao := dao.UserDao{DbEngine: engine.GetOrmEngine()}
	return userDao.QueryByUserId(userId)
}

func (uc *UserService) getCurrentUserId(header string) *api.JsonResult {
	token := util.ParseBearerToken(header)
	if token == "" {
		return api.JsonError(api.JwtExpiresErr)
	}

	jwtToken := util.ParseJwt(token)
	if jwtToken == nil {
		return api.JsonError(api.JwtExpiresErr)
	}

	data, err := json.Marshal(jwtToken.Claims)
	if err != nil {
		return api.JsonError(api.JwtExpiresErr)
	}

	var claims jwt.MapClaims
	err = json.Unmarshal(data, &claims)
	if err != nil {
		return api.JsonError(api.JwtExpiresErr)
	}

	id, ok := claims["userId"]
	log.Printf("%v", claims)
	if id == "" || !ok {
		return api.JsonError(api.JwtExpiresErr)
	}
	return api.JsonData(id)
}

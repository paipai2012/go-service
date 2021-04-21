package test

import (
	"moose-go/engine"
	"moose-go/model"
	"testing"
	// "github.com/mitchellh/mapstructure"
)

func TestExample(t *testing.T) {

	engine.NewOrmEngine()

	dbEngine := engine.GetOrmEngine()
	sql := "select user_id, username, phone, gender, avatar, email, job, address, description from t_user_info where user_id = ? "
	result, _ := dbEngine.Query(sql, "385758996637745152")

	// result
	// []map[string][]byte

	// result[0] ==> map[string][]byte
	var userInfo model.UserInfo
	// mapstructure.Decode(result[0], &userInfo)
	t.Log(result[0])
	t.Log("--------------------------------------")
	t.Logf("%#v", userInfo)
}

package dao

import (
	"moose-go/engine"
)

type PasswordDao struct {
	DbEngine *engine.Orm
}

// query password by userid
func (pd *PasswordDao) QueryPasswordByUserId(userId int64) ([]map[string][]byte, error) {
	sql := "select pwd from t_password where user_id = ?"
	return pd.DbEngine.Query(sql, userId)
}

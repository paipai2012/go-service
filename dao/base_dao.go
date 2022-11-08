package dao

import (
	"sale-service/engine"
)

type BaseDao struct {
	Db *engine.Orm
}

var MyBaseDao *BaseDao

func init() {
	MyBaseDao = &BaseDao{}
	MyBaseDao.Db = engine.GetOrmEngine()
}

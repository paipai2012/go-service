package engine

import (
	"fmt"
	"moose-go/model"

	_ "github.com/go-sql-driver/mysql"

	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var _dbEngine *Orm

type Orm struct {
	*xorm.Engine
}

func GetOrmEngine() *Orm {
	return _dbEngine
}

func NewOrmEngine(appInfo *model.AppInfo) (*xorm.Engine, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", appInfo.UserName, appInfo.Password, appInfo.Host, appInfo.Port, appInfo.DataBase)
	engine, err := xorm.NewEngine(appInfo.DriverName, url)

	fmt.Println(appInfo)

	if err != nil {
		return nil, err
	}

	prefix := names.NewPrefixMapper(names.SnakeMapper{}, "t_")
	engine.SetTableMapper(prefix)

	engine.ShowSQL(true)

	// 创建表
	err = engine.Sync2(new(model.UserInfo))
	if err != nil {
		return nil, err
	}

	orm := new(Orm)
	orm.Engine = engine
	_dbEngine = orm

	return engine, nil
}

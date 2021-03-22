package engine

import (
	"fmt"
	"moose-go/model"

	_ "github.com/go-sql-driver/mysql"

	"xorm.io/xorm"
)

var DbEngine *Orm

type Orm struct {
	*xorm.Engine
}

func OrmEngine(appInfo *model.AppInfo) (*xorm.Engine, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", appInfo.UserName, appInfo.Password, appInfo.Host, appInfo.Port, appInfo.DataBase)
	engine, err := xorm.NewEngine(appInfo.DriverName, url)

	fmt.Println(appInfo)

	if err != nil {
		return nil, err
	}

	// 创建表
	// Sync2 synchronize structs to database tables
	err = engine.Sync2(new(model.UserInfo))
	if err != nil {
		return nil, err
	}

	orm := new(Orm)
	orm.Engine = engine
	DbEngine = orm

	return engine, nil
}

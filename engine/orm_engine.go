package engine

import (
	"fmt"
	"log"
	"moose-go/model"
	"sync"

	_ "github.com/go-sql-driver/mysql"

	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

type Orm struct {
	*xorm.Engine
}

var dbEngine *Orm

var ormOnce sync.Once

func GetOrmEngine() *Orm {
	return dbEngine
}

func NewOrmEngine(appInfo *model.AppInfo) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", appInfo.UserName, appInfo.Password, appInfo.Host, appInfo.Port, appInfo.DataBase)
	engine, err := xorm.NewEngine(appInfo.DriverName, url)

	if err != nil {
		log.Panic(err.Error())
		return
	}

	prefix := names.NewPrefixMapper(names.SnakeMapper{}, "t_")
	engine.SetTableMapper(prefix)

	engine.ShowSQL(true)

	// 创建表
	// err = engine.Sync2(new(model.UserInfo))
	// if err != nil {
	// 	return nil, err
	// }

	// 单利模式
	ormOnce.Do(func() {
		orm := new(Orm)
		orm.Engine = engine
		dbEngine = orm
	})

}

func CloseOrmEngine() {
	if dbEngine == nil {
		return
	}
	if err := dbEngine.Close(); nil != err {
		log.Printf("Disconnect from database failed: %s", err.Error())
	}
}

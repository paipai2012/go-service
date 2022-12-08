package dao

import (
	"database/sql"
	"sale-service/engine"
	"sale-service/model"
)

type LuckDao struct {
	Db *engine.Orm
}

// var LuckDaoInstance = LuckDao{DbEngine: engine.GetOrmEngine()}

// func init() {
// 	fmt.Print("luck dao init")
// 	LuckDaoInstance = &LuckDao{DbEngine: engine.GetOrmEngine()}
// }

// 添加用户
func (ld *LuckDao) InsertLuck(luck *model.Luck) (int64, error) {
	return ld.Db.InsertOne(luck)
}

func (ld *LuckDao) AddDraw(item *model.LuckItem, record *model.LuckRecord, user *model.LuckUser) error {
	session := ld.Db.NewSession()
	defer session.Close()

	// add Begin() before any action
	if err := session.Begin(); err != nil {
		return err
	}

	if user.Id == 0 {
		if _, err := session.Insert(user); err != nil {
			return err
		}
	}

	record.UserId = user.Id
	if _, err := session.Insert(record); err != nil {
		return err
	}

	newItem := &model.LuckItem{Quantity: item.Quantity - 1}
	if _, err := session.ID(item.Id).Update(newItem); err != nil {
		return err
	}
	// if item, err := session.Query(fmt.Sprintf("select quantity from luck_item where id= %d for update", record.ItemId)); err != nil {
	// 	return err
	// }

	// add Commit() after all actions
	return session.Commit()
}

// 添加用户
func (ld *LuckDao) UpdateUserPhone(luckUser *model.LuckUser) (sql.Result, error) {
	sql := "update `luck_user` set mobile=? where username=?"
	return ld.Db.Exec(sql, luckUser.Mobile, luckUser.Username)
}

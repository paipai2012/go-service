package app

import (
	"moose-go/engine"
	"moose-go/util"
)

func InitEngine() {
	config := util.GetConfig()

	//  init mysql
	engine.NewOrmEngine(config)

	// init redis
	engine.NewRedisEngine()
}

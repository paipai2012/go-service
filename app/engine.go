package app

import (
	"sale-service/engine"
)

func InitEngine() {

	// init mysql
	engine.NewOrmEngine()

	// init redis
	engine.NewRedisEngine()
}

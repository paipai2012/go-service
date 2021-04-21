package app

import (
	"moose-go/engine"
)

func InitEngine() {

	// init mysql
	engine.NewOrmEngine()

	// init redis
	engine.NewRedisEngine()
}

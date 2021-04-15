package service

import (
	"context"
	"moose-go/engine"
	"testing"

	"github.com/go-redis/redis/v8"
)

func init() {
	engine.NewRedisEngine()
}

func TestSignIn(t *testing.T) {
	var redisHelper = engine.GetRedisEngine()
	redisHelper.SetBit(context.Background(), "moose-go-signin", 12, 10)
}
func TestGetSignIn(t *testing.T) {
	var redisHelper = engine.GetRedisEngine()
	r, _ := redisHelper.GetBit(context.Background(), "moose-go-signin", 1).Result()
	t.Log(r)
	r, _ = redisHelper.GetBit(context.Background(), "moose-go-signin", 12).Result()
	t.Log(r)
}

func TestAddGeo(t *testing.T) {
	var redisHelper = engine.GetRedisEngine()
	redisHelper.GeoAdd(context.Background(), "moose-go-geo", &redis.GeoLocation{
		Longitude: 106.725241,
		Latitude:  26.594446,
	})
}

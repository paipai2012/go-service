package test

import (
	"context"
	"sale-service/engine"
	"testing"

	"github.com/go-redis/redis/v8"
)

func init() {
	// engine.NewRedisEngine()
}

func TestSignIn(t *testing.T) {
	var redisHelper = engine.GetRedisEngine()
	redisHelper.SetBit(context.Background(), "sale-service-signin", 12, 10)
}
func TestGetSignIn(t *testing.T) {
	var redisHelper = engine.GetRedisEngine()
	r, _ := redisHelper.GetBit(context.Background(), "sale-service-signin", 1).Result()
	t.Log(r)
	r, _ = redisHelper.GetBit(context.Background(), "sale-service-signin", 12).Result()
	t.Log(r)
}

func TestAddGeo(t *testing.T) {
	var redisHelper = engine.GetRedisEngine()
	redisHelper.GeoAdd(context.Background(), "sale-service-geo", &redis.GeoLocation{
		Longitude: 106.725241,
		Latitude:  26.594446,
	})
}

package test

import (
	"context"
	"log"
	"testing"

	"github.com/go-redis/redis/v8"
)

const MY_KEY = "u:969856:202104"

var ctx = context.Background()

func TestRedis(t *testing.T) {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	set := client.Set(ctx, "mykey", "foobar", 0)
	log.Println(set)

	bitCount := client.BitCount(ctx, "mykey", nil)
	log.Println(bitCount)

	bitCount = client.BitCount(ctx, "mykey", &redis.BitCount{Start: 0,
		End: 0})
	log.Println(bitCount)

	bitCount = client.BitCount(ctx, "mykey", &redis.BitCount{
		Start: 1,
		End:   1,
	})
	log.Println(bitCount)

	log.Println("---------------------------------------")

	nn, err := client.BitField(ctx, "mykey", "INCRBY", "i5", 100, 1, "GET", "u4", 0).Result()
	log.Println(nn, err)
}

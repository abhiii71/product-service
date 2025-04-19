package cache

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func ConnectRedis(host, port string) *redis.Client {
	addr := fmt.Sprintf("%s: %s", host, port)
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	return rdb
}

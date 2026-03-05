package cache

import (
	"fmt"
	"context"

	"github.com/redis/go-redis/v9"
)


func ConnectToRedis(host string, port string) {
	ctx := context.Background()

	if len(host) == 0 {
		panic("environment variable REDIS_HOST is not set")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(fmt.Sprintf("[-] Failed to connect to Redis: %v", err))
	} else {
		fmt.Println(fmt.Sprintf("[+] Successfully connected to redis at %s on port %v", host, port))
	}

	return
}
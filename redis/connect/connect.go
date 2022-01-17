package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// Redis - 连接示例

// redis全局变量
var redisClient *redis.Client

func initRedis() (err error) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "192.168.40.132:6379",
		Password: "123456",
		DB:       0,
	})

	// 上下文信息
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("Ping redis server failed, error is %v.\n", err)
		return
	}

	return
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("Connect to redis server failed, error is %v.\n", err)
		return
	}

	fmt.Println("Connect to redis server success...")
}

package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// Redis - ZSet示例

// redis全局变量
var redisClient *redis.Client
var ctx context.Context

func initRedis() (err error) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "192.168.40.132:6379",
		Password: "123456",
		DB:       0,
	})

	// 上下文信息
	//var cancel context.CancelFunc
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()

	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("Ping redis server failed, error is %v.\n", err)
		return
	}

	return
}

// zSetDemo
func zSetDemo() {
	zsetKey := "language_rank"
	languages := []*redis.Z{
		&redis.Z{Score: 90.0, Member: "Golang"},
		&redis.Z{Score: 98.0, Member: "Java"},
		&redis.Z{Score: 95.0, Member: "Python"},
		&redis.Z{Score: 97.0, Member: "JavaScript"},
		&redis.Z{Score: 99.0, Member: "C/C++"},
	}

	// ZADD
	num, err := redisClient.ZAdd(ctx, zsetKey, languages...).Result()
	if err != nil {
		fmt.Printf("ZAdd failed, err:%v\n", err)
		return
	}
	fmt.Printf("ZAdd %d success.\n", num)

	// 把Golang的分数加10
	newScore, err := redisClient.ZIncrBy(ctx, zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("ZIncrBy failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret, err := redisClient.ZRevRangeWithScores(ctx, zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Printf("ZRevRangeWithScores failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = redisClient.ZRangeByScoreWithScores(ctx, zsetKey, &op).Result()
	if err != nil {
		fmt.Printf("ZRangeByScoureWithScores failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("Connect to redis server failed, error is %v.\n", err)
		return
	}

	fmt.Println("Connect to redis server success...")

	zSetDemo()
}

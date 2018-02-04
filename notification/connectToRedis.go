package main

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func connectToRedis(config Config) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
		DB:       config.Redis.Database,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		fmt.Fprintln(os.Stderr, "(REDIS) ERROR:", err)
	}
	fmt.Println("REDIS IS WORKING")
	return redisClient
}

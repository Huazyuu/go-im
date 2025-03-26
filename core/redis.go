package core

import (
	"context"
	"github.com/go-redis/redis"
	"time"
)

func InitRedis(addr, pwd string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
		PoolSize: 100,
	})
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
		return nil
	}
	return client
}

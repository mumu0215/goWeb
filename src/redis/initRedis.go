package redis

import (
	"github.com/go-redis/redis"
	"src/middleaware"
)

var RedisClient * redis.Client

func init() {
	RedisClient=redis.NewClient(&redis.Options{
		Addr:               "127.0.0.1:6379",
		Password:           "",
		DB:                 0,
	})
	_,err:=RedisClient.Ping().Result()
	if err!=nil{
		middleaware.Error(err)
	}
}
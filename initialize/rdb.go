package initialize

import (
	"context"
	"log"

	"github.com/akazwz/go-gin-restful-api/global"
	"github.com/go-redis/redis/v8"
)

func InitRDB() *redis.Client {
	c := global.CFG.RedisDB

	if c.Addr == "" {
		log.Println("config can not be null")
		return nil
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password, // no password set
		DB:       0,          // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Println("connect redis error")
		return nil
	}
	return rdb
}

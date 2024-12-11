package dao

import (
	"github.com/Camelia-hu/gomall/conf"
	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

// RedisInit 1
func RedisInit() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Conf.GetString("data.redis.addr"),
		Password: conf.Conf.GetString("data.redis.password"),
		DB:       0,
	})

	Rdb = rdb
}

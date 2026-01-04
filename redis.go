package seedgo

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var rdb *redis.Client

func ParseRedisConf() {
	if viper.IsSet("redis") {
		rdb = redis.NewClient(&redis.Options{
			Addr:     viper.GetString("redis.host"),
			Password: viper.GetString("redis.password"),
			DB:       viper.GetInt("redis.db"),
		})
	}
}

func RedisDB() *redis.Client {
	return rdb
}

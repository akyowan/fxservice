package common

import (
	"fmt"
	"fxlibraries/redis"
	"fxservice/service/chat/config"
)

var client redis.Client

func init() {
	client = redis.NewClient(&redis.RedisConfig{
		Host: config.Conf.RedisConf.Host,
		Port: config.Conf.RedisConf.Port,
		DB:   config.Conf.RedisConf.DB,
	})
}

func GenerateID8(key string) string {
	redisKey := fmt.Sprintf("MOMO.IDS8.%s", key)
	v, err := client.Incr(redisKey).Result()
	if err != nil {
		return ""
	}
	ids := fmt.Sprintf("%08x", v)
	return ids

}

func GenerateID16(key string) string {
	redisKey := fmt.Sprintf("MOMO.IDS16.%s", key)
	v, err := client.Incr(redisKey).Result()
	if err != nil {
		return ""
	}
	ids := fmt.Sprintf("%016x", v)
	return ids
}

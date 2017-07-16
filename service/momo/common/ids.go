package common

import (
	"fmt"
	"fxlibraries/redis"
	"fxservice/service/momo/config"
)

var client redis.Client

func init() {
	client = redis.NewClient(&(config.Conf.RedisConf))
}

func GenerateID8(key string) string {
	redisKey := fmt.Printf("MOMO.IDS.8")
	v := client.Incr(redisKey).Result()
}

func GenerateID16(key string) string {
}

func GenerateID32(key string) string {
}

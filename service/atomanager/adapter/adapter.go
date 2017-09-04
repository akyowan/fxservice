package adapter

import (
	"fxlibraries/mysql"
	"fxlibraries/redis"
	"fxservice/service/atomanager/config"
)

var dbPool *mysql.DBPool
var redisPool *redis.RedisPool

func init() {
	dbPool = mysql.NewDBPool(mysql.DBPoolConfig{
		Host:         config.Conf.DBConf.Host,
		Port:         config.Conf.DBConf.Port,
		User:         config.Conf.DBConf.User,
		DBName:       config.Conf.DBConf.DBName,
		Password:     config.Conf.DBConf.Password,
		MaxIdleConns: 4,
		MaxOpenConns: 8,
		Debug:        config.IsDebug,
	})

	redisPool = redis.NewPool(&redis.RedisConfig{
		Host: config.Conf.RedisConf.Host,
		DB:   config.Conf.RedisConf.DB,
		Port: config.Conf.RedisConf.Port,
	})
}

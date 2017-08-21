package adapter

import (
	"fxlibraries/mysql"
	"fxservice/service/atomanager/config"
)

var dbPool *mysql.DBPool

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
}

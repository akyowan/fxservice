package adapter

import (
	"fxlibraries/mongo"
	"fxservice/config"
)

var mgoPool *mongo.MgoPool

func init() {
	mgoPool = mongo.NewPool(&mongo.MongodbConfig{
		Host:   config.Conf.MongoDB.Host,
		Port:   config.Conf.MongoDB.Port,
		DBName: config.Conf.MongoDB.DBName,
	})
}

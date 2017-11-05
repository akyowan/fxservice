package adapter

import (
	"aposervice/domain"
	"fxlibraries/loggers"
	"fxlibraries/mongo"
	"fxservice/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
)

var (
	mgoPool *mongo.MgoPool
	storage *oss.Bucket
)

func init() {
	var err error
	mgoPool = mongo.NewPool(&mongo.MongodbConfig{
		Host:   config.Conf.MongoDB.Host,
		Port:   config.Conf.MongoDB.Port,
		DBName: config.Conf.MongoDB.DBName,
	})
	storage, err = newBucket(&config.Conf.Storage)
	if err != nil {
		loggers.Error.Printf("New bucket:%v error:%s", config.Conf.Storage, err.Error())
		panic(err)
	}

}

func newBucket(info *domain.Storage) (*oss.Bucket, error) {
	client, err := oss.New(info.EndPoint, info.AccessID, info.AccessKey)
	if err != nil {
		return nil, err
	}

	if err := client.CreateBucket(info.Bucket); err != nil {
		return nil, err
	}

	bucket, err := client.Bucket(info.Bucket)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}

func PutObject(objectID string, reader io.Reader, options ...oss.Option) error {
	return storage.PutObject(objectID, reader, options...)
}

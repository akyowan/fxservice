package adapter

import (
	"fxlibraries/errors"
	"fxservice/service/chatcenter/domain"
	"math/rand"
	"time"
)

func GetRandNickName() (*domain.NickName, error) {
	db := dbPool.NewConn()
	var nickNames []domain.NickName
	dbResult := db.Find(&nickNames)
	if dbResult.RecordNotFound() {
		return nil, errors.NotFound
	}
	rand.Seed(int64(time.Now().Nanosecond()))
	index := rand.Intn(len(nickNames))
	return &nickNames[index], nil
}

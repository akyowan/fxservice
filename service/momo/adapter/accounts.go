package adapter

import (
	"fxlibraries/errors"
	"fxservice/domain"
	"math/rand"
	"time"
)

func GetNewMomoAccount(province, city string) (*domain.MomoAccount, error) {
	db := dbPool.NewConn().Begin()
	var momoAccount domain.MomoAccount
	dbResult := db.Where("status = ?", domain.MomoAccountUnRegister).Order("tid").First(&momoAccount)
	if dbResult.RecordNotFound() {
		db.Rollback()
		return nil, errors.NotFound
	}
	if dbResult.Error != nil {
		db.Rollback()
		return nil, dbResult.Error
	}
	now := time.Now()
	momoAccount.Status = domain.MomoAccountLocked
	momoAccount.UpdateTime = &now
	if err := db.Save(&momoAccount).Error; err != nil {
		db.Rollback()
		return nil, err
	}
	db.Commit()
	return &momoAccount, nil
}

func AddAccounts(accounts []domain.MomoAccount) error {
	db := dbPool.NewConn().Begin()
	for i := range accounts {
		if err := db.Create(&accounts[i]).Error; err != nil {
			db.Rollback()
			return err
		}
	}
	db.Commit()
	return nil
}

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
package adapter

import (
	"fxlibraries/errors"
	"fxservice/domain"
)

func GetNewMomoAccount(province, city string) (*domain.MomoAccount, error) {
	db := dbPool.NewConn()
	var momoAccount domain.MomoAccount
	dbResult := db.Where("status = ?", domain.UnRegister).Order("tid").First(&momoAccount)
	if dbResult.RecordNotFound() {
		return nil, errors.NotFound
	}
	return &momoAccount, nil
}

//func GetMomoAccount(account string) (*domain.MomoAccount, error) {
//}
//
//func GetDevice(sn string) (*domain.Device, error) {
//}
//
//func GetRandomGPS(province, city string) (*domain.GPSLocation, error) {
//}

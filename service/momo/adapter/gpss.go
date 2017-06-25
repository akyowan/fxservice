package adapter

import (
	"fxlibraries/errors"
	"fxservice/domain"
	"math/rand"
	"time"
)

func GetRandomGPS(province, city string) (*domain.GPSLocation, error) {
	db := dbPool.NewConn()
	var gpss []domain.GPSLocation
	dbResult := db.Where("province = ?", province).Where("city = ?", city).Find(&gpss)
	if dbResult.RecordNotFound() {
		return nil, errors.NotFound
	}
	rand.Seed(int64(time.Now().Nanosecond()))
	index := rand.Intn(len(gpss))
	return &gpss[index], nil
}

func AddGpss(gpss []domain.GPSLocation) error {
	db := dbPool.NewConn().Begin()
	for i := range gpss {
		if err := db.Create(&gpss[i]).Error; err != nil {
			db.Rollback()
			return err
		}
	}
	db.Commit()
	return nil
}

package adapter

import (
	"fxlibraries/errors"
	"fxservice/domain"
	"fxservice/service/momo/common"
	"math/rand"
	"time"
)

const GPS_ID_KEY = "GPS"

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
		gpss[i].GPSID = common.GenerateID8(GPS_ID_KEY)
		if err := db.Create(&gpss[i]).Error; err != nil {
			db.Rollback()
			return err
		}
	}
	db.Commit()
	return nil
}

func GetGPS(GPSID string) (*domain.GPSLocation, error) {
	db := dbPool.NewConn()
	var gps domain.GPSLocation
	dbResult := db.Where("gps_id = ?", GPSID).First(&gps)
	if dbResult.RecordNotFound() {
		return nil, errors.NotFound
	}
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return &gps, nil
}

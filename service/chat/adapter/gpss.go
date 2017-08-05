package adapter

import (
	"fxlibraries/errors"
	"fxservice/domain"
	"fxservice/service/chat/common"
	"math/rand"
	"time"
)

const GPS_ID_KEY = "GPS"

func GetRandomGPS(province, city string) (*domain.GPSLocation, error) {
	db := dbPool.NewConn()
	var gpss []domain.GPSLocation
	dbResult := db.Where("province = ?", province).Where("city = ?", city).Where("type = ?", domain.GPSTypeNormal).Find(&gpss)
	if dbResult.RecordNotFound() {
		dbResult := db.Where("province = ?", province).Where("city = ?", city).Where("type = ?", domain.GPSTypeCentral).Find(&gpss)
		if dbResult.RecordNotFound() {
			return nil, errors.NotFound
		}
		if dbResult.Error != nil {
			return nil, dbResult.Error
		}
	}
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	rand.Seed(int64(time.Now().Nanosecond()))
	index := rand.Intn(len(gpss))
	gps := gpss[index]
	if gps.Type == domain.GPSTypeCentral {
		gps.Latitude += common.RandGPSCentralOffset()
		gps.Longitude += common.RandGPSCentralOffset()
	} else {
		gps.Latitude += common.RandGPSOffset()
		gps.Longitude += common.RandGPSOffset()
	}

	return &gps, nil
}

func AddGpss(gpss []domain.GPSLocation) error {
	db := dbPool.NewConn().Begin()
	for i := range gpss {
		gpss[i].GPSID = common.GenerateID8(GPS_ID_KEY)
		if gpss[i].Type == 0 {
			gpss[i].Type = domain.GPSTypeNormal
		}
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

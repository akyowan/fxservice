package adapter

import (
	"fxlibraries/errors"
	"fxservice/domain"
)

func GetDevice(sn string) (*domain.Device, error) {
	db := dbPool.NewConn()
	var device domain.Device
	dbResult := db.Where("sn = ?", sn).First(&device)
	if dbResult.RecordNotFound() {
		return nil, errors.NotFound
	}
	return &device, nil
}

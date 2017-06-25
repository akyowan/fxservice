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

func AddDevices(devices []domain.Device) error {
	db := dbPool.NewConn().Begin()
	for i := range devices {
		if devices[i].SN == "" || devices[i].IMEI == "" || devices[i].UDID == "" {
			continue
		}
		if err := db.Create(&devices[i]).Error; err != nil {
			db.Rollback()
			return err
		}
	}
	db.Commit()
	return nil
}

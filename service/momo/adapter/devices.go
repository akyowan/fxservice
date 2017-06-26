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
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	return &device, nil
}

func AddDevices(devices []domain.Device) error {
	db := dbPool.NewConn().Begin()
	for i := range devices {
		if devices[i].SN == "" || devices[i].IMEI == "" || devices[i].SEQ == "" {
			continue
		}
		if devices[i].Status == 0 {
			devices[i].Status = domain.DeviceEnable
		}
		if err := db.Create(&devices[i]).Error; err != nil {
			db.Rollback()
			return err
		}
	}
	db.Commit()
	return nil
}

func GetEnableDevice() (*domain.Device, error) {
	db := dbPool.NewConn().Begin()
	var device domain.Device
	dbResult := db.Where("status = ?", domain.DeviceEnable).Order("used").First(&device)
	if dbResult.RecordNotFound() {
		db.Rollback()
		return nil, errors.NotFound
	}
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}
	device.Used += 1
	if err := db.Save(&device).Error; err != nil {
		db.Rollback()
		return nil, dbResult.Error
	}
	db.Commit()
	return &device, nil
}

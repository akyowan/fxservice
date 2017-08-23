package adapter

import (
	"fxservice/service/atomanager/common"
	"fxservice/service/atomanager/domain"
)

type AddDevicesResult struct {
	Exists  []string        `json:"exists"`
	Errors  []domain.Device `json:"errors"`
	Success int             `json:"success"`
}

func AddDevices(devices []domain.Device) (*AddDevicesResult, error) {
	db := dbPool.NewConn().Begin()
	result := AddDevicesResult{
		Success: 0,
		Exists:  []string{},
		Errors:  []domain.Device{},
	}

	var device domain.Device
	for i := range devices {
		d := devices[i]
		if d.Sn == "" {
			result.Errors = append(result.Errors, d)
			continue
		}
		dbResult := db.Select("sn").Where("sn = ?", d.Sn).First(&device)
		if dbResult.Error != nil && !dbResult.RecordNotFound() {
			return nil, dbResult.Error
		}
		if dbResult.Error == nil {
			result.Exists = append(result.Exists, d.Sn)
			continue
		}

		if d.Imei == "" {
			d.Imei = " "
		}

		mKey := " "
		if d.Seq == "" {
			d.Seq = " "
		} else {
			if len(d.Seq) >= 4 {
				mKey = d.Seq[len(d.Seq)-4:]
			}
		}
		if d.Version == "" || d.Version < "9.0" {
			d.Version = common.RandVersion()
		}
		if d.Mac == "" {
			d.Mac = " "
		}
		if d.Wifi == "" {
			d.Wifi = " "
		}
		if d.Model == "" {
			if mKey != "" {
				if v := common.Model(mKey); v != "" {
					d.Model = v
				}
			}
		}
		if d.Model == "" {
			d.Model = common.RandModel()
		}
		d.HardWare = common.HardWare(d.Model)
		d.BuildNum = common.BuildNum(d.Version)
		if d.Imei != "" && len(d.Seq) > 3 {
			d.Type = 1
		}

		if err := db.Create(&d).Error; err != nil {
			db.Rollback()
			return nil, err
		}
		result.Success += 1
	}
	db.Commit()
	return &result, nil
}

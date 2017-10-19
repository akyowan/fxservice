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

func AddDevices(group string, devices []domain.Device) (*AddDevicesResult, error) {
	db := dbPool.NewConn().Begin()
	result := AddDevicesResult{
		Success: 0,
		Exists:  []string{},
		Errors:  []domain.Device{},
	}

	var device domain.Device
	isExists := false
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
			isExists = true
			if d.Seq == "" {
				d.Seq = device.Seq
			}
			if d.BuildNum == "" {
				d.BuildNum = device.BuildNum
			}
			if d.HardWare == "" {
				d.HardWare = device.HardWare
			}
			if d.Imei == "" {
				d.Imei = device.Imei
			}
			if d.Ecid == "" {
				d.Ecid = device.Ecid
			}
			if d.BasebandChipid == "" {
				d.BasebandChipid = device.BasebandChipid
			}
			if d.BasebandVersion == "" {
				d.BasebandVersion = device.BasebandVersion
			}
			if d.Firmware == "" {
				d.Firmware = device.Firmware
			}
			if d.Model == "" {
				d.Model = device.Model
			}
			if d.ModelNum == "" {
				d.ModelNum = device.ModelNum
			}
			if d.HardwareModel == "" {
				d.HardwareModel = device.HardwareModel
			}
			if d.MlbSeq == "" {
				d.MlbSeq = device.MlbSeq
			}
			if d.Region == "" {
				d.Region = device.Region
			}
			if d.Wifi == "" {
				d.Wifi = device.Wifi
			}
            if d.Version == "" {
                d.Version = device.Version
            }
			result.Exists = append(result.Exists, d.Sn)
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
		if d.BuildNum == "" {
			d.BuildNum = common.BuildNum(d.Version)
		}
		if d.HardWare == "" {
			d.HardWare = common.HardWare(d.Model)
		}
		if d.Imei != "" && len(d.Seq) > 3 {
			d.Type = 1
		}
		if group != "" {
			d.Group = group
		}

		if isExists {
			if err := db.Table(d.TableName()).Where("sn = ?", d.Sn).Updates(d).Error; err != nil {
				db.Rollback()
				return nil, err
			}
		} else {
			if err := db.Create(&d).Error; err != nil {
				db.Rollback()
				return nil, err
			}
		}
		result.Success += 1
	}
	db.Commit()
	return &result, nil
}

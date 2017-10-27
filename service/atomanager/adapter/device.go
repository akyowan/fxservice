package adapter

import (
	"encoding/base64"
	"encoding/json"
	"fxlibraries/loggers"
	"fxservice/service/atomanager/common"
	"fxservice/service/atomanager/domain"
	"regexp"
	"strings"
	"time"

	version "github.com/hashicorp/go-version"
)

const DEVICE_MIN_ADD = 1

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

	var (
		device   domain.Device
		isExists bool
	)
	for i := range devices {
		isExists = false
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

		minVersion, _ := version.NewVersion("9.0")
		curVersion, err := version.NewVersion(d.Version)
		if err != nil {
			curVersion = minVersion
		}
		if !curVersion.GreaterThan(minVersion) {
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
			d.Status = 1
			if err := db.Create(&d).Error; err != nil {
				db.Rollback()
				return nil, err
			}
		}
		result.Success += 1
	}
	if err := db.Commit().Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func StoreDevice() error {
	key := "report_list"
	devices := make(map[string]domain.Device)
	reg := regexp.MustCompile(`^[\w]{40}$`)
	for {
		var record map[string]string
		result := redisPool.RPop(key)
		resultBytes, err := result.Bytes()
		if err != nil {
			if redisPool.IsNil(err) {
				loggers.Info.Printf("StoreDevice no device")
				return nil
			}
			loggers.Error.Printf("StoreDevice get record from list err:%s", err.Error())
			return err
		}
		if err := json.Unmarshal(resultBytes, &record); err != nil {
			loggers.Info.Printf("StoreDevice json unmarshal error:%s", err.Error())
			loggers.Info.Println(result.String())
			continue
		}
		if action, ok := record["action"]; ok {
			if action == "device_connect" {
				if content, ok := record["content"]; ok {
					decoded, err := base64.StdEncoding.DecodeString(content)
					if err != nil {
						loggers.Warn.Printf("StoreDevice device_connect content decode err:%s", err.Error())
						loggers.Info.Printf(result.String())
						continue
					}

					arr := strings.Split(string(decoded), "||")
					var device domain.Device
					device.Imei = arr[0]
					device.Sn = arr[1]
					device.Seq = arr[2]
					device.Version = arr[3]
					device.Mac = arr[4]
					device.Wifi = arr[5]
					devices[device.Sn] = device
				}
			}
		}
		if sn, ok := record["sn"]; ok {
			if reg.Match([]byte(sn)) {
				if _, ok := devices[sn]; !ok {
					var device domain.Device
					device.Sn = sn
					devices[sn] = device
				}
			}
		}
		listLen := redisPool.LLen(key).Val()
		if listLen <= 0 || len(devices) >= DEVICE_MIN_ADD {
			group := time.Now().Format("20060102")
			var deviceArr []domain.Device
			for _, d := range devices {
				d.Status = 1
				deviceArr = append(deviceArr, d)
			}
			devices = make(map[string]domain.Device)
			result, err := AddDevices(group, deviceArr)
			if err != nil {
				loggers.Error.Printf("StoreDevice add device to db err:%s", err.Error())
				for _, d := range devices {
					loggers.Warn.Println(d)
				}
			}
			loggers.Info.Printf("StoreDevice success:%d exists:%d error:%d", result.Success, len(result.Exists), len(result.Errors))
		}
	}
	return nil
}

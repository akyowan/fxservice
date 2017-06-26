package common

import (
	"fxservice/domain"
)

func CompletionDevices(devices []domain.Device) {
	for i := range devices {
		if devices[i].IDFA == "" {
			devices[i].IDFA = GenRandIDFA()
		}

		if devices[i].IDFV == "" {
			devices[i].IDFV = GenRandIDFV()
		}
		if devices[i].MAC == "" {
			devices[i].MAC = GenRandMac()
		}
		if devices[i].WIFI == "" {
			devices[i].WIFI = GenRandWifi()
		}
		if devices[i].IOSVersion == "" || devices[i].IOSVersion < "10.3.1" {
			devices[i].IOSVersion = GenRandIOSVersion()
		}
		if devices[i].Region == "" {
			devices[i].Region = GenRandRegion()
		}

		if devices[i].Model == "" || devices[i].Model < "iPhone9,1" || devices[i].ModelNum == "" {
			model := GenRandModel()
			devices[i].Model = model.Model
			devices[i].ModelNum = model.ModelNum
		}

		if devices[i].DeviceName == "" {
			devices[i].DeviceName = GenRandDeviceName()
		}
	}
}

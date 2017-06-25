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
		if devices[i].IOSVersion == "" {
			devices[i].IOSVersion = GenRandIOSVersion()
		}
		if devices[i].Region == "" {
			devices[i].Region = GenRandRegion()
		}
		if devices[i].ModelNum == "" {
			devices[i].ModelNum = GenRandModelNum()
		}
		if devices[i].Model == "" {
			devices[i].Model = GenRandModel()
		}
	}
}

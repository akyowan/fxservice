package common

import (
	"fmt"
	"fxlibraries/stringhelper"
	"fxservice/service/chatcenter/domain"
	"strings"
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

func GenRandIDFA() string {
	idfa := UUID()
	return strings.ToUpper(idfa)

}

func GenRandIDFV() string {
	idfv := UUID()
	return strings.ToUpper(idfv)
}

func GenRandIOSVersion() string {
	return "10.3.1"
}

func GenRandMac() string {
	mac := fmt.Sprintf("%s:%s:%s:%s:%s:%s",
		RandHex(2),
		RandHex(2),
		RandHex(2),
		RandHex(2),
		RandHex(2),
		RandHex(2))
	return mac
}

func GenRandWifi() string {
	wifi := fmt.Sprintf("%s:%s:%s:%s:%s:%s",
		RandHex(2),
		RandHex(2),
		RandHex(2),
		RandHex(2),
		RandHex(2),
		RandHex(2))

	return wifi
}

func GenRandRegion() string {
	return "CH/A"
}

func GenRandPassword(n int) string {
	return stringhelper.GererateHash(n)
}

func GenRandDeviceName() string {
	return fmt.Sprintf("%s的iPhone", GenRandName())
}

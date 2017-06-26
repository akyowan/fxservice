package common

import (
	"fmt"
	"fxlibraries/stringhelper"
	"math/rand"
	"strings"
)

func GenRandIDFA() string {
	idfa := fmt.Sprintf("%s-%d-%s-%s-%s",
		stringhelper.GererateHash(4),
		rand.Intn(9000)+1000,
		stringhelper.GererateHash(4),
		stringhelper.GererateHash(4),
		stringhelper.GererateHash(12))
	return strings.ToUpper(idfa)

}

func GenRandIDFV() string {
	idfv := fmt.Sprintf("%s-%d-%s-%s-%s",
		stringhelper.GererateHash(4),
		rand.Intn(9000)+1000,
		stringhelper.GererateHash(4),
		stringhelper.GererateHash(4),
		stringhelper.GererateHash(12))
	return strings.ToUpper(idfv)
}

func GenRandIOSVersion() string {
	return "10.3.1"
}

func GenRandMac() string {
	mac := fmt.Sprintf("%s:%s:%s:%s:%s:%s",
		stringhelper.GererateHash(2),
		stringhelper.GererateHash(2),
		stringhelper.GererateHash(2),
		stringhelper.GererateHash(2),
		stringhelper.GererateHash(2),
		stringhelper.GererateHash(2))
	return mac
}

func GenRandWifi() string {
	wifi := fmt.Sprintf("%s:%s:%s:%s:%s:%s",
		stringhelper.GererateHash(2),
		stringhelper.GererateHash(2),
		stringhelper.GererateHash(2),
		stringhelper.GererateHash(2),
		stringhelper.GererateHash(2),
		stringhelper.GererateHash(2))
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

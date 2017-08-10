package common

import (
	"math/rand"
	"time"
)

func RandGPSOffset() float32 {
	rand.Seed(int64(time.Now().Nanosecond()))
	return float32((rand.Intn(4000) - 2000)) / 1000000

}

func RandGPSCentralOffset() float32 {
	rand.Seed(int64(time.Now().Nanosecond()))
	return float32((rand.Intn(8000) - 4000)) / 100000
}

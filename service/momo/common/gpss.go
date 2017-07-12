package common

import (
	"math/rand"
	"time"
)

func RandGPSOffset() float32 {
	rand.Seed(int64(time.Now().Nanosecond()))
	return float32((rand.Intn(4000) - 2000)) / 1000000

}

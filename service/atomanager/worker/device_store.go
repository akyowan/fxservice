package worker

import (
	"time"
)

type DeviceStorager struct {
	Interval time.Duration
}

func (worker *DeviceStorager) Run() {
	for {
	}
}

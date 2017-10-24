package worker

import (
	"fxlibraries/loggers"
	"fxservice/service/atomanager/adapter"
	"time"
)

type DeviceStorager struct {
	Interval time.Duration
}

func (worker *DeviceStorager) Run() {
	for {
		loggers.Info.Printf("Start device store")
		adapter.StoreDevice()
		time.Sleep(worker.Interval)
	}
}

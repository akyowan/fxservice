package worker

import (
	"fxlibraries/loggers"
	"fxservice/service/atomanager/adapter"
	"time"
)

var INTERVAL_TIME time.Duration

func init() {
	INTERVAL_TIME = time.Minute * 1
}

func Run() {
	for {
		if err := adapter.RebindAccount(); err != nil {
			loggers.Error.Printf("RebindAccount error %s", err.Error())
		}
		time.Sleep(INTERVAL_TIME)
	}
}

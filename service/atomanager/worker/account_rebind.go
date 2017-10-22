package worker

import (
	"fxlibraries/loggers"
	"fxservice/service/atomanager/adapter"
	"time"
)

type AccountRebinder struct {
	Interval time.Duration
}

func (worker *AccountRebinder) Run() {
	for {
		if err := adapter.RebindAccount(); err != nil {
			loggers.Error.Printf("RebindAccount error %s", err.Error())
		}
		time.Sleep(worker.Interval)
	}
}

package main

import (
	"fxservice/service/atomanager/app"
	"fxservice/service/atomanager/config"
)

func main() {
	app.Start(config.Conf.ServerConf.InternalListenAddress)
}

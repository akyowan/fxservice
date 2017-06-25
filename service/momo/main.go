package main

import (
	"fxservice/service/momo/app"
	"fxservice/service/momo/config"
)

func main() {
	app.Start(config.Conf.ServerConf.InternalListenAddress)
}

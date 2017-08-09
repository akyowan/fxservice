package main

import (
	"fxservice/service/apocenter/app"
	"fxservice/service/apocenter/config"
)

func main() {
	app.Start(config.Conf.ServerConf.InternalListenAddress)
}

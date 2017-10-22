package main

import (
	"fxservice/config"
	"fxservice/service/logcenter/app"
)

func main() {
	app.Start(config.Conf.Server.InternalListenAddress)
}

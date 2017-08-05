package main

import (
	"fxservice/service/chat/app"
	"fxservice/service/chat/config"
)

func main() {
	app.Start(config.Conf.ServerConf.InternalListenAddress)
}

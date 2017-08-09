package main

import (
	"fxservice/service/chatcenter/app"
	"fxservice/service/chatcenter/config"
)

func main() {
	app.Start(config.Conf.ServerConf.InternalListenAddress)
}

package app

import (
	"fxlibraries/httpserver"
	"fxlibraries/loggers"
)

func init() {
	loggers.Info.Printf("Initialize...\n")
}

func Start(addr string) {
	r := httpserver.NewRouter()
	loggers.Info.Printf("Starting User Center External Service [\033[0;32;1mOK\t%+v\033[0m] \n", addr)
	panic(r.ListenAndServe(addr))
}

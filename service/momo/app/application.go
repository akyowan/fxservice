package app

import (
	"fxservice/momo/handlers"

	"fxlibraries/httpserver"
	"fxlibraries/loggers"
	//"fxlibraries/errors"
)

func init() {
	loggers.Info.Printf("Initialize...\n")
}

func Start(addr string) {
	r := httpserver.NewRouter()
	r.RouteHandleFunc("/test", handlers.Test).Methods("GET")

	loggers.Info.Printf("Starting User Center External Service [\033[0;32;1mOK\t%+v\033[0m] \n", addr)
	panic(r.ListenAndServe(addr))

}

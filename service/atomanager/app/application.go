package app

import (
	"fxlibraries/httpserver"
	"fxlibraries/loggers"
	"fxservice/service/atomanager/handlers"
)

func init() {
	loggers.Info.Printf("Initialize...\n")
}

func Start(addr string) {
	r := httpserver.NewRouter()
	r.RouteHandleFunc("/accounts/{brief}", handlers.AddAccount).Methods("POST")
	loggers.Info.Printf("Starting ATO  Center External Service [\033[0;32;1mOK\t%+v\033[0m] \n", addr)
	panic(r.ListenAndServe(addr))
}

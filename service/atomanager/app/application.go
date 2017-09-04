package app

import (
	"fxlibraries/errors"
	"fxlibraries/httpserver"
	"fxlibraries/loggers"
	"fxservice/service/atomanager/config"
	"fxservice/service/atomanager/handlers"
)

func init() {
	loggers.Info.Printf("Initialize...\n")
}

func Auth(f httpserver.HandleFunc) httpserver.HandleFunc {
	return func(r *httpserver.Request) *httpserver.Response {
		appKey := r.Header.Get("AppKey")
		if appKey == "" || appKey != config.Conf.ServerConf.AppKey {
			return httpserver.NewResponseWithError(errors.Forbidden)
		}
		return f(r)
	}
}

func Start(addr string) {
	r := httpserver.NewRouter()
	r.RouteHandleFunc("/accounts/{brief}", Auth(handlers.AddAccount)).Methods("POST")
	r.RouteHandleFunc("/devices", Auth(handlers.AddDevice)).Methods("POST")
	loggers.Info.Printf("Starting ATO  Center External Service [\033[0;32;1mOK\t%+v\033[0m] \n", addr)
	panic(r.ListenAndServe(addr))
}

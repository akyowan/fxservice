package app

import (
	"fxlibraries/errors"
	"fxlibraries/httpserver"
	"fxlibraries/loggers"
	"fxservice/config"
	"fxservice/service/logcenter/handler"
)

func init() {
	loggers.Info.Printf("Initialize...\n")
}

func Auth(f httpserver.HandleFunc) httpserver.HandleFunc {
	return func(r *httpserver.Request) *httpserver.Response {
		appKey := r.Header.Get("AppKey")
		if appKey == "" || appKey != config.Conf.Server.AppKey {
			return httpserver.NewResponseWithError(errors.Forbidden)
		}
		return f(r)
	}
}

func Start(addr string) {
	r := httpserver.NewRouter()
	r.RouteHandleFunc("/pdf", handler.PDFLogReport).Methods("POST")

	loggers.Info.Printf("Starting LogCenter External Service [\033[0;32;1mOK\t%+v\033[0m] \n", addr)
	panic(r.ListenAndServe(addr))
}

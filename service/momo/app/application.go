package app

import (
	"fxlibraries/httpserver"
	"fxlibraries/loggers"

	"fxservice/service/momo/handlers"
)

func init() {
	loggers.Info.Printf("Initialize...\n")
}

func Start(addr string) {
	r := httpserver.NewRouter()
	r.RouteHandleFunc("/status", handlers.Test).Methods("GET")

	r.RouteHandleFunc("/accounts", handlers.GetFreeAccounts).Methods("GET").Queries("action", "online")
	r.RouteHandleFunc("/accounts", handlers.GetMomoAccounts).Methods("GET")
	r.RouteHandleFunc("/accounts", handlers.AddAccounts).Methods("POST")
	r.RouteHandleFunc("/accounts", handlers.GetFreeAccounts).Methods("PATCH").Queries("action", "online")
	r.RouteHandleFunc("/accounts", handlers.PatchMomoAccounts).Methods("PATCH")

	r.RouteHandleFunc("/accounts/new", handlers.UnRegisterMomoAccounts).Methods("GET")
	r.RouteHandleFunc("/accounts/{account}", handlers.CompleteMomoAccount).Methods("PATCH")

	r.RouteHandleFunc("/replys/{account}", handlers.GetAccountReply).Methods("GET")

	r.RouteHandleFunc("/gpss", handlers.AddGPSs).Methods("POST")
	r.RouteHandleFunc("/photos", handlers.AddPhotos).Methods("POST")
	r.RouteHandleFunc("/devices", handlers.AddDevices).Methods("POST")

	loggers.Info.Printf("Starting User Center External Service [\033[0;32;1mOK\t%+v\033[0m] \n", addr)
	panic(r.ListenAndServe(addr))

}

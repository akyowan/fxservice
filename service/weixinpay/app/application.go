package app

import (
	"fxlibraries/httpserver"
	"fxlibraries/loggers"
	"fxservice/service/weixinpay/handler"
)

func init() {
	loggers.Info.Printf("Initialize...\n")
}

func Start(addr string) {
	r := httpserver.NewRouter()
	r.RouteHandleFunc("/order/h5", handler.SubmitH5Order).Methods("GET")
	r.RouteHandleFunc("/order/new", handler.SubmitOrder).Methods("GET")
	r.RouteHandleFunc("/callback", handler.PayCallBack).Methods("POST")
	r.RouteHandleFunc("/order/result/{orderID}", handler.PayResult).Methods("GET")
	loggers.Info.Printf("Starting WXPay Demo Service [\033[0;32;1mOK\t%+v\033[0m] \n", addr)
	panic(r.ListenAndServe(addr))
}

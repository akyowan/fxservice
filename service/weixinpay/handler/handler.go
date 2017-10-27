package handler

import (
	"fxlibraries/errors"
	"fxlibraries/httpserver"
	"fxlibraries/loggers"
	"fxservice/service/weixinpay/adapter"
	"time"
)

func SubmitOrder(r *httpserver.Request) *httpserver.Response {
	orderID := time.Now().Format("20060102150405")
	order := &adapter.Order{
		OrderID:    orderID,
		TotalPrice: 1,
		GoodID:     "HEXA",
		Body:       "Vincross HEXA",
		PayMethod:  "NATIVE",
	}
	ip := "12.13.14.15"
	order, err := adapter.SumitOrder(order, ip)
	if err != nil {
		loggers.Error.Printf("SubmitOrder error %s", err.Error())
		return httpserver.NewResponseWithError(errors.InternalServerError)
	}

	resp := httpserver.NewResponse()
	resp.Data = order
	return resp
}

func PayCallBack(r *httpserver.Request) *httpserver.Response {
	resp := httpserver.NewResponse()
	resp.IsWx = true
	return resp
}

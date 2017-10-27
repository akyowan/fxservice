package handler

import (
	"fxlibraries/errors"
	"fxlibraries/httpserver"
	"fxlibraries/loggers"
	"fxservice/service/weixinpay/adapter"
	"strings"
	"time"
)

func SubmitOrder(r *httpserver.Request) *httpserver.Response {
	payMethod := r.QueryParams.Get("payMethod")
	if payMethod == "" {
		payMethod = "NATIVE"
	}
	orderID := time.Now().Format("20060102150405")
	order := &adapter.Order{
		OrderID:    orderID,
		TotalPrice: 1,
		GoodID:     "HEXA",
		Body:       "Vincross HEXA",
		PayMethod:  payMethod,
	}
	ip := strings.Split(r.RemoteAddr, ":")[0]
	order, err := adapter.SumitOrder(order, ip)
	if err != nil {
		loggers.Error.Printf("SubmitOrder error %s", err.Error())
		return httpserver.NewResponseWithError(errors.InternalServerError)
	}
	if payMethod != "NATIVE" {
		return httpserver.NewResponseForRedirect(order.MWebURL)
	}

	resp := httpserver.NewResponse()
	resp.Data = order
	return resp
}

func SubmitH5Order(r *httpserver.Request) *httpserver.Response {
	resp := httpserver.NewResponse()
	resp.Data = "https://api.vincross.com/wxpay/order/new?payMethod=MWEB"
	return resp
}

func PayCallBack(r *httpserver.Request) *httpserver.Response {
	var payResult adapter.PayResult
	resp := httpserver.NewResponse()
	resp.IsWx = true
	if err := r.ParseByXML(&payResult); err != nil {
		loggers.Warn.Printf("Parse weixin pay callback error")
		return resp
	}
	adapter.WXPayCallBack(payResult.OutTradeNO, adapter.OrderStatusPaid)
	return resp
}

func PayResult(r *httpserver.Request) *httpserver.Response {
	orderID := r.UrlParams["orderID"]

	order, err := adapter.GetOrder(orderID)
	if err != nil {
		return httpserver.NewResponseWithError(errors.InternalServerError)
	}

	resp := httpserver.NewResponse()
	resp.Data = order
	return resp
}

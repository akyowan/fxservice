package handler

import (
	"fxlibraries/errors"
	"fxlibraries/httpserver"
	"fxlibraries/loggers"
	"fxservice/service/logcenter/adapter"
	"strings"
)

func PDFLogReport(r *httpserver.Request) *httpserver.Response {
	var logs []map[string]interface{}
	if err := r.Parse(&logs); err != nil {
		loggers.Warn.Printf("PDFLogReport invalid input error:%s", err.Error())
		return httpserver.NewResponseWithError(errors.NewBadRequest("Invalid input"))
	}
	ip := strings.Split(r.RemoteAddr, ":")[0]
	if err := adapter.PDFLogInput(logs, ip); err != nil {
		loggers.Error.Printf("PDFLogReport input log error %s", err.Error())
		return httpserver.NewResponseWithError(errors.InternalServerError)
	}
	return httpserver.NewResponse()
}

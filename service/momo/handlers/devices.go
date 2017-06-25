package handlers

import (
	"fxlibraries/errors"
	"fxlibraries/httpserver"
	"fxlibraries/loggers"
	"fxservice/domain"
	"fxservice/service/momo/adapter"
	"fxservice/service/momo/common"
)

func AddDevices(req *httpserver.Request) *httpserver.Response {
	var devices []domain.Device
	if err := req.Parse(&devices); err != nil {
		loggers.Warn.Printf("AddDevices parse gps error %s", err.Error())
		return httpserver.NewResponseWithError(errors.NewBadRequest("WRONG PARAMTER"))
	}

	common.CompletionDevices(devices)

	if err := adapter.AddDevices(devices); err != nil {
		loggers.Warn.Printf("AddDevices error %s ", err.Error())
		return httpserver.NewResponseWithError(errors.InternalServerError)
	}

	return httpserver.NewResponse()
}

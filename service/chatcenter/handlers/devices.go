package handlers

import (
	"fxlibraries/errors"
	"fxlibraries/httpserver"
	"fxlibraries/loggers"
	"fxservice/service/chatcenter/adapter"
	"fxservice/service/chatcenter/common"
	"fxservice/service/chatcenter/domain"
)

func AddDevices(req *httpserver.Request) *httpserver.Response {
	var devices []domain.Device
	if err := req.Parse(&devices); err != nil {
		loggers.Warn.Printf("AddDevices parse gps error %s", err.Error())
		return httpserver.NewResponseWithError(errors.ParameterError)
	}

	common.CompletionDevices(devices)

	rows, err := adapter.AddDevices(devices)
	if err != nil {
		loggers.Warn.Printf("AddDevices error %s ", err.Error())
		return httpserver.NewResponseWithError(errors.InternalServerError)
	}
	resp := httpserver.NewResponse()
	resp.Data = rows

	return resp
}

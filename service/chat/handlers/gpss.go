package handlers

import (
	"fxlibraries/errors"
	"fxlibraries/httpserver"
	"fxlibraries/loggers"
	"fxservice/domain"
	"fxservice/service/chat/adapter"
)

func AddGPSs(req *httpserver.Request) *httpserver.Response {
	var gpss []domain.GPSLocation
	if err := req.Parse(&gpss); err != nil {
		loggers.Warn.Printf("AddGPSs parse gps error %s", err.Error())
		return httpserver.NewResponseWithError(errors.NewBadRequest("WRONG PARAMTER"))
	}
	if err := adapter.AddGpss(gpss); err != nil {
		loggers.Warn.Printf("AddGpss error %s ", err.Error())
		return httpserver.NewResponseWithError(errors.InternalServerError)
	}

	return httpserver.NewResponse()
}

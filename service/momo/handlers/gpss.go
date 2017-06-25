package handlers

import (
	"fxlibraries/errors"
	"fxlibraries/httpserver"
	"fxlibraries/loggers"
	"fxservice/domain"
	//"fxservice/service/momo/adapter"
)

func AddGPSs(req *httpserver.Request) *httpserver.Response {
	var gpss []domain.GPSLocation
	if err := req.Parse(&gpss); err != nil {
		loggers.Warn.Printf("AddGPSs parse gps error %s", err.Error())
		return httpserver.NewResponseWithError(errors.NewBadRequest("WRONG PARAMTER"))
	}
	return httpserver.NewResponse()
}

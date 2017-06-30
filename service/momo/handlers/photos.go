package handlers

import (
	"fxlibraries/errors"
	"fxlibraries/httpserver"
	"fxlibraries/loggers"
	"fxservice/domain"
	"fxservice/service/momo/adapter"
)

func AddPhotos(req *httpserver.Request) *httpserver.Response {
	var photos [][]domain.Photo
	if err := req.Parse(&photos); err != nil {
		loggers.Warn.Printf("AddPhotos parse photos error %s", err.Error())
		return httpserver.NewResponseWithError(errors.NewBadRequest("WRONG PARAMTER"))
	}
	if err := adapter.AddPhotos(photos); err != nil {
		loggers.Warn.Printf("AddPhotos error %s ", err.Error())
		return httpserver.NewResponseWithError(errors.InternalServerError)
	}
	return httpserver.NewResponse()
}

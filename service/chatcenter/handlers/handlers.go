package handlers

import (
	"fxlibraries/errors"
	"fxlibraries/httpserver"
)

func Test(req *httpserver.Request) *httpserver.Response {
	return httpserver.NewResponseWithError(errors.InternalServerError)
}

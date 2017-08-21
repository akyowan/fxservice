package handlers

import (
	"fxlibraries/errors"
	"fxlibraries/httpserver"
	"fxlibraries/loggers"
	"fxservice/service/atomanager/adapter"
	"fxservice/service/atomanager/domain"
	"strconv"
)

func AddAccount(r *httpserver.Request) *httpserver.Response {
	brief := r.UrlParams["brief"]
	weightStr := r.QueryParams.Get("weight")
	if brief == "" {
		loggers.Warn.Printf("AddAccount no brief")
		return httpserver.NewResponseWithError(errors.NewBadRequest("no brief"))
	}
	weight := 0
	if weightStr == "" {
		weight = 0
	} else {
		var err error
		weight, err = strconv.Atoi(weightStr)
		if err != nil || weight < 0 {
			loggers.Warn.Printf("AddAccount invalid weight")
			return httpserver.NewResponseWithError(errors.NewBadRequest("invalid weight"))
		}
	}

	var accounts []domain.Account
	if err := r.Parse(&accounts); err != nil {
		loggers.Warn.Printf("AddAccount invalid param input")
		return httpserver.NewResponseWithError(errors.NewBadRequest("invalid input"))
	}

	result, err := adapter.AddAccount(brief, weight, accounts)
	if err != nil {
		loggers.Error.Printf("AddAccount error %s", err.Error())
		return httpserver.NewResponseWithError(errors.InternalServerError)
	}

	resp := httpserver.NewResponse()
	resp.Data = *result

	return resp
}

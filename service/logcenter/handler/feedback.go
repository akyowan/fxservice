package handler

import (
	"fxlibraries/errors"
	"fxlibraries/httpserver"
	"fxlibraries/loggers"
	"fxservice/service/logcenter/adapter"
	"fxservice/service/logcenter/domain"
	"strings"
	"time"
)

func Feedback(r *httpserver.Request) *httpserver.Response {
	var feedback domain.FeedBack
	if err := r.Parse(&feedback); err != nil {
		loggers.Warn.Printf("Feedback invalid input param")
		return httpserver.NewResponseWithError(errors.ParameterError)
	}
	ip := strings.Split(r.RemoteAddr, ":")[0]
	now := time.Now()
	feedback.CreatedAt = &now
	feedback.IP = ip
	if err := adapter.FeedbackAdd(&feedback); err != nil {
		loggers.Warn.Printf("Feedback FeedbackAdd error:%s", err.Error())
		return httpserver.NewResponseWithError(errors.InternalServerError)
	}
	return httpserver.NewResponse()
}

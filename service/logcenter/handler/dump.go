package handler

import (
	"fmt"
	"fxlibraries/httpserver"
	"fxlibraries/loggers"
	"fxservice/service/logcenter/adapter"
	"fxservice/service/logcenter/domain"
	"strings"
	"time"
)

func DumpUpload(r *httpserver.Request) *httpserver.Response {
	var dump domain.Dump
	dump.DeviceID = r.UrlParams["device_id"]
	if dump.DeviceID == "" {
		dump.DeviceID = "UNKOWN"
	}
	dump.OS = r.UrlParams["os"]
	dump.Version = r.UrlParams["version"]
	dump.IP = strings.Split(r.RemoteAddr, ":")[0]
	dump.ObjectID = fmt.Sprintf("dump/%s_%d", dump.DeviceID, time.Now().UnixNano())
	if err := adapter.PutObject(dump.ObjectID, r.BodyBuff); err != nil {
		loggers.Error.Printf("DumpUpload PutObject error:%s", err.Error())
		dump.ObjectID = ""
	}
	if err := adapter.DumpAdd(&dump); err != nil {
		loggers.Error.Printf("DumpUpload DumpAdd error:%s", err.Error())
	}
	return httpserver.NewResponse()
}

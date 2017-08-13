package handlers

import (
	"fxlibraries/errors"
	"fxlibraries/httpserver"
	"fxlibraries/loggers"
	"fxservice/service/chatcenter/adapter"
	"fxservice/service/chatcenter/common"
	"fxservice/service/chatcenter/domain"
	"strconv"
	"strings"
)

func UnRegisterTantanAccounts(req *httpserver.Request) *httpserver.Response {
	province := req.QueryParams.Get("province")
	city := req.QueryParams.Get("city")
	if province == "" {
		loggers.Warn.Printf("UnRegisterTantanAccounts no province")
		return httpserver.NewResponseWithError(errors.NewBadRequest("NO PROVINCE"))
	}
	if city == "" {
		loggers.Warn.Printf("UnRegisterTantanAccounts no city")
		return httpserver.NewResponseWithError(errors.NewBadRequest("NO CITY"))
	}

	gps, err := adapter.GetRandomGPS(province, city)
	if err != nil {
		loggers.Warn.Printf("UnRegisterTantanAccounts get %s:%s gps error %s", province, city, err.Error())
		if err == errors.NotFound {
			return httpserver.NewResponseWithError(errors.NewNotFound("NO GPS INFO"))
		}
		return httpserver.NewResponseWithError(errors.InternalServerError)
	}

	device, err := adapter.GetEnableDevice()
	if err != nil {
		loggers.Warn.Printf("UnRegisterTantanAccounts get enable device error %s", err.Error())
		return httpserver.NewResponseWithError(errors.NewNotFound("NO DEVICE INFO"))
	}

	avatar, err := adapter.GetFreeAvatar()
	if err != nil {
		loggers.Warn.Printf("UnRegisterTantanAccounts get enable avatar error %s", err.Error())
		return httpserver.NewResponseWithError(errors.NewNotFound("NO AVATAR INFO"))
	}

	nickName, err := adapter.GetRandNickName()
	if err != nil {
		loggers.Warn.Printf("UnRegisterTantanAccounts get enable nickname error %s", err.Error())
		return httpserver.NewResponseWithError(errors.NewNotFound("NO NICKNAME INFO"))
	}
	operator := common.GenRandOperator()

	tantaAccount := domain.TantanAccount{
		SN:           device.SN,
		Avatar:       avatar.URL,
		PhotosID:     avatar.PhotosID,
		RegisterHost: strings.Split(req.RemoteAddr, ":")[0],
		Province:     gps.Province,
		City:         gps.City,
		Status:       domain.AccountStatusRegistering,
		Operator:     operator.Operator,
		OperatorMC:   operator.OperatorMC,
		OperatorMN:   operator.OperatorMN,
		Gender:       domain.Female,
		NickName:     nickName.NickName,
	}

	if err := adapter.AddTantanAccount(&tantaAccount); err != nil {
		loggers.Warn.Printf("UnRegisterTantanAccounts get enable nickname error %s", err.Error())
		return httpserver.NewResponseWithError(errors.InternalServerError)
	}

	type Resp struct {
		Profile *domain.TantanAccount `json:"profile"`
		Device  *domain.Device        `json:"device"`
		GPS     *domain.GPSLocation   `json:"gps"`
	}

	resp := httpserver.NewResponse()

	resp.Data = Resp{
		Profile: &tantaAccount,
		Device:  device,
		GPS:     gps,
	}

	return resp
}

func CompleteTantanAccount(req *httpserver.Request) *httpserver.Response {
	var tantanAccount domain.TantanAccount
	v := req.UrlParams["id"]
	if v == "" {
		loggers.Warn.Printf("CompleteTantanAccount no id")
		return httpserver.NewResponseWithError(errors.ParameterError)
	}
	id, err := strconv.ParseInt(v, 10, 64)
	if err == nil {
		loggers.Warn.Printf("CompleteTantanAccount ivalid id %s", err.Error())
		return httpserver.NewResponseWithError(errors.ParameterError)
	}

	if err := req.Parse(&tantanAccount); err != nil {
		loggers.Warn.Printf("CompleteTantanAccount parse tantan account error %s", err.Error())
		return httpserver.NewResponseWithError(errors.ParameterError)
	}
	if tantanAccount.Account == "" {
		loggers.Warn.Printf("CompleteTantanAccount no tantan account")
		return httpserver.NewResponseWithError(errors.NewBadRequest("no tantan account"))
	}
	if tantanAccount.Password == "" {
		loggers.Warn.Printf("CompleteTantanAccount no tantan account")
		return httpserver.NewResponseWithError(errors.NewBadRequest("no tantan password"))
	}

	if tantanAccount.Status == 0 {
		tantanAccount.Status = domain.AccountStatusFree
	}
	tantanAccount.RegisterHost = strings.Split(req.RemoteAddr, ":")[0]
	if err := adapter.CompleteTantanAccount(id, &tantanAccount); err != nil {
		loggers.Warn.Printf("CompleteTantanAccount update tantan account error %s", err.Error())
		return httpserver.NewResponseWithError(errors.InternalServerError)
	}

	return httpserver.NewResponse()
}

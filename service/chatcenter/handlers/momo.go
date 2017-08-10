package handlers

import (
	"fxlibraries/errors"
	"fxlibraries/httpserver"
	"fxlibraries/loggers"
	"fxservice/service/chatcenter/adapter"
	"fxservice/service/chatcenter/domain"
	"strconv"
	"strings"
	"time"
)

func UnRegisterMomoAccounts(req *httpserver.Request) *httpserver.Response {
	province := req.QueryParams.Get("province")
	city := req.QueryParams.Get("city")
	if province == "" {
		loggers.Warn.Printf("UnRegisterMomoAccounts no province")
		return httpserver.NewResponseWithError(errors.NewBadRequest("NO PROVINCE"))
	}
	if city == "" {
		loggers.Warn.Printf("UnRegisterMomoAccounts no city")
		return httpserver.NewResponseWithError(errors.NewBadRequest("NO CITY"))
	}

	gps, err := adapter.GetRandomGPS(province, city)
	if err != nil {
		loggers.Warn.Printf("UnRegisterMomoAccounts get %s:%s gps error %s", province, city, err.Error())
		if err == errors.NotFound {
			return httpserver.NewResponseWithError(errors.NewNotFound("NO GPS INFO"))
		}
		return httpserver.NewResponseWithError(errors.InternalServerError)
	}

	momoAccount, err := adapter.GetNewMomoAccount(gps)
	if err != nil {
		loggers.Warn.Printf("UnRegisterMomoAccounts get new account error %s", err.Error())
		if err == errors.NotFound {
			return httpserver.NewResponseWithError(errors.NewNotFound("NO NEW MOMO ACCOUNT"))
		}
		return httpserver.NewResponseWithError(errors.InternalServerError)
	}

	device, err := adapter.GetDevice(momoAccount.SN)
	if err != nil {
		loggers.Warn.Printf("UnRegisterMomoAccounts get %s device error %s", momoAccount.SN, err.Error())
		if err == errors.NotFound {
			return httpserver.NewResponseWithError(errors.NewNotFound("NO DEVICE INFO"))
		}
		return httpserver.NewResponseWithError(errors.InternalServerError)
	}

	type Resp struct {
		Profile *domain.MomoAccount `json:"profile"`
		Device  *domain.Device      `json:"device"`
		GPS     *domain.GPSLocation `json:"gps"`
	}

	resp := httpserver.NewResponse()

	resp.Data = Resp{
		Profile: &domain.MomoAccount{
			Account:      momoAccount.Account,
			Password:     momoAccount.Password,
			AccountType:  momoAccount.AccountType,
			MomoPassword: momoAccount.MomoPassword,
			NickName:     momoAccount.NickName,
			Gender:       momoAccount.Gender,
			Operator:     momoAccount.Operator,
			OperatorMC:   momoAccount.OperatorMC,
			OperatorMN:   momoAccount.OperatorMN,
			Avatar:       momoAccount.Avatar,
		},
		Device: device,
		GPS:    gps,
	}

	return resp
}

func CompleteMomoAccount(req *httpserver.Request) *httpserver.Response {
	var momoAccount domain.MomoAccount
	account := req.UrlParams["account"]
	if account == "" {
		loggers.Warn.Printf("CompleteMomoAccount no account")
		return httpserver.NewResponseWithError(errors.ParameterError)
	}

	if err := req.Parse(&momoAccount); err != nil {
		loggers.Warn.Printf("CompleteMomoAccount parse momo account error %s", err.Error())
		return httpserver.NewResponseWithError(errors.ParameterError)
	}
	if momoAccount.MomoAccount == "" {
		loggers.Warn.Printf("CompleteMomoAccount no momo account")
		return httpserver.NewResponseWithError(errors.NewBadRequest("no momo account"))
	}
	if momoAccount.Status == 0 {
		momoAccount.Status = domain.AccountStatusFree
	}
	momoAccount.RegisterHost = strings.Split(req.RemoteAddr, ":")[0]
	if err := adapter.CompleteMomoAccount(account, &momoAccount); err != nil {
		loggers.Warn.Printf("CompleteMomoAccount update momo account error %s", err.Error())
		return httpserver.NewResponseWithError(errors.InternalServerError)
	}

	return httpserver.NewResponse()
}

func GetMomoAccounts(req *httpserver.Request) *httpserver.Response {
	param := adapter.AccountQueryParam{
		Limit:  10,
		Offset: 0,
	}
	if v := req.QueryParams.Get("account"); v != "" {
		param.Account = v
	}
	if v := req.QueryParams.Get("province"); v != "" {
		param.Province = v
	}
	if v := req.QueryParams.Get("city"); v != "" {
		param.Province = v
	}
	if v := req.QueryParams.Get("momoAccount"); v != "" {
		param.MomoAccount = v
	}
	if v := req.QueryParams.Get("operator"); v != "" {
		param.Operator = v
	}
	if v := req.QueryParams.Get("limit"); v != "" {
		if i, err := strconv.Atoi(v); (err == nil) && (i < 50) && (i > 0) {
			param.Limit = i
		}
	}
	if v := req.QueryParams.Get("gender"); v != "" {
		if i, err := strconv.Atoi(v); (err == nil) && (i > 0) {
			param.Gender = domain.GenderType(i)
		}
	}
	if v := req.QueryParams.Get("offset"); v != "" {
		if i, err := strconv.Atoi(v); (err == nil) && (i > 0) {
			param.Offset = i
		}
	}

	if v := req.QueryParams.Get("status"); v != "" {
		arrs := strings.Split(v, ",")
		for _, s := range arrs {
			if i, err := strconv.Atoi(s); (err == nil) && (i > 0) {
				param.Status = append(param.Status, domain.AccountStatus(i))
			}
		}
	}

	if v := req.QueryParams.Get("type"); v != "" {
		if i, err := strconv.Atoi(v); (err == nil) && (i > 0) {
			param.Type = domain.AccountType(i)
		}
	}
	if v := req.QueryParams.Get("begin"); v != "" {
		timestap, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			return httpserver.NewResponseWithError(errors.NewBadRequest("begin time format error"))
		}
		t := time.Unix(timestap, 0)
		param.Begin = &t
	}
	if v := req.QueryParams.Get("end"); v != "" {
		timestap, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			return httpserver.NewResponseWithError(errors.NewBadRequest("end tie format error"))
		}
		t := time.Unix(timestap, 0)
		param.End = &t
	}
	loggers.Debug.Println(param)

	accounts, err := adapter.GetMomoAccounts(&param)
	if err != nil {
		loggers.Warn.Printf("GetMomoAccounts error %s", err.Error())
		return httpserver.NewResponseWithError(errors.InternalServerError)
	}

	resp := httpserver.NewResponse()

	resp.Data = accounts
	return resp
}

func PatchMomoAccounts(req *httpserver.Request) *httpserver.Response {
	var accounts []domain.MomoAccount
	if err := req.Parse(&accounts); err != nil {
		loggers.Warn.Printf("PatchMomoAccounts parse param error %s", err.Error())
		return httpserver.NewResponseWithError(errors.ParameterError)
	}
	if err := adapter.PatchMomoAccounts(&accounts); err != nil {
		loggers.Warn.Printf("PatchMomoAccounts error %s", err.Error())
		return httpserver.NewResponseWithError(errors.ParameterError)
	}
	return httpserver.NewResponse()
}

func GetFreeMomoAccounts(req *httpserver.Request) *httpserver.Response {
	province := req.QueryParams.Get("province")
	city := req.QueryParams.Get("city")
	account := req.QueryParams.Get("account")
	if province == "" {
		loggers.Warn.Printf("GetFreeAccounts no province param")
		return httpserver.NewResponseWithError(errors.ParameterError)
	}
	if city == "" {
		loggers.Warn.Printf("GetFreeAccounts no city param")
		return httpserver.NewResponseWithError(errors.ParameterError)
	}
	param := adapter.FreeAccountsQueryParam{
		Province: province,
		City:     city,
		Account:  account,
		Limit:    10,
	}
	if v := req.QueryParams.Get("limit"); v != "" {
		if i, err := strconv.Atoi(v); (err == nil) && (i <= 50) && (i > 0) {
			param.Limit = i
		}
	}
	freeAccounts, err := adapter.GetMomoFreeAccounts(&param)
	if err != nil {
		loggers.Error.Printf("GetFreeAccounts get free accounts error %s", err.Error())
		return httpserver.NewResponseWithError(errors.InternalServerError)
	}

	type Object struct {
		Profile *domain.MomoAccount `json:"profile"`
		Device  *domain.Device      `json:"device"`
		GPS     *domain.GPSLocation `json:"gps"`
		RGPS    *domain.GPSLocation `json:"reg_gps,omitempty"`
	}

	var data []Object
	for i := range *freeAccounts {
		account := (*freeAccounts)[i]
		device, err := adapter.GetDevice(account.SN)
		if err != nil {
			loggers.Error.Printf("GetFreeAccounts get device %s error %s", account.SN, err.Error())
			return httpserver.NewResponseWithError(errors.InternalServerError)
		}
		gps, err := adapter.GetRandomGPS(province, city)
		if err != nil {
			loggers.Error.Printf("GetFreeAccounts get gps %s:%s error %s", account.Province, account.City, err.Error())
			return httpserver.NewResponseWithError(errors.InternalServerError)
		}
		obj := Object{
			Profile: &account,
			Device:  device,
			GPS:     gps,
		}

		if account.GPSID != "" {
			rgps, err := adapter.GetGPS(account.GPSID)
			if err != nil {
				loggers.Error.Printf("GetFreeAccounts get register gps %s error %s ", account.GPSID, err.Error())
				return httpserver.NewResponseWithError(errors.InternalServerError)
			}
			obj.RGPS = rgps
		}
		data = append(data, obj)
	}

	resp := httpserver.NewResponse()
	resp.Data = data
	return resp

}

func GetMomoAccountReply(req *httpserver.Request) *httpserver.Response {
	account := req.UrlParams["account"]
	if account == "" {
		loggers.Warn.Printf("GetAccountReply no account")
		return httpserver.NewResponseWithError(errors.ParameterError)
	}
	reply, err := adapter.GetMomoAccountReply(account)
	reply.CreatedAt = nil
	reply.UpdatedAt = nil
	if err != nil {
		loggers.Warn.Printf("GetAccountReply error %s", err.Error())
		return httpserver.NewResponseWithError(errors.NotFound)
	}

	resp := httpserver.NewResponse()
	resp.Data = reply
	return resp
}

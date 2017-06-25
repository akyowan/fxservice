package handlers

import (
	"fxlibraries/errors"
	"fxlibraries/httpserver"
	"fxlibraries/loggers"
	"fxservice/domain"
	"fxservice/service/momo/adapter"
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

	momoAccount, err := adapter.GetNewMomoAccount(province, city)
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

	gps, err := adapter.GetRandomGPS(province, city)
	if err != nil {
		loggers.Warn.Printf("UnRegisterMomoAccounts get %s:%s gps error %s", province, city, err.Error())
		if err == errors.NotFound {
			return httpserver.NewResponseWithError(errors.NewNotFound("NO GPS INFO"))
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

	loggers.Warn.Println(resp)
	return resp
}

func AddAccounts(req *httpserver.Request) *httpserver.Response {
	return httpserver.NewResponse()
}

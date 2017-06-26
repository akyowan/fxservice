package handlers

import (
	"fxlibraries/errors"
	"fxlibraries/httpserver"
	"fxlibraries/loggers"
	"fxservice/domain"
	"fxservice/service/momo/adapter"
	"fxservice/service/momo/common"
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
	var accounts []domain.MomoAccount
	if err := req.Parse(&accounts); err != nil {
		loggers.Warn.Printf("AddAccounts parse accounts error %s", err.Error())
		return httpserver.NewResponseWithError(errors.ParameterError)
	}
	var newAccounts []domain.MomoAccount
	for i := range accounts {
		if accounts[i].Account == "" || accounts[i].AccountType == 0 || accounts[i].Password == "" {
			loggers.Warn.Printf("AddAccounts invalid account %s:%d:%s", accounts[i].Account, accounts[i].AccountType, accounts[i].Password)
			continue
		}
		photosID := adapter.GetRandomPhotosID()
		avatar, err := adapter.GetAvatar(photosID)
		if err != nil {
			loggers.Warn.Printf("AddAccounts get avatar %s error %s", photosID, err.Error())
			continue
		}
		device, err := adapter.GetEnableDevice()
		if err != nil {
			loggers.Warn.Printf("AddAccounts get enable device error %s", err.Error())
			return httpserver.NewResponseWithError(errors.NewNotFound("No enable devices"))
		}
		nickName, err := adapter.GetRandNickName()
		if err != nil {
			loggers.Warn.Printf("AddAccounts get nickname error %s", err.Error())
			return httpserver.NewResponseWithError(errors.NewNotFound("No nicknames "))
		}

		operator := common.GenRandOperator()

		now := time.Now()
		accounts[i].NickName = nickName.NickName
		accounts[i].MomoPassword = common.GenRandPassword(8)

		accounts[i].Operator = operator.Operator
		accounts[i].OperatorMC = operator.OperatorMC
		accounts[i].OperatorMN = operator.OperatorMN

		accounts[i].PhotosID = photosID
		accounts[i].Avatar = avatar
		accounts[i].CreateTime = &now
		accounts[i].SN = device.SN
		accounts[i].Status = domain.MomoAccountUnRegister
		accounts[i].Gender = domain.Female
		newAccounts = append(newAccounts, accounts[i])
	}
	if err := adapter.AddAccounts(newAccounts); err != nil {
		loggers.Warn.Printf("AddAccounts error %s", err.Error())
		return httpserver.NewResponseWithError(errors.InternalServerError)
	}

	return httpserver.NewResponse()
}

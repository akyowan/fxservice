package handlers

import (
	"fxlibraries/errors"
	"fxlibraries/httpserver"
	"fxlibraries/loggers"
	"fxservice/service/chatcenter/adapter"
	"fxservice/service/chatcenter/common"
	"fxservice/service/chatcenter/domain"
)

func AddAccounts(req *httpserver.Request) *httpserver.Response {
	var accounts []domain.MomoAccount
	if err := req.Parse(&accounts); err != nil {
		loggers.Warn.Printf("AddAccounts parse accounts error %s", err.Error())
		return httpserver.NewResponseWithError(errors.ParameterError)
	}
	var newAccounts []domain.MomoAccount
	for i := range accounts {
		if accounts[i].Account == "" || accounts[i].Password == "" {
			loggers.Warn.Printf("AddAccounts invalid account %s:%s", accounts[i].Account, accounts[i].Password)
			continue
		}
		if accounts[i].AccountType == 0 {
			accounts[i].AccountType = domain.QQ
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
		accounts[i].NickName = nickName.NickName
		accounts[i].MomoPassword = common.GenRandPassword(8)
		accounts[i].Operator = operator.Operator
		accounts[i].OperatorMC = operator.OperatorMC
		accounts[i].OperatorMN = operator.OperatorMN
		accounts[i].SN = device.SN
		accounts[i].Status = domain.AccountStatusUnRegister
		accounts[i].Gender = domain.Female
		newAccounts = append(newAccounts, accounts[i])
	}
	if err := adapter.AddAccounts(&newAccounts); err != nil {
		loggers.Warn.Printf("AddAccounts error %s", err.Error())
		return httpserver.NewResponseWithError(errors.InternalServerError)
	}

	return httpserver.NewResponse()
}

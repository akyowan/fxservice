package common

import (
	"math/rand"
	"time"
)

type TelcomOperator struct {
	Operator   string
	OperatorMC string
	OperatorMN string
}

var operators []TelcomOperator

func init() {
	operators = []TelcomOperator{
		{Operator: "中国移动", OperatorMC: "460", OperatorMN: "00"},
		{Operator: "中国联通", OperatorMC: "460", OperatorMN: "01"},
		{Operator: "中国移动", OperatorMC: "460", OperatorMN: "02"},
		{Operator: "中国电信", OperatorMC: "460", OperatorMN: "03"},
		{Operator: "中国电信", OperatorMC: "460", OperatorMN: "05"},
		{Operator: "中国联通", OperatorMC: "460", OperatorMN: "06"},
		{Operator: "中国移动", OperatorMC: "460", OperatorMN: "07"},
		{Operator: "中国铁通集团有限公司", OperatorMC: "460", OperatorMN: "20"},
	}
}

func GenRandOperator() TelcomOperator {
	len := len(operators)
	rand.Seed(int64(time.Now().Nanosecond()))
	return operators[rand.Intn(len)]
}

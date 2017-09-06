package domain

import (
	"time"
)

// tantan账号信息
type TantanAccount struct {
	ID           int64         `json:"id" gorm:"primary_key;column:tid;unique_index:tantan_account_pkey"`
	Account      string        `json:"account,omitempty" gorm:"default:NULL"`
	AccountType  AccountType   `json:"account_type,omitempty" gorm:"default:NULL"`
	Password     string        `json:"password,omitempty" gorm:"default:NULL"`
	Gender       GenderType    `json:"gender,omitempty"`
	NickName     string        `json:"nick_name,omitempty"`
	SN           string        `json:"sn,omitempty"`
	Operator     string        `json:"operator,omitempty"`
	OperatorMC   string        `json:"operator_mc,omitempty"`
	OperatorMN   string        `json:"operator_mn,omitempty"`
	Province     string        `json:"province,omitempty"`
	City         string        `json:"city,omitempty"`
	GPSID        string        `json:"-" gorm:"column:gps_id"`
	RegisterTime *time.Time    `json:"register_time,omitempty"`
	RegisterHost string        `json:"register_host,omitempty"`
	Status       AccountStatus `json:"status"`
	CreatedAt    *time.Time    `json:"create_time,omitempty" gorm:"column:create_time"`
	UpdatedAt    *time.Time    `json:"update_time,omitempty" gorm:"column:update_time"`
}

func (*TantanAccount) TableName() string {
	return "tantan_accounts"
}

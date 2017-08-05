package domain

import (
	"time"
)

// 陌陌账号信息
type MomoAccount struct {
	ID           int64             `json:"-" gorm:"primary_key;column:tid;unique_index:momo_account_pkey"`
	Account      string            `json:"account,omitempty" gorm:"no null"`
	AccountType  MomoAccountType   `json:"account_type,omitempty"`
	Password     string            `json:"password,omitempty"`
	MomoAccount  string            `json:"momo_account,omitempty"`
	MomoPassword string            `json:"momo_password,omitempty"`
	Avatar       string            `json:"avatar,omitempty"`
	PhotosID     string            `json:"photos_id,omitempty"`
	Gender       GenderType        `json:"gender,omitempty"`
	NickName     string            `json:"nick_name,omitempty"`
	SN           string            `json:"sn,omitempty"`
	Operator     string            `json:"operator,omitempty"`
	OperatorMC   string            `json:"operator_mc,omitempty"`
	OperatorMN   string            `json:"operator_mn,omitempty"`
	Province     string            `json:"province,omitempty"`
	City         string            `json:"city,omitempty"`
	GPSID        string            `json:"-" gorm:"column:gps_id"`
	RegisterTime *time.Time        `json:"register_time,omitempty"`
	RegisterHost string            `json:"register_host,omitempty"`
	Status       MomoAccountStatus `json:"status"`
	CreatedAt    *time.Time        `json:"create_time,omitempty" gorm:"column:create_time"`
	UpdatedAt    *time.Time        `json:"update_time,omitempty" gorm:"column:update_time"`
}

func (*MomoAccount) TableName() string {
	return "momo_accounts"
}

type MomoAccountType int

const (
	_     MomoAccountType = iota
	QQ                    // 1 QQ账号
	Phone                 // 2 手机账号
)

type GenderType int

const (
	_      GenderType = iota
	Man               // 1 男
	Female            // 2 女
)

type MomoAccountStatus int

const (
	_                            MomoAccountStatus = iota
	MomoAccountStatusUnRegister                    // 1 未注册
	MomoAccountStatusFree                          // 2 已注册(可用)
	MomoAccountStatusRegistering                   // 3 正在注册
	MomoAccountStatusDisabled                      // 4 被禁用
	MomoAccountStatusOnline                        // 5 在线中
	MomoAccountStatusLocked                        // 6 锁定中(正在注册)
)

// AccountReply
type MomoReply struct {
	ID          int64           `json:"-" gorm:"primary_key;column:tid;unique_index:photos_pkey"`
	Account     string          `json:"account"`
	AccountType MomoAccountType `json:"account_type"`
	ReplyID     string          `json:"reply_id"`
	CreatedAt   *time.Time      `json:"create_time,omitempty" gorm:"column:create_time"`
	UpdatedAt   *time.Time      `json:"update_time,omitempty" gorm:"column:update_time"`
}

func (*MomoReply) TableName() string {
	return "momo_replys"
}

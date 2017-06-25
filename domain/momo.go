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
	RegisterTime *time.Time        `json:"register_time,omitempty"`
	Status       MomoAccountStatus `json:"-"`

	CreateTime *time.Time `json:"create_time"`
	UpdateTime *time.Time `json:"update_time"`
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
	_          MomoAccountStatus = iota
	UnRegister                   // 1 未注册
	Registered                   // 2 已注册
	Disabled                     // 3 被禁用
)

// 硬件信息
type Device struct {
	ID         int64      `json:"-" gorm:"primary_key;column:tid;unique_index:devices_pkey"`
	SN         string     `json:"sn,omitempty"`
	IMEI       string     `json:"imei,omitempty"`
	UDID       string     `json:"udid,omitempty"`
	IOSVersion string     `json:"ios_version,omitempty"`
	MAC        string     `json:"mac,omitempty"`
	WIFI       string     `json:"wifi,omitempty"`
	Model      string     `json:"model,omitempty"`
	IDFA       string     `json:"idfa,omitempty"`
	IDFV       string     `json:"idfv,omitempty"`
	Region     string     `json:"region,omitempty"`
	ModelNum   string     `json:"model_num,omitempty"`
	DeviceName string     `json:"device_name,omitempty"`
	CreateTime *time.Time `json:"create_time,omitempty"`
	UpdateTime *time.Time `json:"update_time,omitempty"`
}

func (*Device) TableName() string {
	return "devices"
}

// GPS信息
type GPSLocation struct {
	ID        int64   `json:"-" gorm:"primary_key;column:tid;unique_index:gps_locations_pkey"`
	Longitude float32 `json:"longitude,omitempty"`
	Latitude  float32 `json:"latitude,omitempty"`
	Province  string  `json:province,omitempty"`
	City      string  `json:city,omitempty"`
}

func (*GPSLocation) TableName() string {
	return "gpss"
}

// 套图信息
type Photos struct {
	ID       int64  `json:"-" gorm:"primary_key;column:tid;unique_index:photos_pkey"`
	PhotosID string `json:"photos_id,omitempty"`
	Seq      int    `json:"seq,omitempty"`
	URL      string `json:"url,omitempty"`
}

func (*Photos) TableName() string {
	return "photos"
}

// 陌陌图片信息
type MomoPhotos struct {
	ID       int64  `json:"-" gorm:"primary_key;column:tid;unique_index:momo_photos"`
	Account  string `json:"account,omitempty"`
	PhotosID string `json:"photos_id,omitempty"`
}

func (*MomoPhotos) TableName() string {
	return "momo_photos"
}

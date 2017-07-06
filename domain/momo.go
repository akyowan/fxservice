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
	_                     MomoAccountStatus = iota
	MomoAccountUnRegister                   // 1 未注册
	MomoAccountRegistered                   // 2 已注册
	MomoAccountLocked                       // 3 锁定中(正在注册)
	MomoAccountDisabled                     // 4 被禁用
	MomoAccountOnline                       // 5 在线中
)

// 硬件信息
type Device struct {
	ID         int64        `json:"-" gorm:"primary_key;column:tid;unique_index:devices_pkey"`
	SN         string       `json:"sn,omitempty"`
	IMEI       string       `json:"imei,omitempty" gorm:"column:imei"`
	SEQ        string       `json:"seq,omitempty" gorm:"column:seq"`
	IOSVersion string       `json:"ios_version,omitempty" gorm:"column:ios_version"`
	MAC        string       `json:"mac,omitempty" gorm:"column:mac"`
	WIFI       string       `json:"wifi,omitempty" gorm:"column:wifi"`
	Model      string       `json:"model,omitempty"`
	IDFA       string       `json:"idfa,omitempty" gorm:"column:idfa"`
	IDFV       string       `json:"idfv,omitempty" gorm:"column:idfv"`
	Region     string       `json:"region,omitempty"`
	ModelNum   string       `json:"model_num,omitempty"`
	DeviceName string       `json:"device_name,omitempty"`
	Used       int          `json:"-"`
	Status     DeviceStatus `json:"-"`
	CreatedAt  *time.Time   `json:"create_time,omitempty" gorm:"column:create_time"`
	UpdatedAt  *time.Time   `json:"update_time,omitempty" gorm:"column:update_time"`
}

func (*Device) TableName() string {
	return "devices"
}

type DeviceStatus int

const (
	_ DeviceStatus = iota
	DeviceEnable
	DeviceDisabled
)

// GPS信息
type GPSLocation struct {
	ID        int64   `json:"-" gorm:"primary_key;column:tid;unique_index:gps_locations_pkey"`
	GPSID     string  `json:"-" gorm:"column:gps_id"`
	Longitude float32 `json:"longitude,omitempty"`
	Latitude  float32 `json:"latitude,omitempty"`
	Province  string  `json:province,omitempty"`
	City      string  `json:city,omitempty"`
}

func (*GPSLocation) TableName() string {
	return "gpss"
}

type PhotoGroup struct {
	ID       int64        `json:"-" gorm:"primary_key;column:tid;unique_index:photo_groups_pkey"`
	PhotosID string       `json:"photos_id" gorm:"unique_index:photos_id_idx"`
	Status   PhotosStatus `json:"photos_status"`
}

func (*PhotoGroup) TableName() string {
	return "photo_groups"
}

type PhotosStatus int

const (
	_                   PhotosStatus = iota
	PhotosStatusFree                 // 可用
	PhotosStatusUsed                 // 已用
	PhotosStatusDisable              // 禁用
)

// 套图信息
type Photo struct {
	ID       int64  `json:"-" gorm:"primary_key;column:tid;unique_index:photos_pkey"`
	PhotosID string `json:"photos_id,omitempty"`
	Seq      int    `json:"seq,omitempty"`
	URL      string `json:"url,omitempty"`
}

func (*Photo) TableName() string {
	return "photos"
}

type NickName struct {
	ID       int64  `json:"-" gorm:"primary_key;column:tid;unique_index:momo_photos"`
	NickName string `json:"-"`
}

func (*NickName) TableName() string {
	return "nick_names"
}

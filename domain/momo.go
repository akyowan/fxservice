package domain

import (
	"time"
)

type MomoAccount struct {
	ID                   int64           `json:"-" gorm:"primary_key;column:tid;unique_index:momo_account_pkey"`
	Account              string          `json:"account,omitempty" gorm:"no null"`
	AccountType          MomoAccountType `json:"account_type"`
	Password             string          `json:"password"`
	MomoAccount          string          `json:"momo_account"`
	MomoPassword         string          `json:"momo_password"`
	Avatar               string          `json:"avatar"`
	PhotosID             string          `json:"photos_id"`
	Gender               GenderType      `json:"gender"`
	NickName             string          `json:"nick_name"`
	SN                   string          `json:"sn"`
	RegistrationProvince string          `json:"registration_province"`
	RegistrationCity     string          `json:"registration_city"`
	RegisterTime         *time.Time      `json:"register_time"`
	CreateTime           *time.Time      `json:"create_time"`
	UpdateTime           *time.Time      `json:"update_time"`
}

func (*MomoAccount) TableName() string {
	return "momo_accounts"
}

type MomoAccountType string

const (
	QQ    MomoAccountType = "QQ"
	Phone MomoAccountType = "Phone"
)

type GenderType string

const (
	Female GenderType = "Female"
	Man    GenderType = "Man"
)

type Device struct {
	ID         int64      `json:"-" gorm:"primary_key;column:tid;unique_index:devices_pkey"`
	SN         string     `json:"sn"`
	IMEI       string     `json:"imei"`
	UDID       string     `json:"udid"`
	IOSVersion string     `json:"ios_version"`
	MAC        string     `json:"mac"`
	WIFI       string     `json:"wifi"`
	Model      string     `json:"model"`
	IDFA       string     `json:"idfa"`
	IDFV       string     `json:"idfv"`
	Region     string     `json:"region"`
	ModelNum   string     `json:"model_num"`
	DeviceName string     `json:"device_name"`
	CreateTime *time.Time `json:"create_time"`
	UpdateTime *time.Time `json:"update_time"`
}

func (*Device) TableName() string {
	return "devices"
}

type GPSLocation struct {
	ID               int64   `json:"-" gorm:"primary_key;column:tid;unique_index:gps_locations_pkey"`
	Account          string  `json:"-"`
	Longitude        float32 `json:"longitude"`
	Latitude         float32 `json:"latitude"`
	LocationProvince string  `json:location_province"`
	LocationCity     string  `json:location_city"`
}

func (*GPSLocation) TableName() string {
	return "gps_locations"
}

type Photos struct {
	ID       int64  `json:"-" gorm:"primary_key;column:tid;unique_index:photos_pkey"`
	PhotosID string `json:"photos_id"`
	Seq      int    `json:"seq"`
	URL      string `json:"url"`
}

func (*Photos) TableName() string {
	return "photos"
}

type MomoPhotos struct {
	ID       int64  `json:"-" gorm:"primary_key;column:tid;unique_index:momo_photos"`
	Account  string `json:"account"`
	PhotosID string `json:"photos_id"`
}

func (*MomoPhotos) TableName() string {
	return "momo_photos"
}

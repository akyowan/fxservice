package domain

import "time"

// Device information
type Device struct {
	ID              int64      `json:"-" gorm:"primary_key;column:id;unique_index:devices_pkey"`
	SN              string     `json:"sn" gorm:"unique_index:devices_sn_idx"`
	IMEI            string     `json:"imei" gorm:"column:imei"`
	SEQ             string     `json:"seq" gorm:"column:seq"`
	IOSVersion      string     `json:"ios_version" gorm:"column:ios_version"`
	MAC             string     `json:"mac" gorm:"column:mac"`
	WIFI            string     `json:"wifi" gorm:"column:wifi"`
	Model           string     `json:"model"`
	IDFA            string     `json:"idfa" gorm:"column:idfa"`
	IDFV            string     `json:"idfv" gorm:"column:idfv"`
	Region          string     `json:"region"`
	BuildNum        string     `json:"build_num"`
	Hardware        string     `json:"hard_ware"`
	HardwareModel   string     `json:"hard_ware_model"`
	ECID            string     `json:"ecid"`
	ModelNum        string     `json:"model_num"`
	Firmware        string     `json:"firmware"`
	MlbSeq          string     `json:"mlb_seq"`
	BasebandChipID  string     `json:"baseband_chip_id"`
	BasebandVersion string     `json:"baseband_version"`
	BindCount       int        `json:"bind_count"`
	SigninCount     int        `json:"sigin_count"`
	Type            int        `json:"type"`
	Group           string     `json:"group"`
	CreatedAt       *time.Time `json:"create_time" gorm:"column:create_time"`
	UpdatedAt       *time.Time `json:"update_time" gorm:"column:update_time"`
}

func (*Device) TableName() string {
	return "devices"
}

package domain

import (
	"time"
)

// Account information
type Account struct {
	Account   string        `json:"account"`
	Password  string        `json:"password"`
	SN        string        `json:"sn"`
	Group     string        `json:"group"`
	Status    AccountStatus `json:"status"`
	RoundUsed int           `json:"round_used"`
	TotalUsed int           `json:"total_used"`
	Errno     int           `json:"errno"`
	AppList   []int64       `json:"app_list"`
	CreateAt  *time.Time    `json:"create_time"`
	UpdateAt  *time.Time    `json:"update_time"`
}

type AccountStatus int

const (
	_                      AccountStatus = iota
	AccountStatusFree                    // 可用
	AccountStatusLocked                  // 锁定
	AccountStatusRoundRest               // CD中
	AccountStatusDisable                 // 禁用
)

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
	BasebandVersion string     `json"baseband_version"`
	BindCount       int        `json:"bind_count"`
	SigninCount     int        `json:"sigin_count"`
	Type            int        `json:"type"`
	Group           string     `json:"group"`
	DeviceName      string     `json:"device_name"`
	CreatedAt       *time.Time `json:"create_time" gorm:"column:create_time"`
	UpdatedAt       *time.Time `json:"update_time" gorm:"column:update_time"`
}

func (*Device) TableName() string {
	return "devices"
}

// Apo task information
type ApoTask struct {
	ID           int64         `json:"-" gorm:"primary_key;column:id;unique_index:devices_pkey"`
	AppID        string        `json:"app_id"`
	AppName      string        `json:"app_name"`
	BundleID     string        `json:"bundle_id"`
	Total        int           `json:"total"`
	DoneCount    int           `json:"done_count"`
	DoingCount   int           `json:"doing_count"`
	TimeoutCount int           `json:"timeount_count"`
	FailCount    int           `json:"fail_count"`
	ApoKey       int           `json:"apo_key"`
	AccountBrief string        `json:"account_brief"`
	Status       ApoTaskStatus `json:"apo_task_status"`
}

func (*ApoTask) TableName() string {
	return "apo_tasks"
}

type ApoTaskStatus int

const (
	_ ApoTaskStatus = iota
	ApoTaskStatusEnable
	ApoTaskStatusDisable
)

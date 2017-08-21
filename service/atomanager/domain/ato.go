package domain

import (
	"time"
)

type Account struct {
	Id          int        `json:"-"`
	Account     string     `json:"account,omitempty"`
	Passwd      string     `json:"passwd,omitempty"`
	EmailPasswd string     `json:"email_password,omitempty"`
	Brief       string     `json:"brief,omitempty"`
	UsedToday   int        `json:"used_today,omitempty"`
	UsedTotal   int        `json:"used_total,omitempty"`
	Errno       int        `json:"errno,omitempty"`
	Ip          string     `json:"ip,omitempty"`
	Imei        string     `json:"imei,omitempty"`
	Sn          string     `json:"sn,omitempty"`
	Seq         string     `json:"seq,omitempty"`
	Version     string     `json:"version,omitempty"`
	Mac         string     `json:"mac,omitempty"`
	Wifi        string     `json:"wifi,omitempty"`
	Model       string     `json:"model,omitempty"`
	BuildNum    string     `json:"build_num,omitempty"`
	HardWare    string     `json:"hard_ware,omitempty"`
	Status      int        `json:"status,omitempty"`
	Type        int        `json:"type,omitempty"`
	Answer1     string     `json:"answwer1,omitempty"`
	Answer2     string     `json:"answwer1,omitempty"`
	Answer3     string     `json:"answwer1,omitempty"`
	CreateAt    *time.Time `json:"create_time,omitempty" gorm:"column:create_time"`
	UpdateAt    *time.Time `json:"update_time,omitempty" gorm:"column:update_time"`
}

func (*Account) TableName() string {
	return "apo_account_info"
}

type Device struct {
	Id              int    `json:"-"`
	Imei            string `json:"imei,omitempty"`
	Sn              string `json:"sn,omitempty"`
	Seq             string `json:"seq,omitempty"`
	Version         string `json:"version,omitempty"`
	Mac             string `json:"mac,omitempty"`
	Wifi            string `json:"wifi,omitempty"`
	Model           string `json:"model,omitempty"`
	BuildNum        string `json:"build_num,omitempty"`
	HardWare        string `json:"hard_ware,omitempty"`
	HardwareModel   string `json:"hardware_model,omitempty"`
	Ecid            string `json:"ecid,omitempty"`
	Region          string `json:"region,omitempty"`
	ModelNum        string `json:"model_num,omitempty"`
	Firmware        string `json:"firmware,omitempty"`
	MlbSeq          string `json:"mlb_seq,omitempty"`
	BasebandChipid  string `json:"baseband_chipid,omitempty"`
	BasebandVersion string `json:"baseband_version,omitempty"`
	BindCount       int    `json:"bind_count,omitempty"`
	Signin_count    int    `json:"signin_count,omitempty"`
	Status          int    `json:"status,omitempty"`
	Type            int    `json:"type,omitempty"`
	Group           int    `json:"group,omitempty"`
}

func (*Device) TableName() string {
	return "apo_device_info"
}

type AccountGroup struct {
	Id     int    `json:"-"`
	Brief  string `json:"brief,omitempty"`
	Weight int    `json:"weight,omitempty"`
	Total  int    `json:"total,omitempty"`
}

func (*AccountGroup) TableName() string {
	return "apo_account_groups"
}

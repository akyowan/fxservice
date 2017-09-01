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

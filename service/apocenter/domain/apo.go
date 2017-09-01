package domain

import "time"

type ApoTask struct {
	ID           int64         `json:"-" gorm:"primary_key;column:id;unique_index:devices_pkey"`
	AppID        string        `json:"app_id,omitempty"`
	AppName      string        `json:"app_name,omitempty"`
	BundleID     string        `json:"bundle_id,omitempty"`
	Level        int           `json:"level"`
	Total        int           `json:"total"`
	RealTotal    int           `json:"real_total"`
	DoneCount    int           `json:"done_count"`
	DoingCount   int           `json:"doing_count"`
	TimeoutCount int           `json:"timeount_count"`
	FailCount    int           `json:"fail_count,omitempty"`
	ApoKey       int           `json:"apo_key,omitempty"`
	AccountBrief string        `json:"account_brief,omitempty"`
	Cycle        int           `json:"cycel"`
	RemindCycle  int           `json:"remind_cycle"`
	UncatchDay   int           `json:"uncatch_day"`
	TypeModelID  int64         `json:"type_model_id,omitempty"`
	AmoutModelID int64         `json:"amount_model_id,omitempty"`
	Status       ApoTaskStatus `json:"status,omitempty"`
	PreaddCount  int           `json:"preadd_count"`
	PreaddTime   *time.Time    `json:"_time,omitempty"`
	StartTime    *time.Time    `json:"_time,omitempty"`
	EndTime      *time.Time    `json:"_time,omitempty"`
	CreateAt     *time.Time    `json:"_time,omitempty" gorm:"column:create_time"`
	UpdateAt     *time.Time    `json:"_time,omitempty" gorm:"column:update_time"`
}

func (*ApoTask) TableName() string {
	return "apo_tasks"
}

type ApoTaskStatus int

const (
	_                    ApoTaskStatus = iota
	ApoTaskStatusStart                 // 开始
	ApoTaskStatusPause                 // 暂停
	ApoTaskStatusDisable               // 停止,禁止
)

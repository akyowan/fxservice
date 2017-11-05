package domain

import (
	"time"
)

type FeedBack struct {
	Content   string     `json:"content,omitempty" bson:"content,omitempty"`
	Contact   string     `json:"contact,omitempty" bson:"contact,omitempty"`
	IP        string     `json:"ip,omitempty" bson:"ip,omitempty"`
	DeviceID  string     `json:"device_id,omitempty" bson:"device_id,omitempty"`
	OS        string     `json:"os,omitempty" bson:"os,omitempty"`
	Version   string     `json:"version" bson:"version,omitempty"`
	CreatedAt *time.Time `json:"create_time,omitempty" bson:"create_time,omitempty"`
}

package domain

import (
	"time"
)

type Dump struct {
	DeviceID  string     `json:"device_id,omitempty" bson:"device_id,omitempty"`
	OS        string     `json:"os,omitempty" bson:"os,omitempty"`
	Version   string     `json:"version" bson:"version,omitempty"`
	IP        string     `json:"ip,omitempty" bson:"ip,omitempty"`
	ObjectID  string     `json:"object_id,omitempty" bson:"object_id,omitempty"`
	CreatedAt *time.Time `json:"create_time,omitempty" bson:"create_time,omitempty"`
}

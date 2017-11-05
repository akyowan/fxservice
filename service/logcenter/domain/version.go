package domain

import (
	"time"
)

type Version struct {
	NeedUpdate  int             `json:"need_update" bson:"-"`
	Version     string          `json:"version" bson:"version"`
	VersionSeq  int             `json:"-" bson:"version_seq"`
	PkgURL      string          `json:"pkg_url" bson:"pkg_url"`
	PkgSize     int             `json:"pkg_size" bson:"pkg_size"`
	MD5         string          `json:"md5" bson:"md5"`
	UpdateType  UpdateTypeEnum  `json:"update_type" bson:"update_type"`
	VersionType VersionTypeEnum `json:"version_type" bson:"version_type"`
	ReleaseNote string          `json:"release_note" bson:"release_note"`
	ReleaseDate *time.Time      `json:"release_date" bson:"release_date"`
}

type UpdateTypeEnum int

const (
	_                UpdateTypeEnum = iota
	UpdateTypeNormal                // 正常升级类型,手动升级触发
	UpdateTypePop                   // 程序启动,自动弹窗触发
	UpdateTypeSilent                // 服务器自动检测,静默升级
)

type VersionTypeEnum int

const (
	VersionBeat    VersionTypeEnum = iota // beat 版本
	VersionRelease                        // 正式版本
)

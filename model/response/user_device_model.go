package response

import (
	"gin_api_02/global"
	"time"
)

type UserDevice struct {
	global.EXTEND_MODEL
	UserId      int64     `json:"userId" gorm:"column:user_id"` // 用户UUID
	AppType     int       `json:"appType" form:"appType" gorm:"column:app_type" `
	DeviceType  int       `json:"deviceType" form:"deviceType" gorm:"column:device_type" `
	Version     string    `json:"version" form:"version" `
	VersionNum  int       `json:"versionNum,string" form:"versionNum" gorm:"column:version_num" `
	DeviceId    string    `json:"deviceId,string" form:"deviceId" gorm:"column:device_id" `
	DeviceToken string    `json:"deviceToken,string" form:"deviceToken" gorm:"column:device_token" `
	Brand       string    `json:"brand,string" form:"brand" `
	SysVersion  string    `json:"sysVersion,string" form:"sysVersion" gorm:"column:sys_version" `
	LastAt      time.Time `json:"lastAt" form:"lastAt" gorm:"column:last_at" `
}

func (UserDevice) TableName() string {
	return "user_device"
}

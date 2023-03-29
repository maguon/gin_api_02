package response

import (
	"gin_api_02/global"
)

type AppInfo struct {
	global.EXTEND_MODEL
	AppType       int    `json:"appType" form:"appType" gorm:"column:app_type" `
	DeviceType    int    `json:"deviceType" form:"deviceType" gorm:"column:device_type" `
	Version       string `json:"version" form:"version" `
	VersionNum    int    `json:"versionNum,string" form:"versionNum" gorm:"column:version_num" `
	MinVersionNum int    `json:"minVersionNum,string" form:"minVersionNum" gorm:"column:min_version_num" `
	ForceUpdate   int    `json:"forceUpdate" form:"forceUpdate" gorm:"column:force_update" binding:"required"`
	Status        int    `json:"status" form:"status" binding:"required"`
	Url           string `json:"url" form:"url" `
	Remarks       string `json:"remark" form:"remark" `
}

func (AppInfo) TableName() string {
	return "app_info"
}

type AppQuery struct {
	global.EXTEND_SEARCH
	AppInfo
}

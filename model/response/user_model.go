package response

import (
	"gin_api_02/global"
	"time"
)

type UserInfo struct {
	global.EXTEND_MODEL
	Status   int        `json:"status" `
	Phone    string     `json:"phone"  form:"phone"`
	Password string     `json:"-"  form:"password"`
	Email    string     `json:"email"  form:"email"`
	Avatar   string     `json:"avatar" form:"avatar"`
	Name     string     `json:"name" form:"name"`
	Gender   int        `json:"gender" form:"gender"`
	Birth    *time.Time `json:"birth" form:"birth"`
}

type UserQuery struct {
	global.EXTEND_SEARCH
	UserInfo
	BirthStart time.Time `json:"birthStart" form:"birthStart"`
	BirthEnd   time.Time `json:"birthEnd" form:"birthEnd"`
}

type UserDevice struct {
	global.EXTEND_MODEL
	LastLoginAt time.Time `json:"lastLoginAt" form:"lastLoginAt" gorm:"column:last_login_at" `
	UserId      int64     `json:"userId"  form:"userId"  gorm:"column:user_id" `
	AppType     int       `json:"appType"  form:"appType"  gorm:"column:app_type" `
	UserType    int       `json:"userType"  form:"userType"  gorm:"column:user_type" `
	DeviceId    string    `json:"deviceId"  form:"deviceId"  gorm:"column:device_id" `
	DeviceToken string    `json:"deviceToken"  form:"deviceToken"  gorm:"column:device_token" `
	Version     string    `json:"version"  form:"version"  gorm:"column:version" `
	VersionNum  int       `json:"versionNum"  form:"versionNum"  gorm:"column:version_num" `
}

type UserDeviceQuery struct {
	global.EXTEND_SEARCH
	UserDevice
	LastLoginStart time.Time `json:"lastLoginStart" form:"lastLoginStart"`
	LastLoginEnd   time.Time `json:"lastLoginEnd" form:"lastLoginEnd"`
}

func (UserInfo) TableName() string {
	return "user_info"
}

type LoginResponse struct {
	User      UserInfo `json:"user"`
	Token     string   `json:"token"`
	ExpiresAt int64    `json:"expiresAt"`
}

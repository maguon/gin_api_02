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

func (UserInfo) TableName() string {
	return "user_info"
}

type LoginResponse struct {
	User      UserInfo `json:"user"`
	Token     string   `json:"token"`
	ExpiresAt int64    `json:"expiresAt"`
}

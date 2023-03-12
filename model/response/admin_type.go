package response

import (
	"gin_api_02/global"

	"github.com/jackc/pgtype"
)

type AdminTypeQuery struct {
	global.EXTEND_SEARCH
	AdminType
}
type AdminType struct {
	global.EXTEND_MODEL
	TypeName string       `json:"typeName" form:"typeName" gorm:"column:type_name;comment:类型名称"`
	MenuList pgtype.JSONB `json:"menuList" form:"menuList" gorm:"type:jsonb;column:menu_list;comment:类型名称"`
	Remark   string       `json:"remark" form:"remark" gorm:"comment:备注"`
}

func (AdminType) TableName() string {
	return "admin_type"
}

type AdminInfo struct {
	global.EXTEND_MODEL
	Username string `json:"userName" form:"userName"  gorm:"column:user_name;comment:用户登录名"` // 用户登录名
	Password string `json:"password" form:"password" gorm:"comment:用户登录密码"`
	Type     int16  `json:"type" form:"type" gorm:"comment:管理员类型"`
	Gender   int8   `json:"gender" form:"gender" gorm:"comment:管理员性别"`
	Avatar   string `json:"avatar" form:"avatar" gorm:"comment:用户头像"`
	Phone    string `json:"phone" form:"phone" gorm:"comment:用户手机号"` // 用户手机号
	Email    string `json:"email" form:"email" gorm:"comment:用户邮箱"`
	Status   int8   `json:"status" form:"status" gorm:"comment:状态"`
}

func (AdminInfo) TableName() string {
	return "admin_info"
}

type AdminLoginResponse struct {
	Admin     AdminInfo `json:"admin"`
	Token     string    `json:"token"`
	ExpiresAt int64     `json:"expiresAt"`
}

type AdminInfoQueryRes struct {
	AdminInfo
	TypeName string `json:"type_name" form:"type_name" gorm:"column:type_name;comment:类型名称"`
	Remark   string `json:"remark" form:"remark" gorm:"comment:备注"`
}
type AdminUserQuery struct {
	global.EXTEND_SEARCH
	AdminInfo
}

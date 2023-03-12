package request

// Modify password structure
type SysUserPassword struct {
	ID          int64  `json:"-"`           // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

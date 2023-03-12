package service

import (
	"errors"
	"fmt"
	"gin_api_02/global"
	res "gin_api_02/model/response"
	"gin_api_02/utils"
)

type UserService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func (userService *UserService) Login(u *res.UserInfo) (userInter *res.UserInfo, err error) {
	if nil == global.SYS_DB {
		return nil, fmt.Errorf("db not init")
	}

	var user res.UserInfo
	err = global.SYS_DB.Where("user_name = ?", u.Username).First(&user).Error
	global.SYS_LOG.Info(user.Password)
	global.SYS_LOG.Info(utils.BcryptHash(u.Password))
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &user, err
}

//@function: GetUserInfo
//@description: 获取用户信息
//@param: uuid uuid.UUID
//@return: err error, user system.SysUser

func (userService *UserService) GetUserInfo(userId int64) (user res.UserInfo, err error) {
	var reqUser res.UserInfo
	err = global.SYS_DB.First(&reqUser, "id = ?", userId).Error
	if err != nil {
		return reqUser, err
	}
	return reqUser, err
}

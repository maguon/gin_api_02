package service

import (
	"errors"
	"fmt"
	"gin_api_02/global"
	"gin_api_02/utils"

	req "gin_api_02/model/request"
	res "gin_api_02/model/response"
)

type AdminService struct{}

func (adminService *AdminService) AdminLogin(u *res.AdminInfo) (adminRes *res.AdminInfo, err error) {
	if nil == global.SYS_DB {
		return nil, fmt.Errorf("db not init")
	}

	var adminInfo res.AdminInfo
	err = global.SYS_DB.Where("user_name = ?", u.Username).First(&adminInfo).Error
	global.SYS_LOG.Info(adminInfo.Password)
	global.SYS_LOG.Info(utils.BcryptHash(u.Password))
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, adminInfo.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &adminInfo, err
}

func (adminService *AdminService) GetAdminInfo(adminId int64) (adminUser res.AdminInfo, err error) {
	var reqUser res.AdminInfo
	err = global.SYS_DB.First(&reqUser, "id = ?", adminId).Error
	if err != nil {
		return reqUser, err
	}
	fmt.Println(reqUser)
	return reqUser, err
}

func (adminService *AdminService) GetSysUserList(queryModel res.AdminUserQuery) (list interface{}, total int64, err error) {
	limit := queryModel.PageSize
	offset := queryModel.PageSize * (queryModel.PageNumber - 1)
	fmt.Println("limit", limit)
	fmt.Println("offset", queryModel.PageNumber)
	db := global.SYS_DB.Table("admin_info ai ").Joins(" left join admin_type adt on ai.type = adt.id").Select("ai.*,adt.type_name,adt.remark")
	var adminInfoList []res.AdminInfoQueryRes
	if queryModel.ID != 0 {
		db = db.Where("ai.id = ?", queryModel.ID)
	}
	if queryModel.Type != 0 {
		db = db.Where("ai.type = ?", queryModel.Type)
	}
	if queryModel.Gender != 0 {
		db = db.Where("ai.gender = ?", queryModel.Gender)
	}
	if queryModel.Email != "" {
		db = db.Where("ai.email = ?", queryModel.Email)
	}
	if queryModel.Phone != "" {
		db = db.Where("ai.phone = ?", queryModel.Phone)
	}
	if queryModel.Username != "" {
		db = db.Where("ai.user_name = ?", queryModel.Username)
	}
	if queryModel.Status != 0 {
		db = db.Where("ai.status = ?", queryModel.Status)
	}
	if !queryModel.CreatedStart.IsZero() && !queryModel.CreatedEnd.IsZero() {
		db = db.Where("ai.created_at >= ?", queryModel.CreatedStart)
		db = db.Where("ai.created_at <= ?", queryModel.CreatedEnd)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Debug().Limit(limit).Offset(offset).Order("id desc").Find(&adminInfoList).Error
	return adminInfoList, total, err
}

func (adminService *AdminService) AddAdminInfo(adminInfo res.AdminInfo) (adminInfoRes res.AdminInfo, err error) {
	err = global.SYS_DB.Create(&adminInfo).Error
	return adminInfo, err
}

func (adminService *AdminService) UpdateAdminInfo(adminInfo res.AdminInfo) (adminInfoRes res.AdminInfo, err error) {
	err = global.SYS_DB.Where("id = ? ", adminInfo.ID).First(&res.AdminInfo{}).Updates(&adminInfo).Error
	return adminInfo, err
}

func (adminService *AdminService) AddAdminType(adminType res.AdminType) (adminTypeRes res.AdminType, err error) {
	err = global.SYS_DB.Create(&adminType).Error
	return adminType, err
}

func (adminService *AdminService) UpdateAdminType(adminType res.AdminType) (adminTypeRes res.AdminType, err error) {
	err = global.SYS_DB.Where("id = ? ", adminType.ID).First(&res.AdminType{}).Updates(&adminType).Error
	return adminType, err
}

func (adminService *AdminService) UpdateAdminPassword(sysUserPassword req.SysUserPassword) (adminUser *res.AdminInfo, err error) {
	var adminInfo res.AdminInfo
	if err = global.SYS_DB.First(&adminInfo, "id = ?", sysUserPassword.ID).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(sysUserPassword.Password, adminInfo.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	adminInfo.Password = utils.BcryptHash(sysUserPassword.NewPassword)
	err = global.SYS_DB.Save(&adminInfo).Error
	return &adminInfo, err

}
func (adminService *AdminService) GetAdminType(adminTypeQuery res.AdminTypeQuery) (list interface{}, total int64, err error) {
	limit := adminTypeQuery.PageSize
	offset := adminTypeQuery.PageSize * (adminTypeQuery.PageNumber - 1)
	db := global.SYS_DB.Model(&res.AdminType{})
	var adminTypeList []res.AdminType
	if adminTypeQuery.ID != 0 {
		db = db.Where("id = ?", adminTypeQuery.ID)
	}
	if adminTypeQuery.TypeName != "" {
		db = db.Where("type_name = ?", adminTypeQuery.TypeName)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&adminTypeList).Error
	return adminTypeList, total, err
}

func (adminService *AdminService) RemoveAdminType(adminTypeId int64) (total int64, err error) {
	db := global.SYS_DB.Table("admin_info ").Select(" id,user_name ")
	db = db.Where("type = ?", adminTypeId)

	err = db.Count(&total).Error
	if err != nil {
		return total, err
	}
	if total == 0 {
		var adminType res.AdminType
		err = global.SYS_DB.Where("id = ? ", adminTypeId).Delete(&adminType).Error
		if err != nil {
			return total, err
		}
	}
	return total, err
}

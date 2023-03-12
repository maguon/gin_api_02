package service

import (
	"fmt"
	"gin_api_02/global"
	res "gin_api_02/model/response"
)

type AppService struct{}

func (appService *AppService) AddAppInfo(appInfo res.AppInfo) (appInfoRes res.AppInfo, err error) {
	err = global.SYS_DB.Create(&appInfo).Error
	return appInfo, err
}

func (appService *AppService) UpdateAppInfo(appInfo res.AppInfo) (appInfoRes res.AppInfo, err error) {
	err = global.SYS_DB.Where("id = ? ", appInfo.ID).First(&res.AppInfo{}).Updates(&appInfo).Error
	return appInfo, err
}

func (appService *AppService) RemoveAppInfo(appInfoId int64) (err error) {
	var appInfo res.AppInfo
	err = global.SYS_DB.Where("id = ? ", appInfoId).Delete(&appInfo).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (appService *AppService) GetAppInfo(queryModel res.AppQuery) (list interface{}, total int64, err error) {
	limit := queryModel.PageSize
	offset := queryModel.PageSize * (queryModel.PageNumber - 1)
	db := global.SYS_DB.Model(&res.AppInfo{})
	var appInfoList []res.AppInfo
	fmt.Println(queryModel)
	if queryModel.ID != 0 {
		db = db.Where("id = ?", queryModel.ID)
	}
	if queryModel.AppType != 0 {
		db = db.Where("app_type = ?", queryModel.AppType)
	}
	if queryModel.DeviceType != 0 {
		db = db.Where("device_type = ?", queryModel.DeviceType)
	}
	if queryModel.Status != 0 {
		db = db.Where("status = ?", queryModel.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Debug().Limit(limit).Offset(offset).Order("id desc").Find(&appInfoList).Error
	return appInfoList, total, err
}

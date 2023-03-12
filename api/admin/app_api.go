package admin

import (
	"gin_api_02/global"
	"gin_api_02/model/common/response"
	res "gin_api_02/model/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateAppInfo
// @Tags      App
// @Summary   新增AppInfo
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param    data  body      res.AppInfo  true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=res.AppInfo,msg=string}  "返回包括用户组信息"
// @Router    /admin/app [post]
func (b *AdminApi) CreateAppInfo(c *gin.Context) {
	var appInfo res.AppInfo
	err := c.ShouldBindJSON(&appInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if appInfoRes, err := appService.AddAppInfo(appInfo); err != nil {
		global.SYS_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(gin.H{"appInfo": appInfoRes}, "创建成功", c)
	}
}

// UpdateAppInfo
// @Tags      App
// @Summary   更新app信息
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param appId path int true "app ID"
// @Param    data  body      res.AppInfo  true  " "
// @Success  200   {object}  response.Response{data=res.AppInfo,msg=string}  "返回AppInfo"
// @Router    /admin/app/{appId} [put]
func (b *AdminApi) UpdateAppInfo(c *gin.Context) {
	var appInfoOut res.AppInfo
	err := c.ShouldBindJSON(&appInfoOut)
	appId, _ := strconv.ParseInt(c.Param("appId"), 10, 64)
	appInfoOut.ID = appId
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	appInfoRes, err := appService.UpdateAppInfo(appInfoOut)
	if err != nil {
		global.SYS_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{"appInfo": appInfoRes}, "更新成功", c)
}

// RemoveAdminType
// @Tags      App
// @Summary   删除app版本信息
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param typeId path int true "App ID"
// @Success   200   {object}  response.Response{data=response.QueryResult,msg=string}  "AdminType列表,返回包括列表,总数,页码,每页数量"
// @Router    /admin/app/{appId} [delete]
func (s *AdminApi) RemoveAppInfo(c *gin.Context) {
	appId, _ := strconv.ParseInt(c.Param("appId"), 10, 64)

	err := appService.RemoveAppInfo(appId)
	if err != nil {
		global.SYS_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithDetailed(response.QueryResult{}, "删除成功", c)
}

package admin

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"gin_api_02/global"
	"gin_api_02/model/common/response"
	req "gin_api_02/model/request"
	res "gin_api_02/model/response"
	"gin_api_02/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetAdminInfo
// @Tags      Admin
// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取用户信息"
// @Router    /admin/sysUser [get]
func (b *AdminApi) GetAdminUserInfo(c *gin.Context) {
	adminId := utils.GetAdminID(c)
	ReqUser, err := adminService.GetAdminInfo(int64(adminId))
	ReqUser.Password = ""
	if err != nil {
		global.SYS_LOG.Error("获取用户信息失败!", zap.Error(err))
		response.FailWithMessage("获取用户信息失败", c)
		return
	}
	var adminTypeQuery res.AdminTypeQuery
	adminTypeQuery.PageSize = 10
	adminTypeQuery.PageNumber = 1
	adminTypeQuery.ID = int64(ReqUser.Type)
	adminType, _, err := adminService.GetAdminType(adminTypeQuery)
	if err != nil {
		global.SYS_LOG.Error("获取用户权限失败!", zap.Error(err))
		response.FailWithMessage("获取用户权限失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"adminInfo": ReqUser, "adminType": adminType}, "获取成功", c)
}

// GetAdminUserList
// @Tags      Admin
// @Summary   分页获取SysUser列表
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @accept    application/json
// @Produce   application/json
// @Param     data  query     res.AdminUserQuery                       true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  response.Response{data=response.QueryResult,msg=string}  "SysUser列表,返回包括列表,总数,页码,每页数量"
// @Router    /admin/sysUserList [get]
func (s *AdminApi) GetAdminUserList(c *gin.Context) {
	var queryModel res.AdminUserQuery
	err := c.ShouldBindQuery(&queryModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := adminService.GetSysUserList(queryModel)
	if err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.QueryResult{
		List:       list,
		Total:      total,
		PageNumber: queryModel.PageNumber,
		PageSize:   queryModel.PageSize,
	}, "获取成功", c)
}

// CreateAdminType
// @Tags      Admin
// @Summary   新增用户组
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param    data  body      res.AdminType  true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=res.LoginResponse,msg=string}  "返回包括用户组信息"
// @Router    /admin/type [post]
func (b *AdminApi) CreateAdminType(c *gin.Context) {
	var adminType res.AdminType
	err := c.ShouldBindJSON(&adminType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if adminTypeRes, err := adminService.AddAdminType(adminType); err != nil {
		global.SYS_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(gin.H{"adminType": adminTypeRes}, "创建成功", c)
	}
}

// UpdateAdminPassword
// @Tags      Admin
// @Summary   新增用户组
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param    data  body      req.SysUserPassword  true  " 原密码, 新密码"
// @Success  200   {object}  response.Response{msg=string}  "返回更新结果信息"
// @Router    /admin/password [put]
func (b *AdminApi) UpdateAdminPassword(c *gin.Context) {
	var sysUserPassword req.SysUserPassword
	err := c.ShouldBindJSON(&sysUserPassword)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	adminId := utils.GetAdminID(c)
	sysUserPassword.ID = adminId
	_, err = adminService.UpdateAdminPassword(sysUserPassword)
	if err != nil {
		global.SYS_LOG.Error("更新密码失败!", zap.Error(err))
		response.FailWithMessage("更新密码失败 "+err.Error(), c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

// CreateAdminInfo
// @Tags      Admin
// @Summary   新增系统用户
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param    data  body      res.AdminInfo  true  "用户信息"
// @Success  200   {object}  response.Response{data=res.AdminInfo,msg=string}  "返回包括用户组信息"
// @Router    /admin/sysUser [post]
func (b *AdminApi) CreateAdminInfo(c *gin.Context) {
	var adminInfo res.AdminInfo
	err := c.ShouldBindJSON(&adminInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	password := adminInfo.Password
	adminInfo.Password = utils.BcryptHash(adminInfo.Password)
	resFlag := utils.BcryptCheck(password, adminInfo.Password)
	fmt.Println(resFlag)
	if adminInfoRes, err := adminService.AddAdminInfo(adminInfo); err != nil {
		global.SYS_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(gin.H{"adminInfo": adminInfoRes}, "创建成功", c)
	}
}

// UpdateAdminInfo
// @Tags      Admin
// @Summary   新增用户组
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param typeId path int true "sys user ID"
// @Param    data  body      res.AdminInfo  true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=res.AdminInfo,msg=string}  "返回包括用户组信息"
// @Router    /admin/sysUser/{sysUserId} [put]
func (b *AdminApi) UpdateAdmiInfo(c *gin.Context) {
	var adminInfo res.AdminInfo
	err := c.ShouldBindJSON(&adminInfo)
	sysUserId, _ := strconv.ParseInt(c.Param("sysUserId"), 10, 64)
	adminInfo.ID = sysUserId
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	adminInfoRes, err := adminService.UpdateAdminInfo(adminInfo)
	if err != nil {
		global.SYS_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{"adminInfo": adminInfoRes}, "更新成功", c)
}

// UpdateAdminType
// @Tags      Admin
// @Summary   新增用户组
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param typeId path int true "type ID"
// @Param    data  body      res.AdminType  true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=res.LoginResponse,msg=string}  "返回包括用户组信息"
// @Router    /admin/type/{typeId} [put]
func (b *AdminApi) UpdateAdminType(c *gin.Context) {
	var adminType res.AdminType
	err := c.ShouldBindJSON(&adminType)
	typeId, _ := strconv.ParseInt(c.Param("typeId"), 10, 64)
	adminType.ID = typeId
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	adminTypeRes, err := adminService.UpdateAdminType(adminType)
	if err != nil {
		global.SYS_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{"adminType": adminTypeRes}, "更新成功", c)
}

// GetAdminType
// @Tags      Admin
// @Summary   分页获取AdminType列表
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @accept    application/json
// @Produce   application/json
// @Param     data  query     res.AdminTypeQuery                       true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  response.Response{data=response.QueryResult,msg=string}  "AdminType列表,返回包括列表,总数,页码,每页数量"
// @Router    /admin/type [get]
func (s *AdminApi) GetAdminType(c *gin.Context) {
	var queryModel res.AdminTypeQuery
	err := c.ShouldBindQuery(&queryModel)
	fmt.Println(queryModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := adminService.GetAdminType(queryModel)
	if err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.QueryResult{
		List:       list,
		Total:      total,
		PageNumber: queryModel.PageNumber,
		PageSize:   queryModel.PageSize,
	}, "获取成功", c)
}

// RemoveAdminType
// @Tags      Admin
// @Summary   分页获取AdminType列表
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param typeId path int true "type ID"
// @Success   200   {object}  response.Response{data=response.QueryResult,msg=string}  "AdminType列表,返回包括列表,总数,页码,每页数量"
// @Router    /admin/type/{typeId} [delete]
func (s *AdminApi) RemoveAdminType(c *gin.Context) {
	typeId, _ := strconv.ParseInt(c.Param("typeId"), 10, 64)

	total, err := adminService.RemoveAdminType(typeId)
	if err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.QueryResult{
		Total: total,
	}, "获取成功", c)
}

// ExportAdminInfoCsv
// @Tags      Admin
// @Summary   导出Sys AdminInfo CSV
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Success   200   {object}  response.Response{data=response.QueryResult,msg=string}  "AdminInfo导出列表,返回包括列表,总数,页码,每页数量"
// @Router    /admin/sysUser.csv [get]
func (s *AdminApi) ExportAdminInfo(c *gin.Context) {
	var queryModel res.AdminUserQuery
	err := c.ShouldBindQuery(&queryModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, _, err := adminService.GetSysUserList(queryModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var byteList []byte
	byteList, _ = json.Marshal(list)
	var adminInfo []res.AdminInfoQueryRes
	json.Unmarshal(byteList, &adminInfo)
	record := []string{"ID", "用户名称", "用户群组", "手机", "邮箱", "性别", "创建时间", "状态"} // just some test data to use for the wr.Writer() method below.
	fmt.Println(len(adminInfo))
	b := &bytes.Buffer{}   // creates IO Writer
	wr := csv.NewWriter(b) // creates a csv writer that uses the io buffer.
	wr.Write(record)
	for i := 0; i < len(adminInfo); i++ { // make a loop for 100 rows just for testing purposes
		recordTemp := []string{strconv.FormatInt(adminInfo[i].ID, 10), adminInfo[i].Username, adminInfo[i].TypeName, adminInfo[i].Phone,
			adminInfo[i].Email, utils.GetGenderStr(adminInfo[i].Gender), adminInfo[i].CreatedAt.Format("2006-01-02 15:04:05"), utils.GetStatusStr(adminInfo[i].Status)}
		wr.Write(recordTemp) // converts array of string to comma seperated values for 1 row.
	}
	wr.Flush() // writes the csv writer data to  the buffered data io writer(b(bytes.buffer))
	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Write(b.Bytes())

}

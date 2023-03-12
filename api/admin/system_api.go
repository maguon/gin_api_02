package admin

import (
	"gin_api_02/global"
	"gin_api_02/model/common/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetServerInfo
// @Tags      System
// @Summary   获取服务器信息
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取服务器信息"
// @Router    /admin/serverInfo [get]
func (s *AdminApi) GetServerInfo(c *gin.Context) {
	server, err := systemService.GetServerInfo()
	if err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"server": server}, "获取成功", c)
}

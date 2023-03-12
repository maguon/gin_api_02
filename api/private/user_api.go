package private

import (
	"fmt"
	"gin_api_02/global"
	"gin_api_02/model/common/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PrivateApi struct{}

// GetUserInfo
// @Tags      User
// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取用户信息"
// @Router    /private/user [get]
func (b *PrivateApi) GetUserInfo(c *gin.Context) {
	if claims, exists := c.Get("claims"); !exists {
		fmt.Println(exists)
	} else {
		fmt.Println(claims)
	}
	ReqUser, err := userService.GetUserInfo(1000)
	if err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "获取成功", c)
}

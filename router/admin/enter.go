package admin

import (
	"gin_api_02/api"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	AdminRouter
}

type AdminRouter struct{}

func (s *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	adminRouter := Router.Group("admin")
	adminApi := api.ApiGroupApp.AdminApiGroup.AdminApi
	{
		adminRouter.GET("sysUser", adminApi.GetAdminUserInfo)
		adminRouter.GET("sysUserList", adminApi.GetAdminUserList)
		adminRouter.POST("type", adminApi.CreateAdminType)
		adminRouter.POST("sysUser", adminApi.CreateAdminInfo)
		adminRouter.PUT("sysUser/:sysUserId", adminApi.UpdateAdmiInfo)
		adminRouter.PUT("type/:typeId", adminApi.UpdateAdminType)
		adminRouter.DELETE("type/:typeId", adminApi.RemoveAdminType)
		adminRouter.PUT("password", adminApi.UpdateAdminPassword)
		adminRouter.GET("type", adminApi.GetAdminType)
		adminRouter.GET("sysUser.csv", adminApi.ExportAdminInfo)
		adminRouter.GET("serverInfo", adminApi.GetServerInfo)

		adminRouter.POST("app", adminApi.CreateAppInfo)
		adminRouter.PUT("app/:appId", adminApi.UpdateAppInfo)
		adminRouter.DELETE("app/:appId", adminApi.RemoveAppInfo)
	}
	return adminRouter
}

package public

import (
	api "gin_api_02/api"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	PublicRouter
}

type PublicRouter struct{}

func (s *PublicRouter) InitPublicRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	publicRouter := Router.Group("public")
	publicApi := api.ApiGroupApp.PublicApiGroup.PublicApi
	{
		publicRouter.GET("captcha", publicApi.Captcha)
		publicRouter.POST("login", publicApi.Login)
		publicRouter.POST("adminLogin", publicApi.AdminLogin)
		publicRouter.GET("app", publicApi.GetAppInfo)
		publicRouter.GET("film", publicApi.GetFilmInfo)

	}
	return publicRouter
}

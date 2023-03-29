package public

import (
	api "gin_api_02/api"
	"gin_api_02/middleware"

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
		publicRouter.POST("register", publicApi.UserRegister)
		publicRouter.POST("adminLogin", publicApi.AdminLogin)
		publicRouter.GET("app", middleware.CheckCaptcha(), publicApi.GetAppInfo)
		publicRouter.GET("film", publicApi.GetFilmInfo)
		publicRouter.GET("actress", publicApi.GetActressList)

	}
	return publicRouter
}

package initialize

import (
	docs "gin_api_02/docs"
	"gin_api_02/global"
	"gin_api_02/middleware"
	"gin_api_02/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"PUT", "GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:  []string{"Content-Type", "auth-token"},
		ExposeHeaders: []string{"*"},
	}))
	publicRouter := router.RouterGroupApp.Public
	privateRouter := router.RouterGroupApp.Private
	adminRouter := router.RouterGroupApp.Admin
	PublicGroup := Router.Group("/api")
	{
		publicRouter.InitPublicRouter(PublicGroup)
	}
	PrivateGroup := Router.Group("/api")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		privateRouter.InitPrivateRouter(PrivateGroup)
	}
	AdminGroup := Router.Group("/api")
	AdminGroup.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"PUT", "GET", "POST", "PATCH"},
		AllowHeaders:  []string{"Content-Type", "auth-token"},
		ExposeHeaders: []string{"*"},
	}))
	//AdminGroup.Use(middleware.JWTAdminAuth())
	{
		adminRouter.InitAdminRouter(AdminGroup)
	}
	if global.SYS_CONFIG.System.Mode == "debug" {
		docs.SwaggerInfo.BasePath = "/api"
		Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return Router
}

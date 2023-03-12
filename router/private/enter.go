package private

import (
	api "gin_api_02/api"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	PrivateRouter
}

type PrivateRouter struct{}

func (s *PrivateRouter) InitPrivateRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	privateRouter := Router.Group("private")
	privateApi := api.ApiGroupApp.PrivateApiGroup.PrivateApi
	{
		privateRouter.GET("user", privateApi.GetUserInfo)

	}
	return privateRouter
}

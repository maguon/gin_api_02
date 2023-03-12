package router

import (
	"gin_api_02/router/admin"
	"gin_api_02/router/private"
	"gin_api_02/router/public"
)

type RouterGroup struct {
	Public  public.RouterGroup
	Private private.RouterGroup
	Admin   admin.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

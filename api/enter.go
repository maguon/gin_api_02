package api

import (
	"gin_api_02/api/admin"
	"gin_api_02/api/private"
	"gin_api_02/api/public"
)

type ApiGroup struct {
	PublicApiGroup  public.ApiGroup
	PrivateApiGroup private.ApiGroup
	AdminApiGroup   admin.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

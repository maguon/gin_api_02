package admin

import "gin_api_02/service"

type ApiGroup struct {
	AdminApi
}
type AdminApi struct{}

var (
	adminService  = service.ServiceGroupApp.AdminService
	appService    = service.ServiceGroupApp.AppService
	systemService = service.ServiceGroupApp.SystemService
)

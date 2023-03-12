package public

import "gin_api_02/service"

type ApiGroup struct {
	PublicApi
}

var (
	adminService = service.ServiceGroupApp.AdminService
	userService  = service.ServiceGroupApp.UserService
	jwtService   = service.ServiceGroupApp.JwtService
	appService   = service.ServiceGroupApp.AppService
	filmService  = service.ServiceGroupApp.FilmService
)

package service

type ServiceGroup struct {
	JwtService
	UserService
	AdminService
	AppService
	SystemService
	FilmService
}

var ServiceGroupApp = new(ServiceGroup)

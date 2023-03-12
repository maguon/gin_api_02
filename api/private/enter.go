package private

import "gin_api_02/service"

type ApiGroup struct {
	PrivateApi
}

var (
	userService = service.ServiceGroupApp.UserService
)

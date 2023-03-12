package main

import (
	"fmt"
	"gin_api_02/global"
	"gin_api_02/initialize"
	"strconv"
)

func main() {
	global.SYS_VP = initialize.Viper()
	fmt.Println(global.SYS_CONFIG.Zap.OutputPaths)
	global.SYS_LOG = initialize.Zap()
	initialize.Redis()
	initialize.Mongo()
	// global.SYS_DB = initialize.Gorm()

	Router := initialize.Routers()
	global.SYS_LOG.Info("server start ")
	Router.Run(":" + strconv.Itoa(global.SYS_CONFIG.System.Port))
}

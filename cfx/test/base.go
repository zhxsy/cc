package test

import (
	"cfx_server/cfx/server"
	"cfx_server/warehouses/app"
	"cfx_server/warehouses/test"
)

func init() {
	app.InitEnv("dev")
	//app.InitConfig("GAME_CENTER_ENV_PATH", "/Users/nck/workplace/cfx_server/api/game/config")
	app.InitConfig("", "/Users/nck/workplace/cfx_server/cfx/config")
	//app.InitLog("GAME_CENTER_ENV_PATH", "/Users/nck/workplace/cfx_server/api/game/logs")
	app.InitLog("", "/Users/nck/workplace/cfx_server/cfx/logs")

	//初始化路由相关的代码
	test.NewEngine(server.GetEngine)
}

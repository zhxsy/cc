package main

import (
	"cfx_server/cfx/server"
	"cfx_server/warehouses/app"
	"flag"
)

func main() {
	var env string
	flag.StringVar(&env, "env", "", "env: [dev|st|prd]")
	flag.Parse()

	//初始化env
	app.InitEnv(env)
	app.InitConfig("", "cfx/config")
	//初始化日志
	app.InitLog("", "cfx/logs")
	server.Graceful()
}

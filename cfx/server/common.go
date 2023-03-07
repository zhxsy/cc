package server

import (
	"cfx_server/cfx/middleware"
	"cfx_server/cfx/router"
	"cfx_server/warehouses/app"
	"cfx_server/warehouses/library/wrapper"
	"github.com/gin-gonic/gin"
)

//获取engine
func GetEngine() *gin.Engine {
	var engine *gin.Engine
	if app.IsPrd() {
		gin.SetMode(gin.ReleaseMode)
		engine = gin.New()
	} else {
		engine = gin.Default()
	}

	//初始化
	initMiddleware(engine)
	//路由
	router.InitRouter(engine)

	return engine
}

// initMiddleware 加载中间件
func initMiddleware(engine *gin.Engine) {
	engine.Use(wrapper.CheckTraceId())
	engine.Use(middleware.CorsHandle())
}

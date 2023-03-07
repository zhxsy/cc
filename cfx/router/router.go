package router

import (
	"cfx_server/cfx/controller/confulx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
	return
}

func InitRouter(engine *gin.Engine) {
	engine.GET("ping", Ping)

	//这里是提供的cfx的代码
	cfxEngine := engine.Group("/cfx")
	{
		cfxEngine.POST("/walletInfo", confulx.WalletInfo)
		cfxEngine.POST("/mintForUser", confulx.MintForUser)
		cfxEngine.POST("/nftTransfer", confulx.NFTTransfer)
	}

}

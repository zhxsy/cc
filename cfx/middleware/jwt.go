package middleware

import (
	"cfx_server/cfx/form"
	"cfx_server/warehouses/app"
	"cfx_server/warehouses/library/utils"
	"cfx_server/warehouses/output"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			output.Fail(c, form.ErrJwtErr.SetMsg("not login"))
			c.Abort()
			return
		}
		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				output.Fail(c, form.ErrJwtErr.SetMsg("Authorization has expired"))
				c.Abort()
				return
			}
			output.Fail(c, form.ErrJwtErr.SetMsg(err.Error()))
			c.Abort()
			return
		}
		// 用户被删除的逻辑 需要优化 此处比较消耗性能 如果需要 请自行打开
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + app.Config("app").GetInt64("expires_time")
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			//if ok := app.Config("app").GetBool("use_multipoint");ok {
			//	err, RedisJwtToken := jwtService.GetRedisJWT(newClaims.Username)
			//	if err != nil {
			//		global.GVA_LOG.Error("get redis jwt failed", zap.Error(err))
			//	} else { // 当之前的取成功时才进行拉黑操作
			//		_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken})
			//	}
			//	// 无论如何都要记录当前的活跃状态
			//	_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
			//}
		}
		c.Set("claims", claims)
		c.Next()
	}
}

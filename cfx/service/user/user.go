package user

import (
	"cfx_server/warehouses/app"
	"cfx_server/warehouses/library/utils"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Service struct {
}

// 签名
func (s *Service) tokenNext() (string, error) {
	j := utils.NewJWT()
	id := uint64(1)
	userName := "userName"
	baseClaims := utils.BaseClaims{
		ID:       id,
		Username: userName,
	}
	bufferTime := app.Config("app").GetInt64("buffer_time")
	expiresTime := app.Config("app").GetInt64("expires_time")
	customClaims := j.CreateClaims(baseClaims, bufferTime, jwt.StandardClaims{
		NotBefore: time.Now().Unix() - 1000,              // 签名生效时间
		ExpiresAt: time.Now().Unix() + expiresTime,       // 过期时间  配置文件
		Issuer:    app.Config("app").GetString("issuer"), // 签名的发行者
	})
	return j.CreateToken(customClaims)
}

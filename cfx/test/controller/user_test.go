package controller

import (
	_ "cfx_server/cfx/test"
	"cfx_server/warehouses/library/third_api"
	"cfx_server/warehouses/test"
	"testing"
)

func TestLogin(t *testing.T) {
	body := test.PostJson(
		"/user/notice",
		map[string]interface{}{
			"email":    "3264245531@qq.com",
			"password": "123456",
		},
	)

	body.Printf()
}

func TestRegister(t *testing.T) {
	body := test.PostJson(
		"/user/register",
		map[string]interface{}{
			"username":         "vic",
			"email":            "3264245531@qq.com",
			"password":         "123456",
			"confirm_password": "123456",
		},
	)

	body.Printf()
}

func TestInfo(t *testing.T) {
	head := &third_api.HeaderOption{
		Key: "x-token",
		Val: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJiYXNlX2NsYWltcyI6eyJJRCI6MSwiVXNlcm5hbWUiOiJ2aWMiLCJOaWNrTmFtZSI6IiJ9LCJidWZmZXJfdGltZSI6ODY0MDAsImp3dF9zdGFuZGFyZF9jbGFpbXMiOnsiZXhwIjoxNjYxODUxNDQ1LCJpc3MiOiJ2aWMiLCJuYmYiOjE2NjEyNDU2NDV9fQ.nhxqOX2w7uK1LYt-23opxxpO6t46OdVuE9eXxostk7Y",
	}
	body := test.GetHttp(
		"/user/info?id=1",
		head,
	)

	body.Printf()
}

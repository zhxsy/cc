package service

import (
	"cfx_server/cfx/service/user"
)

type ServGroup struct {
	UserServ user.Service
}

var ServGroupApp = new(ServGroup)

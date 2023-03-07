package form

import "cfx_server/warehouses/output"

var (
	// 用户
	ErrInconsistentPasswd = output.NewErr(-10000, "The password and confirm password are different", true)
	ErrRegisterFail       = output.NewErr(-10001, "Register fail", true)
	ErrRegisterAlready    = output.NewErr(-10002, "the phone already register", true)
	ErrPasswordErr        = output.NewErr(-10003, "The password or email is incorrect", true)
	ErrNotLogin           = output.NewErr(-10004, "Not login", true)
	// jwt
	ErrJwtErr = output.NewErr(-10100, "token fail", true)

	GenerateWalletErr = output.NewErr(-10200, "generate wallet error", true)
	MintForUser       = output.NewErr(-10201, "mintForUser error", true)
)

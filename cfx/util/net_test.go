package util

import (
	"cfx_server/cfx/util/net"
	"fmt"
	"testing"
)

func Test_net(t *testing.T) {
	url := "https://www.baidu.com"
	value, err := net.Get(url)
	fmt.Println(value, err)
}

func Test_tt(t *testing.T) {
	uri := "https://test.confluxrpc.com"
	var strs []string
	strs = append(strs, "0xf858ed00012ba2051f161bf1885361def794826bda2a8b447e443989b164f90a")
	mm := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "cfx_getTransactionByHash",
		"params":  strs,
		"id":      1,
	}
	res, err := net.Post(uri, mm)
	if err == nil {
		fmt.Println(res)
	}
}

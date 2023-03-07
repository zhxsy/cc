package rpc

import (
	"cfx_server/cfx/util/net"
	"fmt"
	"testing"
)

func Test_Ping(t *testing.T) {
	url := "http://127.0.0.1:9200/ping"
	var options []net.Option
	value := net.HeaderOption{Key: "key", Val: "value"}
	options = append(options, &value)
	result, err := Do("GET", url, nil, options...)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func Test_value3(t *testing.T) {
	url := "https://testnet.confluxscan.io/address/cfxtest:ach35y6ags44gf3k50pg2vu4v73ywe04xjjywpgv15?limit=10&reverse=true&skip=10"
	url = "https://testnet.snowtrace.io/api?module=logs&action=getLogs&topic0=0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef&apikey=FA5FATX6Z8GTBZXIP6GKQAIWFETV2QVGE4&"
	url = "https://testnet.confluxscan.io/v1/transaction?accountAddress=cfxtest:ach35y6ags44gf3k50pg2vu4v73ywe04xjjywpgv15&limit=10&reverse=true&skip=0"
	var options []net.Option
	value := net.HeaderOption{Key: "key", Val: "value"}
	options = append(options, &value)
	result, err := Do("GET", url, nil, options...)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

//cfx_call 这个也记录一下，后面的可能有用

//trace_transaction 这个也给记录一下，看里面各个字段的含义

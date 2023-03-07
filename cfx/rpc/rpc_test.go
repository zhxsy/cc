package rpc

import (
	"cfx_server/cfx/util"
	"encoding/json"
	"fmt"
	"testing"
)

//相关的文档
//https://developer.confluxnetwork.org/conflux-doc/docs/json_rpc/#cfx_getlogs

// Request
//curl -X POST
//--data '{"jsonrpc":"2.0","method":"cfx_getBalance","params":["cfx:type.user:aarc9abycue0hhzgyrr53m6cxedgccrmmyybjgh4xg", "latest_state"],"id":1}'
//-H "Content-Type: application/json"
//localhost:12539
//-x 请求方式 -data 请求体 -h 请求头 请求域名或者地址

//cfx_getBalance 获取账户的余额
func Test_getBalance(t *testing.T) {
	uri := "https://test.confluxrpc.com"
	var params []string
	params = append(params, "cfxtest:aak3rztcv1pvhf2pk02kt2egumj2dmv37upn5139gk") //钱包地址
	params = append(params, "latest_state")
	mm := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "cfx_getBalance",
		"params":  params,
		"id":      1,
	}
	result, err := Do("POST", uri, mm)
	if err != nil {
		t.Errorf("Test DoPost err:[%v]", err)
		return
	}
	str, _ := json.Marshal(result)
	fmt.Println(string(str))
	userBalance, _ := util.Hex2String(result.Result.(string))
	fmt.Printf("userBalnce:[%v] \n", userBalance)
}

func Test_getTransactionByHash(t *testing.T) {
	uri := "https://test.confluxrpc.com"
	var params []string
	//这个里面存储的是 hash 的交易码
	params = append(params, "0xf858ed00012ba2051f161bf1885361def794826bda2a8b447e443989b164f90a")
	mm := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "cfx_getTransactionByHash",
		"params":  params,
		"id":      1,
	}
	result, err := Do("POST", uri, mm)
	if err != nil {
		t.Errorf("Test DoPost err:[%v]", err)
		return
	}
	str, _ := json.Marshal(result)
	fmt.Println(string(str))
}

//这个记录一下，后续的可能用的到
//cfx_getLogs  这个里买的参数是map的方式
func Test_getLogs(t *testing.T) {
	uri := "https://test.confluxrpc.com"
	var params []map[string]string
	paramMap := map[string]string{}
	paramMap["fromEpoch"] = "0x6853800"
	paramMap["toEpoch"] = "0x6857ee9"
	paramMap["address"] = "cfxtest:acdc69pkv0f3yegsreh8dg6jbb9u23n5u6aydecrkm"
	paramMap["limit"] = "0x5"
	params = append(params, paramMap)
	mm := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "cfx_getLogs",
		"params":  params,
		"id":      1,
	}
	result, err := Do("POST", uri, mm)
	if err != nil {
		t.Errorf("Test DoPost err:[%v]", err)
		return
	}
	str, _ := json.Marshal(result)
	fmt.Println(string(str))
}

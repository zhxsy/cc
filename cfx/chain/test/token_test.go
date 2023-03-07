package cfx_test

import (
	"cfx_server/cfx/chain/cfxcontract"
	"cfx_server/cfx/chain/service"
	"encoding/json"
	"fmt"
	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/bind"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

//token 有三个方法，一个是设置值，一个是获取值，一个是转移
var (
	tokenContract = "cfxtest:acgj7b7x70u8d78fezmty3nnvmxgxwz2pewrfnm7rx"
	chainId       = uint(1)
)

func Test_TokenBalance(t *testing.T) {
	userAddr := "cfxtest:aak3rztcv1pvhf2pk02kt2egumj2dmv37upn5139gk"
	value, err := service.TokenGetBalance(userAddr, tokenContract)
	if err != nil {
		t.Errorf("Test NewToken err :%v", err)
	}
	fmt.Printf("userAddr:%v value:%v\n", userAddr, value)

	userAddr = "cfxtest:aarpdmpkj0ckhzmz7zp7yn25ec57dbsrcymfgj4r79"
	value, err = service.TokenGetBalance(userAddr, tokenContract)
	if err != nil {
		t.Errorf("Test NewToken err :%v", err)
	}
	fmt.Printf("userAddr:%v value:%v\n", userAddr, value)
}

//这里是设置值，使用默认的钱包方式
func Test_TokenValue(t *testing.T) {
	client, err := sdk.NewClient("https://test.confluxrpc.com", sdk.ClientOption{
		KeystorePath: "../keystore",
	})
	if err != nil {
		t.Errorf("Test TokenValue err:%v", err)
	}

	defaultAddr, _ := client.AccountManager.GetDefault()
	fmt.Println(defaultAddr)
	client.AccountManager.Unlock(*defaultAddr, "sznpassword")

	auth := &bind.TransactOpts{}
	//auth.From = defaultAddr
	//auth.GasPrice = types.NewBigInt(0x3b9aca00)
	//auth.Gas = types.NewBigInt(0x6bfb)
	auth.ChainID = types.NewUint(chainId)

	//合约
	contract := cfxaddress.MustNew(tokenContract)
	storageToken, err := cfxcontract.NewToken(contract, client)
	if err != nil {
		t.Errorf("Test NewToken err:%v", err)
	}
	value := big.NewInt(504)
	resp, hash, err := storageToken.MyToken(auth, value)
	if err != nil {
		t.Errorf("")
	}
	//答应相关的信息，以及hash
	resData, _ := json.Marshal(resp)
	fmt.Printf("resp %v\n", string(resData))
	fmt.Printf("hash:[%v]", hash)
}

//这里是用某个人的账号设置值
func Test_ComplexTokenValue(t *testing.T) {
	privateKey := "ff4366c5874602ff63be192cc1b6c38fc3cb3eb464d3805f64a0ece0c46b9b24"
	password := "sznpassword"
	resp, hash, err := service.SetTokenValue(privateKey, password, tokenContract, chainId)
	if err != nil {
		t.Errorf("Test SetTokenValue err:%v", err)
	}
	fmt.Printf("resp %v\n", resp)
	fmt.Printf("hash:[%v]\n", hash)
}

func Test_TokenTransfer(t *testing.T) {
	client, err := sdk.NewClient("https://test.confluxrpc.com", sdk.ClientOption{
		KeystorePath: "../keystore",
	})

	defaultAddr, _ := client.AccountManager.GetDefault()
	fmt.Println(defaultAddr)
	client.AccountManager.Unlock(*defaultAddr, "sznpassword")

	auth := &bind.TransactOpts{}
	auth.From = defaultAddr
	//auth.GasPrice = types.NewBigInt(0x3b9aca00)
	//auth.Gas = types.NewBigInt(0x6bfb)
	//auth.StorageLimit = types.NewUint64(300000)
	auth.ChainID = types.NewUint(chainId)

	//合约
	contract := cfxaddress.MustNew("cfxtest:acgj7b7x70u8d78fezmty3nnvmxgxwz2pewrfnm7rx")
	storageToken, err := cfxcontract.NewToken(contract, client)
	if err != nil {
		t.Errorf("Test NewToken err: %v ", err)
	}
	value := big.NewInt(20)
	add := cfxaddress.MustNewFromBase32("cfxtest:aarpdmpkj0ckhzmz7zp7yn25ec57dbsrcymfgj4r79")
	toAddr := common.HexToAddress(add.GetHexAddress())

	resp, hash, err := storageToken.Transfer(auth, toAddr, value)
	if err != nil {
		t.Errorf("Test Transfer err:%v", err)
	}
	//答应相关的信息，以及hash
	resData, _ := json.Marshal(resp)
	fmt.Printf("resp %v\n", string(resData))
	fmt.Printf("hash:[%v]", hash)
}

func Test_ComplexTokenTransfer(t *testing.T) {
	privateKey := "ff4366c5874602ff63be192cc1b6c38fc3cb3eb464d3805f64a0ece0c46b9b24"
	password := "sznpassword"
	contract := "cfxtest:acgj7b7x70u8d78fezmty3nnvmxgxwz2pewrfnm7rx"
	toUserAddr := "cfxtest:aak3rztcv1pvhf2pk02kt2egumj2dmv37upn5139gk"
	num := int64(40)
	resp, hash, err := service.Transfer(privateKey, password, contract, toUserAddr, num, chainId)
	if err != nil {
		t.Errorf("Test SetTokenValue err:%v", err)
	}
	fmt.Printf("resp %v\n", resp)
	fmt.Printf("hash:[%v]\n", hash)
}

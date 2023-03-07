package service

import (
	"cfx_server/cfx/chain/cfxcontract"
	"cfx_server/cfx/chain/client"
	"encoding/json"
	"fmt"
	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/bind"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func TokenGetBalance(userAddr, contactAddr string) (*big.Int, error) {
	value := big.NewInt(0)
	cli, err := sdk.NewClient("https://test.confluxrpc.com", sdk.ClientOption{
		KeystorePath: "../context/keystore",
	})
	if err != nil {
		return value, err
	}

	contract := cfxaddress.MustNew(contactAddr)
	storageToken, err := cfxcontract.NewToken(contract, cli)
	if err != nil {
		return value, err
	}

	add := cfxaddress.MustNewFromBase32(userAddr)
	findAddr := common.HexToAddress(add.GetHexAddress())
	//这里，打印出 标准的钱包地址 0x1396d5e28Dd913970C4dB097e086829181aA39EC
	fmt.Println(findAddr)
	resp, err := storageToken.BalanceOf(&bind.CallOpts{}, findAddr)
	if err != nil {
		return value, err
	}
	return resp, nil
}

func SetTokenValue(privateKey, password, contractAddr string, chainId uint) (string, string, error) {
	cli, addr, err := client.GetClient(privateKey, password)
	if err != nil {
		return "", "", err
	}
	fmt.Println(addr)
	defer func() {
		cli.AccountManager.Delete(*addr, password)
		cli.Close()
	}()
	cli.AccountManager.Unlock(*addr, password)

	auth := &bind.TransactOpts{}
	auth.From = addr
	auth.ChainID = types.NewUint(chainId)

	//合约
	contract := cfxaddress.MustNew(contractAddr)
	storageToken, err := cfxcontract.NewToken(contract, cli)
	if err != nil {
		return "", "", err
	}
	value := big.NewInt(505)
	resp, hash, err := storageToken.MyToken(auth, value)
	if err != nil {
		return "", "", err
	}
	//答应相关的信息，以及hash
	resData, _ := json.Marshal(resp)
	hashData, _ := json.Marshal(hash)
	return string(resData), string(hashData), err
}

func Transfer(privateKey, password, contractAddr string, toUserAddr string, num int64, chainId uint) (string, string, error) {
	cli, addr, err := client.GetClient(privateKey, password)
	if err != nil {
		return "", "", err
	}
	fmt.Println(addr)
	defer func() {
		cli.AccountManager.Delete(*addr, password)
		cli.Close()
	}()
	cli.AccountManager.Unlock(*addr, password)

	auth := &bind.TransactOpts{}
	auth.From = addr
	auth.ChainID = types.NewUint(chainId)

	//合约
	contract := cfxaddress.MustNew(contractAddr)
	storageToken, err := cfxcontract.NewToken(contract, cli)
	if err != nil {
		return "", "", err
	}
	value := big.NewInt(num)
	add := cfxaddress.MustNewFromBase32(toUserAddr)
	toAddr := common.HexToAddress(add.GetHexAddress())

	resp, hash, err := storageToken.Transfer(auth, toAddr, value)
	if err != nil {
		return "", "", err
	}
	//答应相关的信息，以及hash
	resData, _ := json.Marshal(resp)
	hashData, _ := json.Marshal(hash)
	return string(resData), string(hashData), err
}

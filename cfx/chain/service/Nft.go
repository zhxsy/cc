package service

import (
	"cfx_server/cfx/chain/cfxcontract"
	"cfx_server/cfx/chain/client"
	"cfx_server/warehouses/app"
	"fmt"
	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/bind"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
)

// NftGetBalance 获取ntf的个数
//获取链上的数据的时候，是不需要使用密钥的
func NftGetBalance(userAddr string) (*big.Int, error) {
	rpc := app.Config("app").GetString("cfx_rpc")
	contractAddr := app.Config("app").GetString("ntf_contract_address")
	client, err := sdk.NewClient(rpc)
	if err != nil {
		return big.NewInt(0), err
	}

	contract := cfxaddress.MustNew(contractAddr)
	storageToken, err := cfxcontract.NewMetaderbyHorseNFTToken(contract, client)
	if err != nil {
		return big.NewInt(0), err
	}

	add := cfxaddress.MustNewFromBase32(userAddr)
	findAddr := common.HexToAddress(add.GetHexAddress())
	value, err := storageToken.BalanceOf(&bind.CallOpts{}, findAddr)
	if err != nil {
		return big.NewInt(0), err
	}
	return value, nil
}

// MintForUserByDefault 这里使用默认的账户进行mint
// 涉及到初始化路径的问题，这个方法，暂时只能本地使用
func MintForUserByDefault(userAddr string, tokenId int64) (string, error) {
	password := app.Config("app").GetString("key_password")
	env := app.GetEnv()
	path := "cfx/keystore/" + env
	path = "/Users/nck/workplace/cfx_server/cfx/keystore/dev"
	cli, err := sdk.NewClient("https://test.confluxrpc.com", sdk.ClientOption{
		KeystorePath: path,
	})
	if err != nil {
		return "", err
	}
	defaultAddr, _ := cli.AccountManager.GetDefault()
	cli.AccountManager.Unlock(*defaultAddr, password)

	return MintForNFT(cli, userAddr, tokenId)
}

func MintForUserByPrivate(userAddr string, tokenId int64) (string, error) {
	privateKey := app.Config("app").GetString("private_key")
	password := app.Config("app").GetString("key_password")
	cli, addr, err := client.GetClient(privateKey, password)
	if err != nil {
		return "", err
	}
	defer func() {
		cli.AccountManager.Delete(*addr, password)
		cli.Close()
	}()
	cli.AccountManager.Unlock(*addr, password)
	return MintForNFT(cli, userAddr, tokenId)
}

func MintForNFT(cli *sdk.Client, userAddr string, tokenId int64) (string, error) {
	chainId := app.Config("app").GetInt32("chain_id")
	contractAddr := app.Config("app").GetString("ntf_contract_address")
	auth := &bind.TransactOpts{}
	auth.ChainID = types.NewUint(uint(chainId))

	//合约
	contract := cfxaddress.MustNew(contractAddr)
	storageToken, err := cfxcontract.NewMetaderbyHorseNFTToken(contract, cli)
	if err != nil {
		return "", err
	}
	value := big.NewInt(tokenId)

	add := cfxaddress.MustNewFromBase32(userAddr)
	toAddr := common.HexToAddress(add.GetHexAddress())
	fmt.Println(toAddr)
	log.Default().Printf("MintForUser userAddr:[%v] \n", userAddr)
	_, hash, err := storageToken.MintForUser(auth, value, toAddr)
	if err != nil {
		return "", err
	}
	return hash.String(), err
}

func NFTTransfer(privateKey string, toUserAddr string, tokenId int64) (string, error) {
	password := app.Config("app").GetString("key_password")
	chainId := app.Config("app").GetInt32("chain_id")
	contractAddr := app.Config("app").GetString("ntf_contract_address")
	cli, addr, err := client.GetClient(privateKey, password)
	if err != nil {
		return "", err
	}
	defer func() {
		cli.AccountManager.Delete(*addr, password)
		cli.Close()
	}()
	cli.AccountManager.Unlock(*addr, password)

	auth := &bind.TransactOpts{}
	auth.From = addr
	auth.ChainID = types.NewUint(uint(chainId))

	//合约
	contract := cfxaddress.MustNew(contractAddr)
	storageToken, err := cfxcontract.NewMetaderbyHorseNFTToken(contract, cli)
	if err != nil {
		return "", err
	}
	value := big.NewInt(tokenId)
	fromAddr := common.HexToAddress(addr.GetHexAddress())
	add := cfxaddress.MustNewFromBase32(toUserAddr)
	toAddr := common.HexToAddress(add.GetHexAddress())
	log.Default().Printf("NFTTransfer from:[%v] to [%v]", addr, add)
	_, hash, err := storageToken.TransferFrom(auth, fromAddr, toAddr, value)
	if err != nil {
		return "", err
	}
	//答应相关的信息，以及hash
	return hash.String(), err
}

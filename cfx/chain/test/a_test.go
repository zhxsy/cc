package cfx_test

import (
	"cfx_server/cfx/chain/cfxcontract"
	"fmt"
	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/bind"
	internalcontract "github.com/Conflux-Chain/go-conflux-sdk/contract_meta/internal_contract"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func Test_AdminControl(t *testing.T) {
	client, err := sdk.NewClient("https://test.confluxrpc.com", sdk.ClientOption{
		KeystorePath: "../context/keyStore",
	})
	if err != nil {
		fmt.Println(err)
	}

	defaultAddr, _ := client.AccountManager.GetDefault()
	fmt.Println(defaultAddr)
	//todo 这个查看下是用来做什么的
	client.AccountManager.Unlock(*defaultAddr, "258369Hua")

	//创建一个 adminControl的合约，把 client 给传递进去
	adminControl, _ := internalcontract.NewAdminControl(client)
	//合约
	contractAddr := cfxaddress.MustNew("cfxtest:accx9mzm610p6fbkp898sw8pzf7wamrn1jxc463brw")

	ops := types.ContractMethodCallOption{}

	addr, err := adminControl.GetAdmin(&ops, contractAddr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(addr)

	toAddr := cfxaddress.MustNewFromBase32("cfxtest:aarpdmpkj0ckhzmz7zp7yn25ec57dbsrcymfgj4r79")
	ops2 := types.ContractMethodSendOption{}

	hash, err := adminControl.SetAdmin(&ops2, contractAddr, toAddr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hash)

	addr, err = adminControl.GetAdmin(&ops, contractAddr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(addr)
}

//代付相关的合约
func Test_Sponsor(t *testing.T) {
	client, err := sdk.NewClient("https://test.confluxrpc.com", sdk.ClientOption{
		KeystorePath: "../context/keyStore",
	})
	if err != nil {
		fmt.Println(err)
	}

	//合约  这个合约是 7r49的
	contractAddr := cfxaddress.MustNew("cfxtest:acadkkt6x8y4kvmhpaua75h5ne9s9r0ucp02k1x29m")

	ops := types.ContractMethodCallOption{}
	adminControl, _ := internalcontract.NewAdminControl(client)
	sponsor, _ := internalcontract.NewSponsor(client)
	admin, err := adminControl.GetAdmin(&ops, contractAddr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(admin)

	addr, err := sponsor.GetSponsorForGas(&ops, contractAddr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(addr)
	fmt.Println(addr.GetHexAddress())

	add := cfxaddress.MustNewFromBase32(addr.String())
	userAddr := common.HexToAddress(add.GetHexAddress())
	fmt.Println(userAddr)

	//这个是抵押物相关的代码
	fmt.Println("抵押物")
	addr, _ = sponsor.GetSponsorForCollateral(&ops, contractAddr)
	userAddr = common.HexToAddress(addr.GetHexAddress())
	fmt.Println(userAddr)

	//获取gas 的balance
	value, _ := sponsor.GetSponsoredBalanceForGas(&ops, contractAddr)
	fmt.Println(value.String())

	value2 := big.NewInt(200)

	op := types.ContractMethodSendOption{}
	//auth.Nonce = types.NewBigInt(0x1b)
	op.GasPrice = types.NewBigInt(0x3b9aca00)
	op.Gas = types.NewBigInt(0x6bfb)
	//auth.Value =  types.NewBigInt(128)
	//auth.EpochHeight = types.NewUint64(105745985)
	//op.StorageLimit = types.NewUint64(300000)
	op.ChainID = types.NewUint(1)

	toAddr := cfxaddress.MustNew("cfxtest:aak3rztcv1pvhf2pk02kt2egumj2dmv37upn5139gk")

	hash, err := sponsor.SetSponsorForGas(&op, toAddr, value2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hash)

}

//代付相关的合约
func Test_Sponsor2(t *testing.T) {
	client, err := sdk.NewClient("https://test.confluxrpc.com", sdk.ClientOption{
		KeystorePath: "../context/keyStore",
	})
	if err != nil {
		fmt.Println(err)
	}

	//defaultAddr, _ := client.AccountManager.GetDefault()
	//fmt.Println(defaultAddr)
	//
	//client.AccountManager.Unlock(*defaultAddr, "258369Hua")

	//合约  这个合约是 7r49的
	contractAddr := cfxaddress.MustNew("cfxtest:acdv7e8fzn2k0hfa5bm181gtyxc1jb1h765dewfvt3")

	ops := types.ContractMethodCallOption{}
	adminControl, _ := internalcontract.NewAdminControl(client)
	sponsor, _ := internalcontract.NewSponsor(client)
	admin, err := adminControl.GetAdmin(&ops, contractAddr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(admin)

	addr, err := sponsor.GetSponsorForGas(&ops, contractAddr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(addr)
	fmt.Println(addr.GetHexAddress())

	add := cfxaddress.MustNewFromBase32(addr.String())
	userAddr := common.HexToAddress(add.GetHexAddress())
	fmt.Println(userAddr)

	//这个是抵押物相关的代码
	fmt.Println("抵押物")
	addr, _ = sponsor.GetSponsorForCollateral(&ops, contractAddr)
	userAddr = common.HexToAddress(addr.GetHexAddress())
	fmt.Println(userAddr)

	//获取gas 的balance
	value, _ := sponsor.GetSponsoredBalanceForGas(&ops, contractAddr)
	fmt.Println(value.String())
	//v := 9223372036854775807
	v := 1000 * 5000000000
	value2 := big.NewInt(int64(v))
	//value2.SetString(num,0)
	//fmt.Println(value2.String())

	va, _ := sponsor.GetSponsoredGasFeeUpperBound(&ops, contractAddr)
	fmt.Println(va.String())

	val, _ := sponsor.GetSponsoredBalanceForGas(&ops, contractAddr)
	fmt.Println(val)

	op := types.ContractMethodSendOption{}

	hash, err := sponsor.SetSponsorForGas(&op, contractAddr, value2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hash)
}

//2147483647
//-2147483648

//let upper_bound = min + ((max - min) as f64 * 0.8) as u64;

func TestMin(t *testing.T) {
}

func Test_Sponsor23(t *testing.T) {
	client, err := sdk.NewClient("https://test.confluxrpc.com", sdk.ClientOption{
		KeystorePath: "../context/keyStore",
	})
	if err != nil {
		fmt.Println(err)
	}

	userList := client.GetAccountManager().List()
	defaultAddr, _ := client.AccountManager.GetDefault()
	fmt.Println(defaultAddr)

	fmt.Println(userList)

	userAdd := userList[1]

	client.AccountManager.Unlock(userAdd, "258369Hua")

	//合约  这个合约是 7r49的
	contractAddr := cfxaddress.MustNew("cfxtest:acdv7e8fzn2k0hfa5bm181gtyxc1jb1h765dewfvt3")

	ops := types.ContractMethodCallOption{}

	adminControl, _ := internalcontract.NewAdminControl(client)
	sponsor, _ := internalcontract.NewSponsor(client)
	admin, err := adminControl.GetAdmin(&ops, contractAddr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(admin)

	addr, err := sponsor.GetSponsorForGas(&ops, contractAddr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(addr)
	fmt.Println(addr.GetHexAddress())

	add := cfxaddress.MustNewFromBase32(addr.String())
	userAddr := common.HexToAddress(add.GetHexAddress())
	fmt.Println(userAddr)

	//这个是抵押物相关的代码
	fmt.Println("抵押物")
	addr, _ = sponsor.GetSponsorForCollateral(&ops, contractAddr)
	userAddr = common.HexToAddress(addr.GetHexAddress())
	fmt.Println(userAddr)

	//获取gas 的balance
	//value, _ := sponsor.GetSponsoredBalanceForGas(&ops, contractAddr)
	//fmt.Println(value.String())
	//v := 9223372036854775807
	v := 10000000000000
	v = 500000 * 5 * 1200
	value2 := big.NewInt(int64(v))
	//value2.SetString(num,0)
	//fmt.Println(value2.String())

	//va, _ := sponsor.GetSponsoredGasFeeUpperBound(&ops, contractAddr)
	//fmt.Println(va.String())
	//
	//val,_ := sponsor.GetSponsoredBalanceForGas(&ops,contractAddr)
	//fmt.Println(val)

	op := types.ContractMethodSendOption{}
	//op.Value =  types.NewBigInt(400000)
	op.GasPrice = types.NewBigInt(5)
	op.Gas = types.NewBigInt(50000)
	//auth.StorageLimit = types.NewUint64(300000)
	//op.ChainID = types.NewUint(1)

	//无差别的存储代付：1 * 10 cfx； 无差别的燃气代付：5 * 1 Gdrip，
	//upper bound 500,000 drip（大概就是 Gas price为 1 ，Gas limit不超过 50 万）

	hash, err := sponsor.SetSponsorForGas(&op, contractAddr, value2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hash)

	var users []types.Address
	users = append(users, *defaultAddr)
}

//查看之前的配置，是否有这个人相关的配置，若是没有的话，给加载进来
//若是有的话，找个这个用户，然后，设置为默认的客户端，执行相关的代码，但是，这里是有一定的问题的
//前提 执行这个操作的时候，需要知道用户相关的钱包地址，以及用户的私钥

//这里尝试一下，第二种的方式，把用户相关的nft的类型，这里设置为任何人都可以进行转账的
//但是，还是有一个问题，那就是用户在执行一个操作的时候，还是需要用户相关的签名的操作的

//把tokenId 授权给所有的人，在代码操作的时候，需要继续，查看下，是否其他人能够执行这个 tokenId

func Test_pg(t *testing.T) {
	//KeystorePath 这个目录下，存储的是用户的相关信息
	client, err := sdk.NewClient("https://test.confluxrpc.com", sdk.ClientOption{
		KeystorePath: "../context/keystore",
	})

	fmt.Println(client.AccountManager.GetDefault())
	//client.AccountManager.Delete()

	//这里是合约地址

	contractAddr := cfxaddress.MustNew("cfxtest:acgj7b7x70u8d78fezmty3nnvmxgxwz2pewrfnm7rx")
	storageToken, err := cfxcontract.NewToken(contractAddr, client)
	if err != nil {
		panic(err)
	}

	add := cfxaddress.MustNewFromBase32("cfxtest:aak3rztcv1pvhf2pk02kt2egumj2dmv37upn5139gk")
	userAddr := common.HexToAddress(add.GetHexAddress())
	resp, err := storageToken.BalanceOf(&bind.CallOpts{}, userAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println("结果", resp)
}

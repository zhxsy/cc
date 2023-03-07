package cfx_test

import (
	"cfx_server/cfx/chain/generate"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

const (
	privateKeyStr = "0xd28edbdb7bbe75787b84c5f525f47666a3274bb06561581f00839645f3c26f66"
	publicKeyStr  = "0xc42b53ae2ef95fee489948d33df391c4a9da31b7a3e29cf772c24eb42f74e94ab3bfe00bf29a239c17786a5b921853b7c5344d36694db43aa849e401f91566a5"
	addressStr    = "0x1cecb4a2922b7007e236daf0c797de6e55496e84"
)

func TestPrivateKeyToPublic(t *testing.T) {
	expect := publicKeyStr
	pubKey := utils.PrivateKeyToPublicKey(privateKeyStr)
	if expect != pubKey {
		t.Errorf("Test PrivateKeyToPublicKey failed, expect %v, actual %v", expect, pubKey)
	}
}

func TestPublicKeyToAddr(t *testing.T) {
	expect := addressStr[2:]
	actual, err := utils.PublicKeyToCommonAddress(publicKeyStr)
	if err != nil {
		t.Fatal(err)
	}
	if expect != hex.EncodeToString(actual[:]) {
		t.Errorf("Test PublicKeyToAddr failed,except %v, actual %v", expect, actual)
	}
}

//=====================生成一个相关的私钥====================================================
//私钥可以生成公钥，然后，得到钱包地址

func Test_GenerateWallet(t *testing.T) {
	commAddr, pk, err := generate.GenerateKeyPair()
	if err != nil {
		t.Errorf("Test generateWallet err:%v", err)
	}
	fmt.Printf("addr:%v \n pk: %v \n", commAddr, pk)
	//GetCfxAddr(commAddr)
}

//"derby_hoof_coin_contract_addr": "0xee1c38135561e3cefaf30d84d22804f459a98d26",
//"derby_hoof_coin_contract_key": "6197a903105772358f355e6d7f2b9a46e40872d1eef6bee3c848b64d51a75d8d",

//0x9C0c35D8e1998fe104d8E4e5bF2629977c4E1295
//0x1c0c35D8e1998fe104D8e4e5bF2629977C4e1295
func Test_GetAddr(t *testing.T) {
	private := "0xae2e51689f785fac3f8f7093ee796facd2df8d30a23481507541aaeab743037e"
	private = "a871a4542bf668b40a902a18c8497926d6a81830da0815f0352ef91936677e20"
	//private = "6197a903105772358f355e6d7f2b9a46e40872d1eef6bee3c848b64d51a75d8d"
	//private = "d46e7e9cff4f5bc8f38183b06053171ddefb102362cbd0f438efed2af57abc19"
	commAddr, err := GetAddrByPrivateStr(private)
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	//0x57f367169c7183f5d7aec6e1d0f3ed951150f11e
	fmt.Printf("commonAddr:[%v]\n", commAddr)

	value, err := generate.GetCfxAddr(commAddr, 1)
	if err == nil {
		fmt.Printf("%v\n", value)
	}
	value, err = generate.GetCfxAddr(commAddr, 1029)
	if err == nil {
		fmt.Printf("%v\n", value)
	}
}

func GetAddrByPrivateStr(privateStr string) (string, error) {
	pubKey := utils.PrivateKeyToPublicKey(privateStr)
	actual, err := utils.PublicKeyToCommonAddress(pubKey)
	if err != nil {
		return "", err
	}
	return actual.String(), nil
}

func TestKeccak256(t *testing.T) {

	inputBytes := []byte{0x12, 0x34, 0x56, 0x78}
	hash := crypto.Keccak256(inputBytes)
	expect := "0x" + hex.EncodeToString(hash)
	// t.Error(expect)

	input := "0x12345678"
	actual, err := utils.Keccak256(input)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(actual)
	if actual != expect {
		t.Errorf("Test Keccak256 failed, expect %+v, actual %+v", expect, actual)
	}
}

//0x1396d5e28Dd913970C4dB097e086829181aA39EC
func Test_value3(t *testing.T) {
	add := cfxaddress.MustNewFromBase32("cfxtest:aak3rztcv1pvhf2pk02kt2egumj2dmv37upn5139gk")
	userAddr := common.HexToAddress(add.GetHexAddress())
	fmt.Println(userAddr)

	add2 := cfxaddress.MustNewFromBase32("cfx:aak3rztcv1pvhf2pk02kt2egumj2dmv37ugajh9zcd")
	userAddr2 := common.HexToAddress(add2.GetHexAddress())
	fmt.Println(userAddr2)
}

func GetCfxAddr(commAddr string) {
	fmt.Println(commAddr)
	//主网Id
	v, _ := cfxaddress.NewFromHex(commAddr, 1029)
	fmt.Println(v)

	//测试网id
	v, _ = cfxaddress.NewFromHex(commAddr, 1)
	fmt.Println(v)
}

func CreateKey(privateKey string) (addrs string) {
	//创建私钥
	//privateKey, err := crypto.GenerateKey()
	key, _ := crypto.HexToECDSA(privateKey)
	//privateKeyBytes := crypto.FromECDSA(key)
	//priv := hexutil.Encode(privateKeyBytes)[2:]
	publicKey := key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		//log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	//fmt.Println(address)
	fmt.Println(address)
	return address
}

//以太坊和 conflux 利用私钥转化出来的地址是有些不一样的
//这个是需要注意下的
func Test_Value(t *testing.T) {
	private := "d46e7e9cff4f5bc8f38183b06053171ddefb102362cbd0f438efed2af57abc19"
	CreateKey(private)
}

package generate

import (
	"cfx_server/cfx/common"
	"cfx_server/warehouses/app"
	"fmt"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/Conflux-Chain/go-conflux-sdk/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// GenerateKeyPair 用户的钱包地址和私钥
func GenerateKeyPair() (addr string, pk string, err error) {
	privateKey, _ := crypto.GenerateKey()
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateStr := hexutil.Encode(privateKeyBytes)
	pubKey := utils.PrivateKeyToPublicKey(privateStr)
	address, err := utils.PublicKeyToCommonAddress(pubKey)
	if err != nil {
		return "", "", fmt.Errorf("GenerateKeyPair PublicKeyToCommonAddress")
	}
	chainId := app.Config("app").GetInt32("chain_id")
	walletAddress := address.String()
	cfxUserAddr, err := GetCfxAddr(walletAddress, uint32(chainId))
	if err != nil {
		return "", "", err
	}
	return cfxUserAddr, privateStr, nil
}

// GetCfxAddr 使用commAddr获取用户的conflux地址，测试链或正式链上的地址
//这个 commAddr 是标准的钱包地址，例如 0x1fAbe8fDea8d0Af3EE582388dd6983Bc6123Fb3c
func GetCfxAddr(commAddr string, chainId uint32) (string, error) {
	if chainId != common.ChainIdTest && chainId != common.ChainIdFormal {
		return "", fmt.Errorf("chainId:[%v] not exists ", chainId)
	}
	v, err := cfxaddress.NewFromHex(commAddr, chainId)
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

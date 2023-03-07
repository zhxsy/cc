package client

import (
	"cfx_server/warehouses/app"
	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
)

func GetClient(privateKey string, password string) (*sdk.Client, *types.Address, error) {
	rpc := app.Config("app").GetString("cfx_rpc")
	path := "cfx/keystore/" + app.GetEnv()
	client, err := sdk.NewClient(rpc, sdk.ClientOption{
		KeystorePath: path,
	})
	if err != nil {
		return nil, nil, err
	}
	//第二个参数，这里是要和 unlocked 这里像配对的，两者的值不一样的话，签名是会产生错误的
	addr, err := client.AccountManager.ImportKey(privateKey, password)
	if err != nil {
		panic(err)
	}
	return client, &addr, nil
}

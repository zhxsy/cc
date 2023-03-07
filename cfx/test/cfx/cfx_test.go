package cfx

import (
	_ "cfx_server/cfx/test"
	"cfx_server/warehouses/test"
	"testing"
)

func Test_walletInfo(t *testing.T) {
	body := test.PostJson(
		"/cfx/walletInfo",
		map[string]interface{}{
			"phone": "15518956870",
		},
	)
	body.Printf()
}

func Test_MintForUser(t *testing.T) {
	//privateKey := "a871a4542bf668b40a902a18c8497926d6a81830da0815f0352ef91936677e20" //seed2
	//privateKey := "ff4366c5874602ff63be192cc1b6c38fc3cb3eb464d3805f64a0ece0c46b9b24" //seed1
	//privateKey = "666d1f143619e8704b6b7ea4b933361be77dc1d65c3acfa50705f05c7cf4fa73" //seed3
	userAddr := "cfxtest:aamprvzy4xenn7g0gu81s3vguj8e9jkep6xhvntzud"
	tokenId := int64(18)
	body := test.PostJson(
		"/cfx/mintForUser",
		map[string]interface{}{
			"user_addr": userAddr,
			"token_id":  tokenId,
		},
	)
	body.Printf()
}

func Test_NFTTransfer(t *testing.T) {
	privateKey := "666d1f143619e8704b6b7ea4b933361be77dc1d65c3acfa50705f05c7cf4fa73"
	privateKey = "9db0357bbfaa46ae422a81def866119ef3ea38287b8cd10b148e22c1ce53b0eb"
	toAddr := "cfxtest:aak3rztcv1pvhf2pk02kt2egumj2dmv37upn5139gk"
	tokenId := int64(5)
	body := test.PostJson(
		"/cfx/nftTransfer",
		map[string]interface{}{
			"private_key": privateKey,
			"to_addr":     toAddr,
			"token_id":    tokenId,
		},
	)
	body.Printf()
}

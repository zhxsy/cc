package confulx

import (
	"cfx_server/cfx/chain/generate"
	"cfx_server/cfx/chain/service"
	"cfx_server/cfx/form"
	"cfx_server/warehouses/output"
	"github.com/gin-gonic/gin"
	"log"
)

//获取钱包的信息
func WalletInfo(ctx *gin.Context) {
	req := &form.WalletInfoReq{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		output.Fail(ctx, output.ErrFormValidation.SetMsg(err.Error()).Log(err))
		return
	}

	walletInfo, privateKey, err := generate.GenerateKeyPair()
	if err != nil {
		output.Fail(ctx, form.GenerateWalletErr.SetMsg(err.Error()).Log(err))
		return
	}

	res := form.WalletInfoRes{}
	res.Wallet = walletInfo
	res.PrivateKey = privateKey
	output.Ok(ctx, res)
	log.Default().Printf("WalletInfo phone [%v] res [%v]\n", req.Phone, res)
}

func MintForUser(ctx *gin.Context) {
	req := &form.MintForUserReq{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		output.Fail(ctx, output.ErrFormValidation.SetMsg(err.Error()).Log(err))
		return
	}
	log.Default().Printf("MintForUser key:[%v] tokenId:[%v]\n", req.UserAddr, req.TokenId)

	hash, err := service.MintForUserByPrivate(req.UserAddr, req.TokenId)
	if err != nil {
		output.Fail(ctx, form.MintForUser.SetMsg(err.Error()).Log(err))
		return
	}

	res := form.MintForUserRes{}
	res.Hash = hash
	output.Ok(ctx, res)
	log.Default().Printf("MintForUser success hash:[%v]", hash)
}

func NFTTransfer(ctx *gin.Context) {
	req := &form.NFTTransferReq{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		output.Fail(ctx, output.ErrFormValidation.SetMsg(err.Error()).Log(err))
		return
	}
	log.Default().Printf("NFTTransfer key:[%v] toAddr:[%v] tokenId:[%v]\n", req.PrivateKey, req.ToAddr, req.TokenId)

	hash, err := service.NFTTransfer(req.PrivateKey, req.ToAddr, req.TokenId)
	if err != nil {
		output.Fail(ctx, form.MintForUser.SetMsg(err.Error()).Log(err))
		return
	}

	res := form.NFTTransferRes{}
	res.Hash = hash
	output.Ok(ctx, res)
	log.Default().Printf("NFTTransfer success hash:[%v]", hash)
}

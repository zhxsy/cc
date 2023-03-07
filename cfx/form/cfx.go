package form

type WalletInitReq struct {
	PrivateKey string `json:"private_key"` //用户的私钥
}
type WalletInitRes struct {
	Wallet string `json:"wallet"` //用户的钱包地址
}

type WalletInfoReq struct {
	Phone string `json:"phone" binding:"required"`
}

type WalletInfoRes struct {
	Wallet     string `json:"wallet"`      //用户的钱包地址 cfx格式的地址
	PrivateKey string `json:"private_key"` //用户的私钥
}

type MintForUserReq struct {
	UserAddr string `json:"user_addr" binding:"required"` //用户的私钥
	TokenId  int64  `json:"token_id"`                     //mint的tokenId，对于合约来说是唯一的
}

type MintForUserRes struct {
	Hash string `json:"hash"`
}

type NFTTransferReq struct {
	PrivateKey string `json:"private_key" binding:"required"` //用户的私钥
	ToAddr     string `json:"to_addr"  binding:"required"`    //转移的地址 cfx格式的地址
	TokenId    int64  `json:"token_id"`                       //mint的tokenId，对于合约来说是唯一的
}

type NFTTransferRes struct {
	Hash string `json:"hash"`
}

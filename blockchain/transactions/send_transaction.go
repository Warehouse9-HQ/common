package transactions

type Send struct {
	WalletFrom string `json:"wallet_from"`
	WalletTo   string `json:"wallet_to"`

	Amount int64  `json:"amount"`
	Hash   string `json:"hash"`

	Signature string `json:"signature"`
}

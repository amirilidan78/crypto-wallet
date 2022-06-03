package wallet

import (
	"crypto-wallet/pkg/twallet"
	"crypto-wallet/pkg/uuid"
)

type Wallet interface {
	GenerateWallet() (string, string)
	AddressIsValid(blockchain string, address string) bool
	GetAddress(seed string, passphrases string, coin string) string
	GetPrivateKey(seed string, passphrases string, coin string) string
}

type wallet struct {
	tw twallet.TWallet
}

func (w wallet) GenerateWallet() (string, string) {
	passphrases := uuid.NewUId()
	seed := w.tw.GenerateHDWallet(passphrases)
	return passphrases, seed
}

func (w wallet) GetAddress(seed string, passphrases string, coin string) string {
	return w.tw.GetAddressForCoin(seed, passphrases, coin)
}

func (w wallet) GetPrivateKey(seed string, passphrases string, coin string) string {
	return w.tw.GetPrivateKeyForCoin(seed, passphrases, coin)
}

func (w wallet) AddressIsValid(blockchain string, address string) bool {
	return w.tw.IsAddressValid(address, blockchain)
}

func NewWalletService(tw twallet.TWallet) Wallet {
	return &wallet{tw}
}

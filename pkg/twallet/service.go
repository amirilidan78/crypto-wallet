package twallet

import (
	"github.com/golang/protobuf/proto"
)

type TWallet interface {
	GenerateHDWallet(passphrase string) string
	GetAddressForCoin(seed string, passphrase string, coin string) string
	GetPrivateKeyForCoin(seed string, passphrase string, coin string) string
	GetTransactionInputScriptForCoin(address string, coin string) []byte
	SignTransaction(coin string, pb proto.Message) ([]byte, error)
	IsAddressValid(address string, coin string) bool
}

type tWallet struct {
}

func (s *tWallet) GenerateHDWallet(passphrase string) string {
	return hDWalletKey(passphrase)
}

func (s *tWallet) GetAddressForCoin(seed string, passphrase string, coin string) string {
	return hdWalletAddressForCoin(seed, passphrase, coin)
}

func (s *tWallet) GetPrivateKeyForCoin(seed string, passphrase string, coin string) string {
	return hdWalletPrivateKeyForCoin(seed, passphrase, coin)
}

func (s *tWallet) GetTransactionInputScriptForCoin(address string, coin string) []byte {
	return hdWalletScriptBuildForAddress(address, coin)
}

func (s *tWallet) SignTransaction(coin string, pb proto.Message) ([]byte, error) {
	return signTransaction(coin, pb)
}

func (s *tWallet) IsAddressValid(address string, coin string) bool {
	return isAddressValid(address, coin)
}

func NewTWallet() TWallet {
	return &tWallet{}
}

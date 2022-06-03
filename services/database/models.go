package database

import "crypto-wallet/services/wallet"

type WalletModel struct {
	ID         uint   `gorm:"column:id"`
	Username   string `gorm:"column:username"`
	Seed       string `gorm:"column:seed"`
	Passphrase string `gorm:"column:passphrase"`
}

func (w *WalletModel) GetAddress(wallet wallet.Wallet, coin string) string {
	return wallet.GetAddress(w.Seed, w.Passphrase, coin)
}

func (w *WalletModel) GetPrivateKey(wallet wallet.Wallet, coin string) string {
	return wallet.GetPrivateKey(w.Seed, w.Passphrase, coin)
}

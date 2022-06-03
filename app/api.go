package main

import (
	"crypto-wallet/config"
	"crypto-wallet/pkg/twallet"
	"crypto-wallet/services/api"
	"crypto-wallet/services/database"
	"crypto-wallet/services/logger"
	"crypto-wallet/services/wallet"
)

func main() {
	c := config.NewConfig()
	l := logger.NewLogService(c)
	db := database.NewDatabaseService(c, l)
	tw := twallet.NewTWallet()
	w := wallet.NewWalletService(tw)
	service := api.NewApiService(c, db, w)
	service.Serve()
}

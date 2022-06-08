package main

import (
	"crypto-wallet/config"
	"crypto-wallet/pkg/blockbook"
	"crypto-wallet/pkg/httpClient"
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
	hc := httpClient.NewHttpClient()
	bb := blockbook.NewHttpBlockBookService(c, hc)
	tw := twallet.NewTWallet()
	w := wallet.NewWalletService(tw)
	aw := wallet.NewApiWalletService(c, bb)
	service := api.NewApiService(c, db, w, aw)
	service.Serve()
}

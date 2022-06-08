package wallet

import (
	"crypto-wallet/config"
	"crypto-wallet/pkg/blockbook"
	"fmt"
	"math"
	"strconv"
)

type ApiWallet interface {
	GetWalletBalance(blockchain string, address string) (float64, error)
	GetWalletTransactions(blockchain string, address string) ([]string, error)
	GetTransaction(blockchain string, txId string) (blockbook.TransactionResponse, error)
}

type apiWallet struct {
	c  config.Config
	bb blockbook.HttpBlockBook
}

func (w *apiWallet) GetWalletBalance(blockchain string, address string) (float64, error) {

	resp, err := w.bb.GetAddress(blockchain, address)

	fmt.Println(resp)
	fmt.Println(err)

	if err != nil {
		return 0, err
	}

	balance, errParse := strconv.ParseFloat(resp.Balance, 32)

	fmt.Println(errParse)

	if errParse != nil {
		return 0, errParse
	}

	subAmount := w.c.GetInt("coins." + blockchain + ".subAmount")

	balanceInCurrency := balance / math.Pow10(subAmount)

	return balanceInCurrency, nil
}

func (w *apiWallet) GetWalletTransactions(blockchain string, address string) ([]string, error) {

	resp, err := w.bb.GetAddress(blockchain, address)

	if err != nil {
		return nil, err
	}
	return resp.TxIds, nil
}

func (w *apiWallet) GetTransaction(blockchain string, txId string) (blockbook.TransactionResponse, error) {

	resp, err := w.bb.GetTransaction(blockchain, txId)

	if err != nil {
		return blockbook.TransactionResponse{}, err
	}

	return resp, nil
}

func NewApiWalletService(c config.Config, bb blockbook.HttpBlockBook) ApiWallet {
	return &apiWallet{c, bb}
}

package blockbook

import (
	"crypto-wallet/config"
	"crypto-wallet/pkg/httpClient"
)

type HttpBlockBook interface {
	GetAddress(coin string, address string) (AddressResponse, error)
	GetTransaction(coin string, txId string) (TransactionResponse, error)
}

type httpBlockBook struct {
	c  config.Config
	hc httpClient.HttpClient
}

func (b *httpBlockBook) getHost(coin string) string {

	url := b.c.GetString("coins." + coin + ".node")

	if url == "" {
		panic("error in getting block book node")
	}

	return url
}

func (b *httpBlockBook) get(coin string, path string, res interface{}) error {

	host := b.getHost(coin)

	url := host + path

	err := b.hc.SimpleGet(url, res)

	return err
}

func (b *httpBlockBook) GetAddress(coin string, address string) (AddressResponse, error) {

	res := AddressResponse{}

	path := AddressPath + address

	err := b.get(coin, path, &res)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (b *httpBlockBook) GetTransaction(coin string, txId string) (TransactionResponse, error) {

	res := TransactionResponse{}

	path := TransactionPath + txId

	err := b.get(coin, path, &res)

	if err != nil {
		return res, err
	}

	return res, nil
}

func NewHttpBlockBookService(c config.Config, hc httpClient.HttpClient) HttpBlockBook {
	return &httpBlockBook{c, hc}
}

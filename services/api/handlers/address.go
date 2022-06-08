package handlers

import (
	"crypto-wallet/pkg/blockbook"
	"crypto-wallet/services/database"
	"crypto-wallet/services/wallet"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type newAddressResponse struct {
	Address    string `json:"address"`
	PrivateKey string `json:"privateKey"`
}

type addressBalanceResponse struct {
	Balance float64 `json:"balance"`
}

type addressTransactionsResponse struct {
	Transactions []string `json:"transactions"`
}

type transactionResponse struct {
	Transaction blockbook.TransactionResponse `json:"Transaction"`
}

type errorResponse struct {
	Error string `json:"error"`
}

type newAddressParams struct {
	Coin     string `json:"coin" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type addressBalanceParams struct {
	Blockchain string `json:"blockchain" binding:"required"`
	Address    string `json:"address" binding:"required"`
}

type addressTransactionsParams struct {
	Blockchain string `json:"blockchain" binding:"required"`
	Address    string `json:"address" binding:"required"`
}

type transactionParams struct {
	Blockchain string `json:"blockchain" binding:"required"`
	TxId       string `json:"txId" binding:"required"`
}

func NewAddress(db database.Database, w wallet.Wallet) gin.HandlerFunc {
	return func(c *gin.Context) {
		var walletModel database.WalletModel
		p := newAddressParams{}
		err := c.BindJSON(&p)

		if err != nil {
			c.JSON(http.StatusBadRequest, errorResponse{
				Error: "coin and username is required provider",
			})
			return
		}

		username := p.Username
		coin := strings.ToUpper(p.Coin)

		if db.WalletExist(username) {

			walletModel, err = db.GetWallet(username)

			if err != nil {
				c.JSON(http.StatusBadRequest, errorResponse{
					Error: err.Error(),
				})
				return
			}

		} else {

			passphrases, seed := w.GenerateWallet()
			walletModel, err = db.StoreWallet(username, seed, passphrases)

			if err != nil {
				c.JSON(http.StatusBadRequest, errorResponse{
					Error: err.Error(),
				})
				return
			}
		}

		c.JSON(http.StatusOK, newAddressResponse{
			Address:    walletModel.GetAddress(w, coin),
			PrivateKey: walletModel.GetPrivateKey(w, coin),
		})
		return
	}
}

func AddressBalance(w wallet.ApiWallet) gin.HandlerFunc {
	return func(c *gin.Context) {

		p := addressBalanceParams{}
		err := c.BindJSON(&p)

		if err != nil {
			c.JSON(http.StatusBadRequest, errorResponse{
				Error: "blockchain and address is required",
			})
			return
		}

		Balance, err := w.GetWalletBalance(p.Blockchain, p.Address)

		if err != nil {
			c.JSON(http.StatusBadRequest, errorResponse{
				Error: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, addressBalanceResponse{
			Balance: Balance,
		})
		return
	}
}

func AddressTransactions(w wallet.ApiWallet) gin.HandlerFunc {
	return func(c *gin.Context) {

		p := addressTransactionsParams{}
		err := c.BindJSON(&p)

		if err != nil {
			c.JSON(http.StatusBadRequest, errorResponse{
				Error: "blockchain and address is required",
			})
			return
		}

		transactions, err := w.GetWalletTransactions(p.Blockchain, p.Address)

		if err != nil {
			c.JSON(http.StatusBadRequest, errorResponse{
				Error: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, addressTransactionsResponse{
			Transactions: transactions,
		})
		return
	}
}

func Transaction(w wallet.ApiWallet) gin.HandlerFunc {
	return func(c *gin.Context) {

		p := transactionParams{}
		err := c.BindJSON(&p)

		if err != nil {
			c.JSON(http.StatusBadRequest, errorResponse{
				Error: "blockchain and txId is required",
			})
			return
		}

		resp, err := w.GetTransaction(p.Blockchain, p.TxId)

		if err != nil {

			if err != nil {
				c.JSON(http.StatusBadRequest, errorResponse{
					Error: err.Error(),
				})
				return
			}

		}

		c.JSON(http.StatusOK, transactionResponse{
			Transaction: resp,
		})
		return
	}
}

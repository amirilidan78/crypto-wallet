package handlers

import (
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

type errorResponse struct {
	Error string `json:"error"`
}

type newAddressParams struct {
	Coin     string `json:"coin" binding:"required"`
	Username string `json:"username" binding:"required"`
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

package database

import (
	"crypto-wallet/config"
	"crypto-wallet/pkg/sql"
	"crypto-wallet/services/logger"
	"gorm.io/gorm"
)

type Database interface {
	Migrate()
	GetWallet(username string) (WalletModel, error)
	WalletExist(username string) bool
	StoreWallet(username string, seed string, passphrase string) (WalletModel, error)
}

type database struct {
	db  *gorm.DB
	log logger.LogService
}

func (d *database) GetWallet(username string) (WalletModel, error) {

	var wallet WalletModel

	result := d.db.Where("username = ?", username).First(&wallet)

	return wallet, result.Error
}

func (d *database) WalletExist(username string) bool {

	var exists bool

	err := d.db.Model(WalletModel{}).Select("count(*) > 0").Where("username = ?", username).Find(&exists).Error

	if err != nil {
		return true
	}

	return exists
}

func (d *database) StoreWallet(username string, seed string, passphrase string) (WalletModel, error) {

	model := WalletModel{
		Username:   username,
		Seed:       seed,
		Passphrase: passphrase,
	}

	result := d.db.Create(&model)

	return model, result.Error
}

func (d *database) Migrate() {
	d.db.AutoMigrate(&WalletModel{})
}

func NewDatabaseService(c config.Config, log logger.LogService) Database {
	return &database{sql.NewDatabaseConnection(c.GetString("database.dsn")), log}
}

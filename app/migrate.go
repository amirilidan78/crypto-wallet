package main

import (
	"crypto-wallet/config"
	"crypto-wallet/services/database"
	"crypto-wallet/services/logger"
	"log"
)

func main() {

	c := config.NewConfig()
	l := logger.NewLogService(c)
	db := database.NewDatabaseService(c, l)
	db.Migrate()
	log.Println("Migrated successfully")

}

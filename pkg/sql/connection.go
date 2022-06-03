package sql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewDatabaseConnection(dsn string) *gorm.DB {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

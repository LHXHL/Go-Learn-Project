package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func Connect() {
	open, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_URL")), &gorm.Config{})
	if err != nil {
		panic("Failed to connect...")
	} else {
		log.Println("Connect success!")
	}
	db = open
}

func GetDB() *gorm.DB {
	return db
}

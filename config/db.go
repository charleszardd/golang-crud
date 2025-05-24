package config 

import (
	"fmt"
	"os"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", 
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))

	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Panicf("Could not connect to the database: %v", err)
	}
	
	DB = db
	fmt.Println("Connected to the database")
}
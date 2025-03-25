package db

import (
	"fmt"
	"log"
	dbmodels "stock_broker_application/models"
	"stock_broker_application/src/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	cfg := utils.AppConfig.Database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to Connect to DB: %v", err)
	}
	fmt.Println("Connected to DB")

	DB.AutoMigrate(&dbmodels.User{})
}

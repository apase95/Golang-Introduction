package config

import (
	"fmt"
	"log"
	"mvc-orm/internal/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {
	if err := godotenv.Load(); err != nil {
		log.Println("[INFO] No .env file found!")
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("[ERROR] Failed to connect to database: %v", err)
	}
	err = database.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("[ERROR] Failed to auto-migrate database schema: %v", err)
	}

	DB = database
	fmt.Println("[SUCCESS] Connected to PostgresSQL & Migrated Schema successfully!")
}
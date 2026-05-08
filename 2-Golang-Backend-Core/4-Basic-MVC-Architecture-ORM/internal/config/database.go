package config

import (
	"fmt"
	"log"
	"mvc-orm/internal/models"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {
	if err := godotenv.Load(".env.example"); err != nil {
		log.Println("[INFO] No .env.example file found!")
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	customLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Error,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: customLogger,
	})
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
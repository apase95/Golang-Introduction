package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	AppEnv 		string 
	ServerPort 	string
	DBHost 		string
	DBPort 		int
	DBName 		string
	DBUser 		string
	DBPassword 	string
	AuthKey 	string
}

func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	strValue := getEnv(key, "")
	if value, err := strconv.Atoi(strValue); err == nil {
		return value
	}
	return fallback
}

func LoadConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Println("[INFO] No .env file found")
	}

	return &AppConfig{
		AppEnv: 	getEnv("APP_ENV", "development"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
		DBHost: 	getEnv("DB_HOST", "localhost"),
		DBPort: 	getEnvAsInt("DB_PORT", 5432),
		DBName: 	getEnv("DB_NAME", "Golang-Mini-Runner"),
		DBUser: 	getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "supersecret"),
		AuthKey: 	getEnv("AUTH_KEY", "supersecret"),
	}
}

func PrintConfig(config *AppConfig) {
	fmt.Printf("[ENV] Environment : %s\n", config.AppEnv)
	fmt.Printf("[ENV] Server Port : %s\n", config.ServerPort)
	fmt.Printf("[ENV] DB Host     : %s\n", config.DBHost)
	fmt.Printf("[ENV] DB Port     : %d\n", config.DBPort)
	fmt.Printf("[ENV] DB User     : %s\n", config.DBUser)
	fmt.Printf("[ENV] DB User     : %s\n", config.DBName)
	if config.AppEnv != "production" {
		fmt.Printf("[ENV] DB Password : %s\n", config.DBPassword)
		fmt.Printf("[ENV] Auth Key    : %s\n", config.AuthKey)
	}
}
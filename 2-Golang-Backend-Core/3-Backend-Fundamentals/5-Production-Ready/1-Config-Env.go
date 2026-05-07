package main

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
	DBUser 		string
	DBPassword 	string
}

func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}; return fallback
}

func getEnvAsInt(key string, fallback int) int {
	strValue := getEnv(key, "")
	if value, err := strconv.Atoi(strValue); err == nil {
		return value
	}; return fallback
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
		DBUser: 	getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "supersecret"),
	}
}

func main() {
	config := LoadConfig()
	fmt.Printf("[CONFIG] Environment : %s\n", config.AppEnv)
	fmt.Printf("[CONFIG] Server Port : %s\n", config.ServerPort)
	fmt.Printf("[CONFIG] DB Host     : %s\n", config.DBHost)
	fmt.Printf("[CONFIG] DB Port     : %d\n", config.DBPort)
	fmt.Printf("[CONFIG] DB User     : %s\n", config.DBUser)
	if config.AppEnv != "production" {
		fmt.Printf("[CONFIG] DB Password : %s\n", config.DBPassword)
	}
}

/*
Path: Golang-Backend-Core/3.Backend-Fundamentals/5-Production-Ready/Config-Env.go
Test 1:
```bash
go run Config-Env.go
```

Test 2:
```bash
echo -e "APP_ENV=production\nSERVER_PORT=8081\nDB_USER=admin\nDB_PASSWORD=my_secure_pass" > .env
	go run Config-Env.go
```

Test 3:
```bash
SERVER_PORT=8082 DB_HOST=aws-rds.amazon.com go run Config-Env.go
``` 
*/
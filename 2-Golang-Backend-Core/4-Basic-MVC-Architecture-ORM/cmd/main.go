package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"mvc-orm/internal/config"
	"mvc-orm/internal/routes"
)

func main() {
	fmt.Println("[START] Initialing System...")
	config.ConnectDatabase()
	fmt.Println("[FINISH] Database check")

	mux := http.NewServeMux()
	routes.SetupRoutes(mux)
	port := os.Getenv("PORT")
	if port == "" { port = "8080" }
	
	serverAddr := ":" + port
	fmt.Printf("🚀 Server is running on http://localhost%s\n", serverAddr)
	if err := http.ListenAndServe(serverAddr, mux); err != nil {
		log.Fatalf("[FATAL] Server crashed: %v", err)
	}
}

/*
Path: Golang-Backend-Core/3.Backend-Fundamentals/4-Basic-MVC-Architecture-ORM/cmd/main.go
```bash
go mod init mvc-orm
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/joho/godotenv
go get golang.org/x/crypto/bcrypt

docker stop $(docker ps -aq)
[OPT]: docker rm $(docker ps -aq)

docker run --name pg-mvc -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=user_management -p 5432:5432 -d postgres

go run cmd/main.go
curl -X POST http://localhost:8081/api/v1/users/register \
-H "Content-Type: application/json" \
-d '{"name": "Nooby", "email": "nooby_admin@gmail.com", "password": "supersecretpassword"}'
```
*/
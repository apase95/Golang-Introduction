package main

import (
	"fmt"
	"time"
)

type BaseModel struct {
	ID 			string
	Status 		string
	CreatedAt 	time.Time
}

func (b *BaseModel) UpdateStatus(newStatus string) {
	b.Status = newStatus
	fmt.Printf("[%s] Status updated to: %s\n", b.ID, b.Status)
}

type Server struct {
	BaseModel	//Embedding
	IPAddress 	string
	OS 			string
}

type Database struct { 
	BaseModel	//Embedding
	Engine 		string
	Port 		int
}

func main() {
	webServer := Server{
		BaseModel: BaseModel{
			ID: "hdtd-web-01",
			Status: "Starting",
			CreatedAt: time.Now(),
		},
		IPAddress: "192.168.1.14",
		OS: "Fedora 43",
	}
	fmt.Printf("Server ID: %s, IP: %s, OS: %s\n", webServer.ID, webServer.IPAddress, webServer.OS)

	webServer.UpdateStatus("Running")
	fmt.Println("---")

	pgDB := Database{
		BaseModel: BaseModel{
			ID: "hdtd-db-01",
			Status: "Pending",
			CreatedAt: time.Now(),
		},
		Engine: "PostgreSQL",
		Port: 5432,
	}
	fmt.Printf("Database ID: %s, Engine: %s, Port: %d\n", pgDB.ID, pgDB.Engine, pgDB.Port)
	pgDB.UpdateStatus("Connected")
}
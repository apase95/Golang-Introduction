package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

const dsn2 = "host=localhost port=5432 user=admin password=secret dbname=jobdb sslmode=disable"

func monitorPoolStats(db *sql.DB) {
	for {
		stats := db.Stats()
		fmt.Printf("   [STATS] InUse: %d | Idle: %d | WaitCount: %d (People waiting for connection)\n", stats.InUse, stats.Idle, stats.WaitCount)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	fmt.Println(" === CONFIGURING DATABASE CONNECTION POOL ===")
	
	// Initial Connection Pool
	db, err := sql.Open("postgres", dsn2)
	if err != nil {
		log.Fatalf("Failed to open DB pool: %v", err)
	}; defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatalf("Database connection failed; %v", err)
	}

	// Set Rules for Connection pool
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(10 * time.Minute)
	fmt.Println("[SUCCESS] Connection Pool configured with MaxOpen=5, MaxIdle=5")

	// Simulating high traffic
	fmt.Println("[SIMULATION]")
	var wg sync.WaitGroup
	totalRequests := 20
	go monitorPoolStats(db)
	
	startTime := time.Now()
	for i := 1; i <= totalRequests; i++ {
		wg.Add(1)
		go func(requestID int) {
			defer wg.Done()
			_, err := db.Exec("SELECT pg_sleep(1)")
			if err != nil {
				log.Printf("[REQ %02d] Query failed: %v\n", requestID, err)
				return
			}; fmt.Printf("[REQ %02d] Completed!\n", requestID)
		}(i)
	}; wg.Wait()
	fmt.Printf("\n[FINISHED IN] %v\n", totalRequests, time.Since(startTime))
}

/*
Path: Golang-Backend-Core/3.Backend-Fundamentals/4-Database/Postgres-DB-Conn.go
```bash
go get github.com/lib/pq
docker run --name pg-test -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=jobdb -p 5432:5432 -d postgres
go run Postgres-DB-Conn.go
```
*/
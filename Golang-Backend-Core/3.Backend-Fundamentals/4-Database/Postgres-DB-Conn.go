package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

const dsn = "host=localhost port=5432 user=admin password=secret dbname=jobdb sslmode=disable"

func main() {
	fmt.Println("=== 1. CONNECTING TO POSTGRESQL ===")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open DB pool: %v", err)
	}; defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}; fmt.Println("[SUCCESS] Connected to PostgreSQL successfully!")


	fmt.Println("\n=== 2. EXECUTING DDL (CREATE TABLE) ===")
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS jobs (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		status VARCHAR(20) DEFAULT 'PENDING'
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}; fmt.Println("[SUCCESS] Table 'jobs' is ready!")


	fmt.Println("\n=== 3. EXECUTING DML (INSERT DATA) ===")
	insertQuery := `INSERT INTO jobs (name, status) VALUES ($1, $2) RETURNING id;`
	var newJobID int
	err = db.QueryRow(insertQuery, "Backup Database", "RUNNING").Scan(&newJobID)
	if err != nil {
		log.Fatalf("Failed to insert job: %v", err)
	}; fmt.Printf("[SUCCESS] Inserted new job with ID: %d\n", newJobID)


	fmt.Println("\n=== 4. EXECUTING DQL (QUERY DATA) ===")
	var jobName, jobStatus string
	selectQuery := `SELECT name, status FROM jobs WHERE id = $1`
	err = db.QueryRow(selectQuery, newJobID).Scan(&jobName, &jobStatus)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("[WARNING] No job found with that ID.")
		} else {
			log.Fatalf("Failed to query job: %v", err)
		}
	} else {
		fmt.Printf("[SUCCESS] Fetched Job -> Name: '%s', Status: '%s'\n", jobName, jobStatus)
	}
}

/*
Path: Golang-Backend-Core/3.Backend-Fundamentals/4-Database/Postgres-DB-Conn.go
```bash
go get github.com/lib/pq
docker run --name pg-test -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=jobdb -p 5432:5432 -d postgres
go run Postgres-DB-Conn.go
```
*/
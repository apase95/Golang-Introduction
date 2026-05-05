package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func submitJobHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[ERROR]: Handling new job...")
	time.Sleep(1 * time.Second)
	
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{
		"status": "Created",
		"job_id": "JOB-01"
	}`))
}

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[CHECKPOINT 1] --> Request started: %s %s", r.Method, r.URL.Path)
		next(w, r)
		log.Printf("[CHECKPOINT 1] <-- Request completed: %s %s (Took: %v)", r.Method, r.URL.Path, time.Since(start))
	}
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("[CHECKPOINT 2] Checking security badge (API Key)...")
		
		apiKey := r.Header.Get("X-API-Key")
		if apiKey != "super-secret-token" {
			log.Println("[CHECKPOINT 2] ALERT: Invalid API Key! Access blocked.")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error": "Unauthorized"}`))
			return
		}

		log.Println("[CHECKPOINT 2] Badge valid, allowed to pass!")
		next(w, r)
	}
}


func main() {
	mux := http.NewServeMux()
	finalHandler := LoggingMiddleware(AuthMiddleware(submitJobHandler2))
	mux.HandleFunc("POST /api/v1/jobs", finalHandler)
	fmt.Println("🚀 Server is running at http://localhost:8088")
	log.Fatal(http.ListenAndServe(":8088", mux))
}

/*
Path: Golang-Backend-Core/3.Backend-Fundamentals/3-Middleware/Middleware-Pattern.go
```bash
go run Middleware-Pattern.go
curl -X POST http://localhost:8088/api/v1/jobs
curl -X POST -H "X-API-Key: super-secret-token" http://localhost:8088/api/v1/jobs
```
*/
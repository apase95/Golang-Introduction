package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func CreatePipeline(coreHandler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for i := len(middlewares) -1; i >= 0; i-- {
		coreHandler = middlewares[i](coreHandler)
	}
	return coreHandler
}

func RecoverMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[EMERGENCY RECOVERED] Panic prevented: %v\n", err)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "Internal Server Error"}`))
			}
		}()
		next(w, r)
	}
}

func LoggingMiddleware2(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[LOGGING] --> Request started: %s %s\n", r.Method, r.URL.Path)
		next(w, r)
		log.Printf("[LOGGING] <-- Request completed: %s %s (Took: %v)\n", r.Method, r.URL.Path, time.Since(start))
	}
}

func AuthMiddleware2(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("[AUTH] Checking security API Key...")
		apiKey := r.Header.Get("X-API-Key")
		if apiKey != "super-secret-token" {
			log.Println("[AUTH] ALERT: Invalid API Key! Access blocked.")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error": "Unauthorized"}`))
			return
		}
		log.Println("[AUTH] Badge valid, allowed to pass!")
		next(w, r)		
	}
}

func submitJobHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("   [CORE] Handling new job...")
	if r.Header.Get("X-Trigger-Panic") == "true" {
		panic("Database connection lost unexpectedly!")
	}
	time.Sleep(500 * time.Millisecond)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{
		"status": "Created",
		"job_id": "JOB-PIPELINE-01"
	}`))
}

func main() {
	mux := http.NewServeMux()
	jobPipeline := CreatePipeline(
		submitJobHandler,
		RecoverMiddleware,
		LoggingMiddleware2,
		AuthMiddleware2,
	)
	mux.HandleFunc("POST /api/v1/jobs", jobPipeline)
	port := ":8082"
	fmt.Printf("🚀 Pipeline Server is running at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, mux))
}

/*
Path: Golang-Backend-Core/3.Backend-Fundamentals/3-Middleware/Middleware-Pipeline.go
```bash
go run Middleware-Pipeline.go
curl -X POST http://localhost:8082/api/v1/jobs
curl -X POST -H "X-API-Key: super-secret-token" http://localhost:8082/api/v1/jobs
curl -X POST -H "X-API-Key: super-secret-token" -H "X-Trigger-Panic: true" http://localhost:8082/api/v1/jobs
```
*/
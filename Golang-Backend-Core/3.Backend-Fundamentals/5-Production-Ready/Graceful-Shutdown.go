package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func slowTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[API] Request reveived. Processing task...")
	time.Sleep(5 * time.Second)

	log.Println("[API] Task completed successfully!")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "Success", "message": "Slow task finished!"}`))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/slow-task", slowTaskHandler)
	
	port := ":8080"
	srv := &http.Server {
		Addr: 		port,
		Handler: 	mux,
	}

	go func() {
		fmt.Printf("🚀 Server is runnig on http://localhost%s\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen and Serve Error: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	fmt.Println("\n[SYSTEM] Received shutdown signal. Initiating Graceful Shutdown...")
	
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("[SYSTEM] Server forced to shutdown due to error: %v", err)
	}; fmt.Println("[SYSTEM] Server exited gracefully!")
}

/*
Path: Golang-Backend-Core/3.Backend-Fundamentals/5-Production-Ready/Graceful-Shutdown.go
```bash
go run Graceful-Shutdown.go
curl -X GET http://localhost:8080/api/v1/slow-task
``` 
*/
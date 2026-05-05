package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! Welcome to the basic Go HTTP Server.\n")
	fmt.Fprintf(w, "You requested the path: %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", helloHandler)
	
	port := ":8081"
	fmt.Printf("🚀 Basic Server is running on http://localhost%s\n", port)
	
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server crashed: %v", err)
	}
}

/*
Path: Golang-Backend-Core/3.Backend-Fundamentals/1-Server-Routing/HTTP-Server.go
```bash
go run HTTP-Server.go
curl http://localhost:8081/
curl http://localhost:8081/jobs/job-01
```
*/
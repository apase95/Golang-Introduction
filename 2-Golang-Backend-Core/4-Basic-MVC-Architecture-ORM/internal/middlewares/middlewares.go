package middlewares

import (
	"log"
	"net/http"
	"time"
)

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[API_START] --> %s %s\n", r.Method, r.URL.Path)
		next(w, r)
		log.Printf("[API_END] <-- %s %s (Took: %v)\n", r.Method, r.URL.Path, time.Since(start))
	}
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey != "super-secret" {
			log.Println("[AUTH] Unauthorized access attempt blocked!")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error": "Unauthorized. Please provide a valid X-API-Key"}`))
		}
		next(w, r)
	}
}
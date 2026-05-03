package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type APIResponse struct {
	Status 	string	`json:"status"`
	Message string	`json:"message"`
	JobID 	string	`json:"job_id,omitempty"`
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(APIResponse{
		Status: "OK",
		Message: "Server is up and running smoothly",
	})
}

func submitJobHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(APIResponse{
		Status: "Created",
		Message: "Job has been added to queue",
		JobID: "JOB-01",
	})
}

func getJobHandler(w http.ResponseWriter, r *http.Request) {
	jobID := r.PathValue("id")
	w.Header().Set("Content-Type", "application/json")
	
	if jobID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(APIResponse {
			Status:  "Error",
			Message: "Missing JobID",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(APIResponse {
		Status:  "Success",
		Message: fmt.Sprintf("Found details for job %s", jobID),
		JobID:   jobID,
	})
}


func main() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("GET /health", healthCheckHandler)
	mux.HandleFunc("POST /api/v1/jobs", submitJobHandler)
	mux.HandleFunc("GET /api/v1/jobs/{id}", getJobHandler)

	port := ":8088"
	fmt.Printf("🚀 Server running on http://localhost%s\n", port)

	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
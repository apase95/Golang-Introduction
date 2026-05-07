package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type jobPayload struct {
	Name 		string 	`json:"name"`
	Command 	string 	`json:"command"`
	Retries 	int 	`json:"retries"`
	AlertEmail 	string 	`json:"alert_email,omitempty"`
	internalID 	string 	`json:"-"`
}

func main() {
	fmt.Println("=== 1. UNMARSHAL ===")
	jsonInput :=[]byte(`{
		"name": "Backup Database",
		"command": "pg_dump -U admin -d myDB",
		"retries": 3
	}`)
	
	var job jobPayload
	err := json.Unmarshal(jsonInput, &job)
	if err != nil {
		log.Fatalf("ERROR: Decode JSON: %v", err)
	}
	fmt.Printf("[SUCCESS]: Job: %s | Command: %s | Email: '%s' \n", job.Name, job.Command, job.AlertEmail)


	fmt.Println("=== 2. MARSHAL ===")
	response := map[string]any{
		"status": "success",
		"job_id": "JOB-01",
		"details": "job",
	}
	responseBytes, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatalf("ERROR: Decode JSON: %v", err)
	}
	fmt.Printf("Generated Data:\n%s\n", string(responseBytes))


	fmt.Println("=== 3. JSON STREAM ===")
	streamInput := strings.NewReader(`{
		"name": "Clean Logs",
		"command": "rm -rf /var/log/*", 
		"retries": 0,
		"alert_email": "noobyhandsome.company@gmail.com"
	}`)

	var streamJob jobPayload
	decoder := json.NewDecoder(streamInput)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&streamJob); err != nil {
		log.Fatalf("ERROR: Decode JSON: %v", err)
	}
	fmt.Printf("[SUCCESS] Job: %s | Email: %s\n", streamJob.Name, streamJob.AlertEmail)
}
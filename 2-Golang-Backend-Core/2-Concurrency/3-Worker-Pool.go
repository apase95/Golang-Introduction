package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID 		int
	Command string
}

type Result struct {
	JobID 	int
	Output 	string
}

func worker(workerID int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("[Worker %d] Started processing Job %d (%s)\n", workerID, job.ID, job.Command)
		time.Sleep(1 * time.Second)
		fmt.Printf("[Worker %d] Finished Job %d\n", workerID, job.ID)
		results <- Result {
			JobID: job.ID,
			Output: fmt.Sprintf("Success run %s", job.Command),
		}
	}
}

func main() {
	const numJobs = 5
	const numWorkers = 3
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)
	var  wg sync.WaitGroup

	fmt.Println ("--- STARTING WORKER POOL ---")
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- Job {
			ID: j, 
			Command: fmt.Sprintf("script_%d.sh", j),
		}
	}; close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("\n--- COLLECTING RESULT ---")
	for res := range results {
		fmt.Printf("[Result] Job %d Output: %s\n", res.JobID, res.Output)
	}
}
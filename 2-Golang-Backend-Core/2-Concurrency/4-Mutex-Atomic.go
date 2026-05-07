package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type ServerStats struct {
	mu sync.Mutex
	JobStatus map[string]int
	TotalProcessed int64
}

func (s *ServerStats) RecordJob(status string) {
	atomic.AddInt64(&s.TotalProcessed, 1)
	s.mu.Lock()
	s.JobStatus[status]++
	s.mu.Unlock()
}

func main() {
	stats := &ServerStats{
		JobStatus: make(map[string]int),
	}
	
	fmt.Println("Simulating 1000 concurrent jobs reporting their status...")
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if id % 10 == 0 {
				stats.RecordJob("FAILED")
			} else {
				stats.RecordJob("SUCCESS")
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("\n--- FINAL SERVER STATS ---")
	fmt.Printf("Total Jobs Processed : %d\n", stats.TotalProcessed)
	fmt.Printf("Successful Jobs      : %d\n", stats.JobStatus["SUCCESS"])
	fmt.Printf("Failed Jobs          : %d\n", stats.JobStatus["FAILED"])
}
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func checkServer(server string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("[...]Checking server %s\n", server)

	delay := time.Duration(rand.Intn(3)+1) * time.Second
	time.Sleep(delay)
	fmt.Printf("[OK] %s responded in %v\n", server, delay)
}

func main() {
	servers :=[]string{"web-01", "db-01", "cache-01", "worker-01"}
	var wg sync.WaitGroup

	fmt.Printf("STARTING CHECKS\n")
	startTime := time.Now()
	for _, srv := range servers {
		wg.Add(1)
		go checkServer(srv, &wg)
	}

	extraJobs :=[]string{"backup-job", "clean-log-job"}
	for _, job := range extraJobs {
		wg.Add(1)

		go func(jobName string) {
			defer wg.Done()
			fmt.Printf("[...]Starting job %s\n", jobName)
			time.Sleep(2 * time.Second)
			fmt.Printf("[OK] %s completed\n", jobName)
		}(job)
	}

	fmt.Printf("Main is blocking and waiting for jobs to complete...\n")
	wg.Wait()
	fmt.Printf("ALL CHECKS AND JOBS COMPLETED in %v\n", time.Since(startTime))
}
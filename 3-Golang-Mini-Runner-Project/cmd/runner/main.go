package main

import (
	"flag"
	"fmt"
	"time"

	"golang-mini-runner-project/internal/config"
	"golang-mini-runner-project/internal/models"
	"golang-mini-runner-project/internal/queue"
	"golang-mini-runner-project/internal/repository"
	"golang-mini-runner-project/internal/worker"
)

func main() {
	workersCount := flag.Int("workers", 2, "Number of concurrent workers in the pool")
	flag.Parse()
	
	appConfig := config.LoadConfig()
	config.PrintConfig(appConfig)

	fmt.Println("\n=== INITIALIZING SYSTEM ===")
	store := repository.NewMemoryStore()
	jobQueue := queue.NewMemoryQueue(10)
	
	pool := worker.NewWorkerPool(*workersCount, jobQueue, store)
	pool.Start()

	fmt.Println("=== INJECTING TEST JOBS ===")
	job1 := &models.Job{Name: "List Files", Command: "echo 'Starting...' && ls -la && echo 'Done!'"}
	store.Save(job1)
	jobQueue.Enqueue(job1)

	job2 := &models.Job{Name: "Fail Task", Command: "ls /folder_does_not_exist"}
	store.Save(job2)
	jobQueue.Enqueue(job2)

	job3 := &models.Job{Name: "Slow Task", Command: "echo 'Sleeping for 3 seconds...' && sleep 3 && echo 'Woke up!'"}
	store.Save(job3)
	jobQueue.Enqueue(job3)

	time.Sleep(5 * time.Second)

	fmt.Println("\n=== SHUTTING DOWN SYSTEM ===")
	jobQueue.Close()
	pool.Stop()

	fmt.Println("\n=== DATABASE FINAL STATE ===")
	for i := uint(1); i <= 3; i++ {
		j, _ := store.GetByID(i)
		fmt.Printf("Job #%d [%s]: %s | Log saved at: %s\n", j.ID, j.Status, j.Name, j.LogsPath)
	}
}
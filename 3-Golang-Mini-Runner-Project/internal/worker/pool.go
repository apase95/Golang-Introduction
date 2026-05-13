package worker

import (
	"golang-mini-runner-project/internal/queue"
	"golang-mini-runner-project/internal/repository"
	"log"
	"sync"
	"time"
)

type WorkerPool struct {
	workerCount int
	queue 		*queue.MemoryQueue
	store 		repository.JobStore
	wg			sync.WaitGroup
}

func NewWorkerPool(count int, q *queue.MemoryQueue, s repository.JobStore) *WorkerPool {
	return &WorkerPool{
		workerCount: 	count,
		queue: 			q,
		store: 			s,
	}
}

func (p *WorkerPool) Start() {
	log.Printf("[WorkerPool] Starting %d workers....\n", p.workerCount)
	for i := 1; i <= p.workerCount; i++ {
		p.wg.Add(1)
		go p.workerLoop(i)
	}
}

func (p *WorkerPool) workerLoop(id int) {
	defer p.wg.Done()
	log.Printf("[Worker %d] Ready and listening for jobs...\n", id)

	for job := range p.queue.Dequeue() {
		log.Printf("[Worker %d] Picked up job #%d: %s\n", id, job.ID, job.Command)

		job.Status = "RUNNING"
		_ = p.store.Update(job)
		time.Sleep(2 * time.Second)

		now := time.Now()
		job.Status = "SUCCESS"
		job.FinishedAt = &now
		_ = p.store.Update(job)

		log.Printf("[Worker %d] Successfully finished Job #%d\n", id, job.ID)
	}

	log.Printf("[Worker %d] Queue closed. Worker shutting down...\n", id)
}

func (p *WorkerPool) Stop() {
	log.Println("[WorkerPool] Waiting for active workers to finish...")
	p.wg.Wait()
	log.Println("[WorkerPool] All workers successfully stopped.")
}
package queue

import (
	"errors"
	"golang-mini-runner-project/internal/models"
)

type MemoryQueue struct {
	jobs chan *models.Job
}

func NewMemoryQueue(bufferSize int) *MemoryQueue {
	return &MemoryQueue{
		jobs: make(chan *models.Job, bufferSize),
	}
}

func (q *MemoryQueue) Enqueue(job *models.Job) error {
	select {
	case q.jobs <- job:
		return nil
	default:
		return errors.New("queue is full, pls try again later")
	}
}

func (q *MemoryQueue) Dequeue() <-chan *models.Job {
	return q.jobs
}

func (q *MemoryQueue) Close() {
	close(q.jobs)
}
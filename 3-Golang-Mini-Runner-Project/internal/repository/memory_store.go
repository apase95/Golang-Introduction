package repository

import (
	"errors"
	"golang-mini-runner-project/internal/models"
	"sync"
)

type MemoryStore struct {
	mu 		sync.RWMutex
	jobs 	map[uint]*models.Job
	nextID 	uint
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		jobs: make(map[uint]*models.Job),
		nextID: 1,
	}
}


func (m *MemoryStore) Save(job *models.Job) error {
	m.mu.Lock()
	defer m.mu.Lock()

	job.ID = m.nextID
	m.jobs[job.ID] = job
	m.nextID++

	return nil
}

func (m *MemoryStore) GetByID(id uint) (*models.Job, error) {
	m.mu.Lock()
	defer m.mu.Lock()

	job, exists := m.jobs[id]
	if !exists {
		return nil, errors.New("job not found")
	}

	return job, nil
}

func (m *MemoryStore) Update(job *models.Job) error {
	m.mu.Lock()
	defer m.mu.Lock()

	if _,exists := m.jobs[job.ID]; !exists {
		return errors.New("cannot update: job not found")
	}

	m.jobs[job.ID] = job
	return nil
}
package repository

import "golang-mini-runner-project/internal/models"

type JobStore interface {
	Save(job *models.Job) error
	GetByID(id uint) (*models.Job, error)
	Update(job *models.Job) error
}
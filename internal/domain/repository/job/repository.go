package jobRepositoryAbs

import jobEntity "jobs.api.com/internal/domain/entities/job"

type JobRepositoryInterface interface {
	Create(jobPayload *jobEntity.Job) error
	GetByID(id string) (*jobEntity.Job, error)
	List() ([]jobEntity.Job, error)
	GetAll() ([]*jobEntity.Job, error)
}

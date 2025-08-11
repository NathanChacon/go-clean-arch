package jobUsecase

import (
	jobEntity "jobs.api.com/internal/domain/entities/job"
	jobRepositoryAbs "jobs.api.com/internal/domain/repository/job"
	uuidInterface "jobs.api.com/internal/infrastructure/utils/interfaces/uuid"
)

type JobDTO struct {
	UUID        string `json:"uuid,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
	CompanyID   string `json:"company_id"`
	CreatedBy   string `json:"created_by"`
}

type UseCase interface {
	PostJob(payload JobDTO) error
	GetById(uuid string) (*jobEntity.Job, error)
	GetAll() ([]*jobEntity.Job, error)
}

type JobUseCase struct {
	repository    jobRepositoryAbs.JobRepositoryInterface
	uuidGenerator uuidInterface.UuidGenerator
}

func NewJobUseCase(repository jobRepositoryAbs.JobRepositoryInterface, uuidGenerator uuidInterface.UuidGenerator) *JobUseCase {
	return &JobUseCase{repository: repository, uuidGenerator: uuidGenerator}
}

func (useCase *JobUseCase) PostJob(payload JobDTO) error {
	var formattedPayload = jobEntity.NewJobParams{
		UUID:        useCase.uuidGenerator.NewUuid(),
		Title:       payload.Title,
		Description: payload.Description,
		Location:    payload.Location,
		CompanyID:   payload.CompanyID,
		CreatedBy:   payload.CreatedBy,
	}

	job, jobError := jobEntity.NewJobEntity(formattedPayload)

	if jobError != nil {
		return jobError
	}

	err := useCase.repository.Create(job)

	return err
}

func (useCase *JobUseCase) GetAll() ([]*jobEntity.Job, error) {
	jobs, err := useCase.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func (useCase *JobUseCase) GetById(uuid string) (*jobEntity.Job, error) {
	job, err := useCase.repository.GetByID(uuid)

	return job, err
}

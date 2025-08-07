package jobUsecase

import (
	jobEntity "jobs.api.com/internal/domain/entities/job"
	jobRepositoryAbs "jobs.api.com/internal/domain/repository/job"
	uuidInterface "jobs.api.com/internal/interfaces/uuid"
)

type UseCase interface {
	PostJob(payload jobEntity.Job) error
	GetById(uuid string) (*jobEntity.Job, error)
}

type JobUseCase struct {
	repository    jobRepositoryAbs.JobRepositoryInterface
	uuidGenerator uuidInterface.UuidGenerator
}

func NewJobUseCase(repository jobRepositoryAbs.JobRepositoryInterface, uuidGenerator uuidInterface.UuidGenerator) *JobUseCase {
	return &JobUseCase{repository: repository, uuidGenerator: uuidGenerator}
}

func (useCase *JobUseCase) PostJob(payload jobEntity.Job) error {
	payload.UUID = useCase.uuidGenerator.NewUuid()

	err := useCase.repository.Create(payload)

	return err
}

func (useCase *JobUseCase) GetById(uuid string) (*jobEntity.Job, error) {
	job, err := useCase.repository.GetByID(uuid)

	return job, err
}

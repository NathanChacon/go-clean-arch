package jobUsecase

// add job entity validation
// check if user existis on database (needs to receive User repo)
import (
	jobEntity "jobs.api.com/internal/domain/entities/job"
	jobRepositoryAbs "jobs.api.com/internal/domain/repository/job"
)

type UseCase interface {
	PostJob(payload *jobEntity.Job) error
	GetById(uuid string) (*jobEntity.Job, error)
}

type JobUseCase struct {
	repository jobRepositoryAbs.JobRepositoryInterface
}

func NewJobUseCase(repository jobRepositoryAbs.JobRepositoryInterface) *JobUseCase {
	return &JobUseCase{repository: repository}
}

func (useCase *JobUseCase) PostJob(payload *jobEntity.Job) error {
	err := useCase.repository.Create(payload)

	return err
}

func (useCase *JobUseCase) GetById(uuid string) (*jobEntity.Job, error) {
	job, err := useCase.repository.GetByID(uuid)

	return job, err
}

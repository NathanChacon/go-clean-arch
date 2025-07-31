package jobUsecase

// add job entity validation
// check if user existis on database (needs to receive User repo) ;;;
import (
	jobEntity "jobs.api.com/internal/domain/entities/job"
	jobRepositoryAbs "jobs.api.com/internal/domain/repository/job"
	userRepositoryAbs "jobs.api.com/internal/domain/repository/user"
)

type UseCase interface {
	PostJob(payload *jobEntity.Job) error
	GetById(uuid string) (*jobEntity.Job, error)
}

type JobUseCase struct {
	repository     jobRepositoryAbs.JobRepositoryInterface
	userRepository userRepositoryAbs.UserRepositoryAbs
}

func NewJobUseCase(repository jobRepositoryAbs.JobRepositoryInterface, userRepository userRepositoryAbs.UserRepositoryAbs) *JobUseCase {
	return &JobUseCase{repository: repository, userRepository: userRepository}
}

func (useCase *JobUseCase) PostJob(payload *jobEntity.Job) error {
	err := useCase.repository.Create(payload)

	return err
}

func (useCase *JobUseCase) GetById(uuid string) (*jobEntity.Job, error) {
	job, err := useCase.repository.GetByID(uuid)

	return job, err
}

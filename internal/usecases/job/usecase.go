package jobUsecase

import (
	jobEntity "jobs.api.com/internal/domain/entities/job"
	userEntity "jobs.api.com/internal/domain/entities/user"
	jobRepositoryAbs "jobs.api.com/internal/domain/repository/job"
	userRepositoryAbs "jobs.api.com/internal/domain/repository/user"
)

type UuidGenerator interface {
	NewUuid() string
}

type UseCase interface {
	PostJob(payload jobEntity.Job) error
	GetById(uuid string) (*jobEntity.Job, error)
}

type JobUseCase struct {
	repository     jobRepositoryAbs.JobRepositoryInterface
	userRepository userRepositoryAbs.UserRepositoryAbs
	uuidGenerator  UuidGenerator
}

func (usecase *JobUseCase) isUserRegistered(userId string) (userEntity.User, error) {
	user, err := usecase.userRepository.GetById(userId)

	return user, err
}

func NewJobUseCase(repository jobRepositoryAbs.JobRepositoryInterface, userRepository userRepositoryAbs.UserRepositoryAbs, uuidGenerator UuidGenerator) *JobUseCase {
	return &JobUseCase{repository: repository, userRepository: userRepository, uuidGenerator: uuidGenerator}
}

func (useCase *JobUseCase) PostJob(payload jobEntity.Job) error {
	payload.UUID = useCase.uuidGenerator.NewUuid()
	_, userErr := useCase.isUserRegistered(payload.CreatedBy)

	if userErr != nil {
		return userErr
	}

	err := useCase.repository.Create(payload)

	return err
}

func (useCase *JobUseCase) GetById(uuid string) (*jobEntity.Job, error) {
	job, err := useCase.repository.GetByID(uuid)

	return job, err
}

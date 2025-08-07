package userUseCase

import (
	"errors"
	"fmt"

	userEntity "jobs.api.com/internal/domain/entities/user"
	domainErrors "jobs.api.com/internal/domain/errors"
	userRepositoryAbs "jobs.api.com/internal/domain/repository/user"
	uuidInterface "jobs.api.com/internal/interfaces/uuid"
)

type CreateAccountParams struct {
	Name     string
	Password string
	Email    string
}

type UseCase interface {
	GetById(id string) (userEntity.User, error)
	CreateUser(payload CreateAccountParams) error
}

type UserUseCase struct {
	repository    userRepositoryAbs.UserRepositoryAbs
	uuidGenerator uuidInterface.UuidGenerator
}

func NewUserUseCase(repository userRepositoryAbs.UserRepositoryAbs, uuidGenerator uuidInterface.UuidGenerator) *UserUseCase {
	return &UserUseCase{repository: repository, uuidGenerator: uuidGenerator}
}

func (useCase *UserUseCase) GetById(id string) (userEntity.User, error) {
	user, err := useCase.repository.GetById(id)
	if errors.Is(err, domainErrors.ErrUserNotFound) {
		return user, fmt.Errorf("GetByID: job %s not found: %w", id, err)
	}
	return user, err
}

func (useCase *UserUseCase) CreateUser(payload CreateAccountParams) error {
	uuid := useCase.uuidGenerator.NewUuid()
	user, err := userEntity.NewUserEntity(uuid, payload.Name, payload.Email, payload.Password)

	if err != nil {
		return err
	}

	repoErr := useCase.repository.CreateUser(user)

	if repoErr != nil {
		return repoErr
	}

	return nil

}

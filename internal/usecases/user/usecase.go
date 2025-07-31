package userUseCase

import (
	userEntity "jobs.api.com/internal/domain/entities/user"
	userRepositoryAbs "jobs.api.com/internal/domain/repository/user"
)

type UseCase interface {
	CreateUser(payload *userEntity.User) error
	GetById(id string) (*userEntity.User, error)
}

type UserUseCase struct {
	repository userRepositoryAbs.UserRepositoryAbs
}

func NewUserUseCase(repository userRepositoryAbs.UserRepositoryAbs) *UserUseCase {
	return &UserUseCase{repository: repository}
}

func (useCase *UserUseCase) CreateUser(payload userEntity.User) error {
	err := useCase.repository.CreateUser(payload)

	return err
}

func (useCase *UserUseCase) GetById(id string) (userEntity.User, error) {
	user, err := useCase.repository.GetById(id)

	return user, err
}

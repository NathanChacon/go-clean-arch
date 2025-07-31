package userRepositoryAbs

import userEntity "jobs.api.com/internal/domain/entities/user"

type UserRepositoryAbs interface {
	CreateUser(payload userEntity.User) error
	GetById(id string) (userEntity.User, error)
}

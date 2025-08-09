package authenticationUseCase

import (
	userEntity "jobs.api.com/internal/domain/entities/user"
	domainErrors "jobs.api.com/internal/domain/errors"
	userRepositoryAbs "jobs.api.com/internal/domain/repository/user"
	hasherInterface "jobs.api.com/internal/infrastructure/utils/interfaces/hasher"
)

type Authenticator interface {
	Login(email string, password string) (userEntity.User, error)
}

type AuthenticationUseCase struct {
	userRepository userRepositoryAbs.UserRepositoryAbs
	passwordHasher hasherInterface.PasswordHasher
}

func NewAutheticationUseCase(userRepository userRepositoryAbs.UserRepositoryAbs) *AuthenticationUseCase {
	return &AuthenticationUseCase{userRepository: userRepository}
}

func (useCase *AuthenticationUseCase) Login(email string, password string) (userEntity.User, error) {
	user, err := useCase.userRepository.GetByEmail(email)

	if err != nil {
		return user, err
	}

	hashErr := useCase.passwordHasher.Compare(user.Password, password)

	if hashErr != nil {
		return user, domainErrors.ErrInvalidPasswordLogin
	}

	return user, nil
}

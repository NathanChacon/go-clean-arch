package authenticationUseCase

import (
	"fmt"

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

func NewAutheticationUseCase(userRepository userRepositoryAbs.UserRepositoryAbs, passwordHasher hasherInterface.PasswordHasher) *AuthenticationUseCase {
	return &AuthenticationUseCase{userRepository: userRepository, passwordHasher: passwordHasher}
}

func (useCase *AuthenticationUseCase) Login(email string, password string) (userEntity.User, error) {
	user, err := useCase.userRepository.GetByEmail(email)

	if err != nil {
		fmt.Println("get error", err)
		return user, err
	}

	hashErr := useCase.passwordHasher.Compare(user.Password, password)

	if hashErr != nil {
		fmt.Println("hash error", hashErr)
		return user, domainErrors.ErrInvalidPasswordLogin
	}

	return user, nil
}

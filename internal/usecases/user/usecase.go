package userUseCase

import (
	"errors"
	"fmt"
	"regexp"
	"unicode"

	userEntity "jobs.api.com/internal/domain/entities/user"
	domainErrors "jobs.api.com/internal/domain/errors"
	userRepositoryAbs "jobs.api.com/internal/domain/repository/user"
)

func isValidName(name string) bool {
	if len(name) < 4 {
		return false
	}

	for _, c := range name {
		if !unicode.IsLetter(c) {
			return false
		}
	}

	return true

}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func isValidPassword(password string) bool {
	var hasMinLen = len(password) >= 8
	var hasNumber, hasSpecial bool

	for _, c := range password {
		if unicode.IsNumber(c) {
			hasNumber = true
		} else if unicode.IsPunct(c) || unicode.IsSymbol(c) {
			hasSpecial = true
		}

		if hasNumber && hasSpecial {
			break
		}
	}

	return hasMinLen && hasNumber && hasSpecial
}

type UseCase interface {
	CreateUser(payload userEntity.User) error
	GetById(id string) (userEntity.User, error)
	CreateAccount(payload userEntity.User) error
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
	if errors.Is(err, domainErrors.ErrUserNotFound) {
		return user, fmt.Errorf("GetByID: job %s not found: %w", id, err)
	}
	return user, err
}

func (useCase *UserUseCase) CreateAccount(payload userEntity.User) error {
	if !isValidEmail(payload.Email) {
		return fmt.Errorf("invalid user email: %s: %w", payload.Email, domainErrors.ErrInvalidEmailFormat)
	}

	if !isValidName(payload.Name) || !isValidPassword(payload.Password) {
		return fmt.Errorf("invalid user credentials: %w", domainErrors.ErrInvalidUserCredential)
	}

	err := useCase.repository.CreateUser(payload)

	if err != nil {
		return err
	}

	return nil

}

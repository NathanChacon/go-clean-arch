package userEntity

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	domainErrors "jobs.api.com/internal/domain/errors"
)

type Password struct {
	Plain  string
	Hashed string
}

type User struct {
	UUID     string
	Email    string
	Password string
	Name     string
	Location string
	Bio      string
}

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
	fmt.Print(email)
	email = strings.TrimSpace(email)
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

func NewUserEntity(uuid string, name string, email string, password Password) (user User, err error) {

	if uuid == "" {
		return User{}, domainErrors.ErrInvalidUuid
	}

	if !isValidEmail(email) {
		return User{}, domainErrors.ErrInvalidEmailFormat
	}

	if !isValidName(name) || !isValidPassword(password.Plain) {
		return User{}, domainErrors.ErrInvalidUserCredential
	}

	return User{
		UUID:     uuid,
		Name:     name,
		Password: password.Hashed,
		Email:    email,
	}, nil

}

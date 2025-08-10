package domainErrors

import "errors"

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrInvalidUserCredential = errors.New("invalid username or password")
	ErrInvalidEmailFormat    = errors.New("email format is invalid")
	ErrInvalidUuid           = errors.New("invalid uuid")
	ErrUserAlreadyRegistered = errors.New("user already registered")
	ErrInvalidPasswordLogin  = errors.New("wrong password")
)

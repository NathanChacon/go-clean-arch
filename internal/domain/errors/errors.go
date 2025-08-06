package domainErrors

import "errors"

var (
	ErrJobNotFound           = errors.New("job not found")
	ErrUserNotFound          = errors.New("user not found")
	ErrInvalidUserCredential = errors.New("invalid username or password")
	ErrInvalidEmailFormat    = errors.New("email format is invalid")
)

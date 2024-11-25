package common

import "errors"

var (
	ErrInvalidEmailOrPassword = errors.New("Invalid Email or Password")
	ErrEmailAlreadyExist      = errors.New("Email already exist")
	ErrUserNotFound           = errors.New("User not found")
	ErrInternalSystem         = errors.New("Internal system error")
)
